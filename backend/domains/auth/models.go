package auth

type tokenBody struct {
	AccessToken string `json:"access_token"`
}

type terms map[string]bool

type user struct {
	AcceptedTerms terms  `json:"acceptedTerms"`
	Uid           string `json:"uid"`
	Mail          string `json:"mail"`
	Nick          string `json:"nickname"`
}

const (
	BookitTerms         string = "bookit_terms"
	ChalmersitTerms     string = "chalmersit_terms"
	GeneralAccountTerms string = "general_account_terms"
	HubbitTerms         string = "hubbit_terms"
)
