package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var code string
var TokenR TokenResp

// TOKEN
type Token struct {
	Grant_type    string `json:"grant_type"`
	Client_id     int    `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Code          string `json:"code"`
	Redirect_uri  string `json:"redirect_uri"`
}

type TokenResp struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Expires_in    int    `json:"expires_in"`
	Scope         string `json:"scope"`
	User_id       int    `json:"user_id"`
	Refresh_token string `json:"refresh_token"`
}

// FUNCIONES PARA INTERCAMBIAR EL CODE POR UN ACCESS TOKEN
func GetToken(c *gin.Context) {
	code = c.Query("code")
	TokenRequest(code, c)
}

func TokenRequest(code string, c *gin.Context) {
	token := Token{
		Grant_type:    "authorization_code",
		Client_id:     6589031130474375,
		Client_secret: "HLa9vLNaV5w4SEqwUVUcFtwgV5ZbWywf",
		Code:          code,
		Redirect_uri:  "http://localhost:8080/auth/code/",
	}

	jsonToken, err := json.Marshal(token)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonToken))

	// Intercambiamos code por token
	resp, err := http.Post("https://api.mercadolibre.com/oauth/token", "application/json; application/x-www-form-urlencoded", bytes.NewBuffer(jsonToken))

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	bodyString := string(data)
	fmt.Println(bodyString)

	json.Unmarshal(data, &TokenR)
	fmt.Printf("%+v\n", TokenR)

	c.JSON(200, TokenR)
}
