package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func (a *App) frontendRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style=\"font-family: sans-serif\">Little Blue</h1>"))
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

	go webhookBody.ProcessWebhookEvent(a)

	log.Println(webhookBody)

	w.Write([]byte("OK"))
}
