package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/webhook"
)

func (a *App) webhookUpdate(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var webhookBody webhook.GithubWebhookRequest
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
