package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type urlConf struct {
	clientID     string
	redirect     string
	clientSecret string
}

func userAcceptTerms(r *http.Request, termsDiff terms, conf urlConf) {
	code := r.URL.Query().Get("code")
	usr := authenticateAndExtractUserInformation(code, conf)

	if usr.Uid != "" {
		usr.AcceptedTerms = termsDiff
		upsertUser(usr)
	}
}

func authenticateAndExtractUserInformation(betaCode string, conf urlConf) user {
	var usr user
	response, err := http.PostForm("https://beta-account.chalmers.it/oauth/token", url.Values{
		"client_id":     {conf.clientID},
		"redirect_uri":  {conf.redirect},
		"grant_type":    {"authorization_code"},
		"code":          {betaCode},
		"client_secret": {conf.clientSecret},
	})

	if err != nil {
		fmt.Println("Err", err)
		return usr
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var raw tokenBody
	json.Unmarshal(body, &raw)
	//fmt.Println("AccessToken:", raw.AccessToken)

	payload := url.Values{}
	payload.Add("token", raw.AccessToken)
	req, err := http.NewRequest("GET", "https://beta-account.chalmers.it/me.json?"+payload.Encode(), nil)
	req.Header.Add("authorization", "Bearer "+raw.AccessToken)

	client := http.Client{}
	resp, err := client.Do(req)

	//fmt.Println(resp)

	res, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(res))

	json.Unmarshal(res, &usr)
	return usr
}
