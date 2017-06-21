package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Handlers ...
// Server is the reverse-proxy/load-balancer engine
type Handlers struct {
	globalConfiguration GlobalConfiguration
}

// NewHandlers ...
func NewHandlers(globalConfiguration *GlobalConfiguration) *Handlers {
	handlers := new(Handlers)
	handlers.globalConfiguration = *globalConfiguration

	return handlers
}

// HealthPage ...
func (handlers *Handlers) HealthPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	fmt.Println("Endpoint Hit: healthPage")
}

// DockerWebhook ...
func (handlers *Handlers) DockerWebhook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t DockerRelease

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("%s/%s:%s", t.Repository.Namespace, t.Repository.Name, t.PushedData.Tag)

	globalConfiguration := handlers.globalConfiguration

	fmt.Println(time.Now().String(), " : ", message)

	if globalConfiguration.Slack != nil {
		SendSlackMessage(&globalConfiguration, message)
	}

}
