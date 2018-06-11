package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func (a *App) frontendRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style=\"font-family: sans-serif\">Ricky</h1>"))
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

	newJob := BuildContext{

		// TODO: Have GithubWebhookRequest implement interface to build a GithubRepository
		Source: GitRepository{
			FullName: webhookBody.Repository.FullName,
			DashedName: repoDashedName,
			RepoName: webhookBody.Repository.Name,
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
	}

	a.Jobs[repoDashedName] = newJob

	newJob.addMessage(MSG_LEVEL_DEBUG, "Event received")

	go a.ProcessWebhook(&newJob, &webhookBody)

	log.Println(webhookBody)

	w.Write([]byte("OK"))
}
