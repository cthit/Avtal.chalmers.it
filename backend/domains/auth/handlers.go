package auth

import (
	"encoding/json"
	"net/http"

	"github.com/avtal.chalmers.it/backend/config"
	"github.com/gorilla/mux"
)

func userAcceptGeneralTerms(w http.ResponseWriter, r *http.Request) {
	userAcceptTerms(
		r,
		terms{GeneralAccountTerms: true},
		urlConf{
			clientID:     config.BetaAppId_Accounts,
			redirect:     "http://localhost:3000/terms/accept/general",
			clientSecret: config.BetaAppSecrets.Accounts,
		},
	)
}

func userAcceptHubbit(w http.ResponseWriter, r *http.Request) {
	userAcceptTerms(
		r,
		terms{HubbitTerms: true},
		urlConf{
			clientID:     config.BetaAppId_Hubbit,
			redirect:     "http://localhost:3000/terms/accept/hubbit",
			clientSecret: config.BetaAppSecrets.Hubbit,
		},
	)
}

func userAcceptBookit(w http.ResponseWriter, r *http.Request) {
	userAcceptTerms(
		r,
		terms{BookitTerms: true},
		urlConf{
			clientID:     config.BetaAppId_Bookit,
			redirect:     "http://localhost:3000/terms/accept/bookit",
			clientSecret: config.BetaAppSecrets.Bookit,
		},
	)
}

func userAcceptChalmersit(w http.ResponseWriter, r *http.Request) {
	userAcceptTerms(
		r,
		terms{ChalmersitTerms: true},
		urlConf{
			clientID:     config.BetaAppId_Chalmersit,
			redirect:     "http://localhost:3000/terms/accept/chalmersit",
			clientSecret: config.BetaAppSecrets.Chalmersit,
		},
	)
}

func getUserStatus(w http.ResponseWriter, r *http.Request) {
	uid := mux.Vars(r)["uid"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	terms := getUser(uid).AcceptedTerms
	userjson, _ := json.Marshal(terms)
	w.Write(userjson)
}
