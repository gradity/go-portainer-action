package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gradity/go-portainer-action/config"
)

type LoginCred struct {
	Username string
	Password string
}

type AuthResponse struct {
	Jwt string `json:"jwt"`
}

type ListEndpointResponse struct {
	Id   int
	Name string
}

func GetJWTToken() *string {

	postBody, _ := json.Marshal(LoginCred{
		Username: config.Portainer.Username,
		Password: config.Portainer.Password,
	})

	requestBody := bytes.NewBuffer(postBody)

	authURL := fmt.Sprintf("%v/auth", config.Portainer.Url)

	res, err := http.Post(authURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var response AuthResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err)
	}

	return &response.Jwt
}

func ListEndpoints(token string) *[]ListEndpointResponse {

	endpointURL := fmt.Sprintf("%v/endpoints", config.Portainer.Url)

	req, err := http.NewRequest(http.MethodGet, endpointURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var response []ListEndpointResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err)
	}

	return &response
}
