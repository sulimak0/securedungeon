package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sulimak0/securedungeon/internal/models"
	"github.com/sulimak0/securedungeon/internal/utils"
	"net/http"
)

type AlarmHandler struct {
	router   *mux.Router
	messages chan models.Alarm
}

func NewAlarmHandler(router *mux.Router, messages chan models.Alarm) {
	handler := &AlarmHandler{router: router, messages: messages}
	router.HandleFunc("/api/alarms", handler.Fetch).Methods("POST")
}

func (a *AlarmHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	alarm := &models.Alarm{}
	err := json.NewDecoder(r.Body).Decode(alarm)
	if err != nil {
		utils.Respond(w, utils.Messge(false, "Invalid request"))
		return
	}
	if len(alarm.Message) == 0 {
		alarm.Message = "unknown alarm from gw"
	}
	a.messages <- *alarm
}
