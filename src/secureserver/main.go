package main

import (
	"github.com/gorilla/mux"
	"github.com/sulimak0/securedungeon/internal/backend"
	"github.com/sulimak0/securedungeon/internal/bot"
	"github.com/sulimak0/securedungeon/internal/controllers"
	"github.com/sulimak0/securedungeon/internal/models"
	"log"
	"net/http"
	"time"
)

const (
	token   = ""
	offset  = 0
	timeout = 60
	port    = "9944"
	portgw  = "9943"
	url     = "http://localhost:"
	userID  = 869847693
)

func main() {
	a := make(chan models.Alarm, 1)
	s := backend.NewSecuritySystem(url+portgw, a)
	b := bot.NewBot(token, offset, timeout, s, []int64{userID})
	go b.CommandProcessor()

	go func() {
		for {
			err := b.Ping()
			if err != nil {
				b.Notify(*models.New("Alarm! No response from camera gw!"))
			}
			log.Println("Succesfull ping")
			time.Sleep(time.Second * 30)
		}

	}()
	go func() {
		for {
			b.GWMessagesProcessor()
		}

	}()
	log.Println("Starting application on port: ", port)

	router := mux.NewRouter()
	controllers.NewAlarmHandler(router, a)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Panic(err)
	}
}
