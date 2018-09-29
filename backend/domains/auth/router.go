package auth

import "github.com/avtal.chalmers.it/backend/domains/util"

func GetRoutes() []util.Route {
	return []util.Route{
		{
			Name:        "UserAccept",
			Method:      "GET",
			Pattern:     "/terms/accept/general",
			HandlerFunc: userAcceptGeneralTerms,
		},
		{
			Name:        "UserAccept",
			Method:      "GET",
			Pattern:     "/terms/accept/hubbit",
			HandlerFunc: userAcceptHubbit,
		},
		{
			Name:        "UserAccept",
			Method:      "GET",
			Pattern:     "/terms/accept/bookit",
			HandlerFunc: userAcceptBookit,
		},
		{
			Name:        "UserAccept",
			Method:      "GET",
			Pattern:     "/terms/accept/chalmersit",
			HandlerFunc: userAcceptChalmersit,
		},
		{
			Name:        "UserStatus",
			Method:      "GET",
			Pattern:     "/terms/user/{uid}",
			HandlerFunc: getUserStatus,
		},
	}
}
