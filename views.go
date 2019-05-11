package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) frontendRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style=\"font-family: sans-serif\">Little Blue</h1>"))
}

func (a *App) getJobsRoute(w http.ResponseWriter, r *http.Request) {
	type WebSafeJob struct {
		RepoName	string		`json:"repo_name"`
		Messages	[]Message	`json:"messages"`
	}

	webSafeJobs := []WebSafeJob{}

	for _, buildContext := range a.buildContexts {
		webSafeJobs = append(webSafeJobs, WebSafeJob{
			RepoName: buildContext.Source.FullName,
			Messages: buildContext.Messages,
		})
	}

	respondWithJSON(w, http.StatusOK, webSafeJobs)
}

func (a *App) webhookUpdate(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	providedSignature := []byte(r.Header.Get("X-Hub-Signature"))

	hmacSignatureValid := a.VerifyRequestBodyHmac(bodyBytes, []byte(a.AppSettings.githubWebhookSecret), providedSignature)

	if !hmacSignatureValid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Signature verification failed. Please check your application configuration."))
		return
	}

	var webhookBody GithubWebhookRequest
	if err := json.Unmarshal(bodyBytes, &webhookBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to decode JSON body"))
		return
	}

	repoDashedName := webhookBody.getDashedName()

	newBuildContext := &BuildContext{

		// TODO: Have GithubWebhookRequest implement interface to build a GithubRepository
		Source: GitRepository{
			FullName: webhookBody.Repository.FullName,
			DashedName: repoDashedName,
			RepoName: webhookBody.Repository.Name,
			GitRefSpec:	RefSpec(webhookBody.Ref),
		},

		// TODO: Pull from a mapping
		Docker: DockerBuildSpec{
			RegistryURL:a.AppSettings.dockerRegistryURL,
			RegistryUsername:a.AppSettings.dockerRegistryUsername,
			RegistryPassword:a.AppSettings.dockerRegistryPassword,
			Tag: "latest",
		},
		Messages: []Message{},
		broadcastChannel: &a.wsBroadcast,
		BuildIdentifier: fmt.Sprint(int64(time.Now().Unix())),
	}

	oldBuildCtx := a.buildContexts
	a.buildContexts = append(oldBuildCtx, newBuildContext)

	go a.ProcessWebhook(newBuildContext, &webhookBody)

	log.Println(webhookBody)

	w.Write([]byte("OK"))
}
