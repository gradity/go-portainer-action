package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gradity/go-portainer-action/config"
)

type LoginCred struct {
	Username string
	Password string
}

type JWTResponse struct {
	Jwt string `json:"jwt"`
}

type ListEndpointResponse struct {
	Id   int
	Name string
}

func GetJWTToken() *JWTResponse {

	loginCred := LoginCred{
		Username: config.Portainer.Username,
		Password: config.Portainer.Password,
	}
	postBody, _ := json.Marshal(loginCred)
	log.Println("postBody: ", postBody)

	requestBody := bytes.NewBuffer(postBody)

	log.Printf("requestBody: %T", requestBody)

	authURL := fmt.Sprintf("%v/auth", config.Portainer.Url)

	res, err := http.Post(authURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var response JWTResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err)
	}

	return &response
}

func (jwt *JWTResponse) ListEndpoints() *[]ListEndpointResponse {

	endpointURL := fmt.Sprintf("%v/endpoints", config.Portainer.Url)

	req, err := http.NewRequest(http.MethodGet, endpointURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", jwt.Jwt))

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
