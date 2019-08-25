package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/webhook"
)

func (s *apiServer) WebhookJobHandler(w http.ResponseWriter, r *http.Request) {
	repoUuid := mux.Vars(r)["repo_uuid"]
	providedSignature := []byte(r.Header.Get("X-Hub-Signature"))

	repo, err := s.Storage.FindRepoByUUID(repoUuid)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !webhook.VerifySha1(reqBody, []byte(repo.AuthenticationCodeSecret), providedSignature) {
		httputils.RespondWithError(
			w,
			http.StatusUnauthorized,
			"Signature verification failed. Please check your application configuration.",
		)
		return
	}

	provider, err := s.Storage.GetProviderForRepo(repo.ID)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	configs, err := s.Storage.LookupWebhookRepoConfigurations(repoUuid)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	webhookBody := &webhook.GithubWebhookBody{}
	if err := json.Unmarshal(reqBody, webhookBody); err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sourceRef := webhookBody.Ref
	if sourceRef == "" {
		sourceRef = fmt.Sprintf("refs/heads/%s", webhookBody.Repository.DefaultBranch)
	}

	buildDefinitions := []*entities.BuildDefinition{}
	for _, cfg := range configs {
		job := &entities.BuildJob{
			Status:               entities.BuildJobStatusCreated,
			BuildConfigurationID: cfg.ID,
			SourceUri:            webhookBody.Repository.CloneURL,
			SourceRevision:       &sourceRef,
		}
		newJob, err := s.Storage.CreateBuildJob(job)
		if err != nil {
			httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		buildDef := &entities.BuildDefinition{
			Provider: provider,
			Config:   &cfg,
			Repo:     repo,
			Job:      newJob,
		}
		buildDefinitions = append(buildDefinitions, buildDef)
	}

	for _, d := range buildDefinitions {
		err = s.Builder.TaskQueue.EnqueueJob(d)
		if err != nil {
			httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	httputils.RespondWithStatus(w, http.StatusOK, "ok")
}
