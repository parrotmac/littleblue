package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

func initializeFrontendRoutes(router *mux.Router) {

	staticUrlPrefix := "/static/"
	clientDirectoryPath := "client/build/static/"

	staticFileServer := http.FileServer(http.Dir(clientDirectoryPath))
	router.PathPrefix(staticUrlPrefix).Handler(http.StripPrefix(staticUrlPrefix, staticFileServer))

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/build/")
	}
	router.PathPrefix("/").HandlerFunc(indexHandler)
}

func (a *App) getJobsRoute(w http.ResponseWriter, r *http.Request) {
	type WebSafeJob struct {
		RepoName string    `json:"repo_name"`
		Messages []Message `json:"messages"`
	}

	webSafeJobs := []WebSafeJob{}

	for _, buildContext := range a.buildContexts {
		webSafeJobs = append(webSafeJobs, WebSafeJob{
			RepoName: buildContext.Source.FullName,
			Messages: buildContext.Messages,
		})
	}

	httputils.RespondWithJSON(w, http.StatusOK, webSafeJobs)
}

func (a *App) webhookUpdate(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	providedSignature := []byte(r.Header.Get("X-Hub-Signature"))

	hmacSignatureValid := a.VerifyRequestBodyHmac(bodyBytes, []byte(a.config.GithubConfig.WebhookSecret), providedSignature)

	if !hmacSignatureValid {
		httputils.RespondWithError(w, http.StatusUnauthorized, "Signature verification failed. Please check your application configuration.")
		return
	}

	var webhookBody GithubWebhookRequest
	if err := json.Unmarshal(bodyBytes, &webhookBody); err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, "Unable to decode JSON body")
		return
	}

	repoDashedName := webhookBody.getDashedName()

	newBuildContext := &BuildContext{

		// TODO: Have GithubWebhookRequest implement interface to build a GithubRepository
		Source: GitRepository{
			FullName:   webhookBody.Repository.FullName,
			DashedName: repoDashedName,
			RepoName:   webhookBody.Repository.Name,
			GitRefSpec: RefSpec(webhookBody.Ref),
		},

		// TODO: Pull from a mapping
		Docker: DockerBuildSpec{
			RegistryURL:      a.config.DockerRegistryConfig.URL,
			RegistryUsername: a.config.DockerRegistryConfig.Username,
			RegistryPassword: a.config.DockerRegistryConfig.Password,
			Tag:              "latest",
		},
		Messages:         []Message{},
		broadcastChannel: &a.wsBroadcast,
		BuildIdentifier:  fmt.Sprint(int64(time.Now().Unix())),
	}

	oldBuildCtx := a.buildContexts
	a.buildContexts = append(oldBuildCtx, newBuildContext)

	go a.ProcessWebhook(newBuildContext, &webhookBody)

	log.Println(webhookBody)

	httputils.RespondWithStatus(w, http.StatusOK, "OK")
}
