package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/meta"
)

type BuildJobRouter struct {
	StorageService *db.Storage
	JobQueue       meta.JobQueue
}

func (router *BuildJobRouter) CreateBuildJobHandler(w http.ResponseWriter, r *http.Request) {
	buildJob := &entities.BuildJob{}
	err := httputils.ReadJsonBodyToEntity(r.Body, buildJob)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	_, err = router.StorageService.CreateBuildJob(buildJob)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *BuildJobRouter) WebhookJobHandler(w http.ResponseWriter, r *http.Request) {
	repoUuid, ok := mux.Vars(r)["repo_uuid"]
	if !ok {
		httputils.RespondWithError(w, http.StatusBadRequest, errors.New("malformed URL"))
		return
	}

	sourceRepo, err := router.StorageService.GetSourceRepositoryByUUID(repoUuid)
	if err != nil {
		httputils.RespondWithError(w, http.StatusNotFound, err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	providedSignature := []byte(r.Header.Get("X-Hub-Signature"))

	hmacSignatureValid, err := httputils.VerifyRequestBodyHmac(bodyBytes, []byte(sourceRepo.AuthenticationCodeSecret), providedSignature)
	if err != nil || !hmacSignatureValid {
		if err != nil {
			logrus.Error("hmac validation failed", err)
		}
		httputils.RespondWithError(w, http.StatusUnauthorized, errors.New("Signature verification failed. Please check your application configuration."))
		return
	}

	buildConfigs, err := router.StorageService.LookupWebhookRepoConfigurations(repoUuid)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	webhookRequest := &meta.GithubWebhookRequest{}
	err = json.Unmarshal(bodyBytes, webhookRequest)
	if err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	sourceRef := webhookRequest.Ref

	for _, buildConfig := range buildConfigs {
		buildIdentifier := uuid.New().String()
		job := &entities.BuildJob{
			BuildIdentifier:      buildIdentifier,
			BuildConfigurationID: buildConfig.ID,
			SourceUri:            webhookRequest.Repository.CloneURL,
			SourceReference:      &sourceRef,
			StartTime:            time.Now(),
		}
		createdJob, err := router.StorageService.CreateBuildJob(job)
		if err != nil {
			logrus.Errorln("failed to create build job", err)
			continue
		}
		logrus.Infoln("ID", strconv.Itoa(int(createdJob.ID)))
		router.JobQueue <- *createdJob
	}

	httputils.RespondWithStatus(w, http.StatusOK, "ok")
}
