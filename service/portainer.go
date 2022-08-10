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

	// option 1 - start
	// loginCred := LoginCred{
	// 	Username: config.Portainer.Username,
	// 	Password: config.Portainer.Password,
	// }

	// postBody, _ := json.Marshal(loginCred)
	// log.Println("postBody: ", postBody)
	// requestBody := bytes.NewBuffer(postBody)
	// option 1 - end

	// option 2 - start
	// postBody, _ := json.Marshal([]byte(config.Portainer.ManifestLogin))
	// log.Println("postBody: ", postBody)
	// requestBody := bytes.NewBuffer(postBody)
	// option 2 - end

	// option 3 - start (works)
	// requestBody := bytes.NewBuffer(config.Portainer.ByteManifest)
	// option 3 - end

	// option 4 - start (works)
	var jsonBody map[string]interface{}
	json.Unmarshal([]byte(config.Portainer.ManifestLogin), &jsonBody)
	log.Println("option 4 - jsonBody: ", jsonBody)
	postBody, _ := json.Marshal(jsonBody)
	log.Println("option 4 - postBody: ", postBody)
	requestBody := bytes.NewBuffer(postBody)
	// option 4 - end

	log.Printf("requestBody: %s", requestBody)

	authURL := fmt.Sprintf("%v/auth", config.Portainer.Url)

	res, err := http.Post(authURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// resBody := res.Body
	// body, _ := ioutil.ReadAll(resBody)
	// testBody := string(body)
	// log.Println("BodyString", testBody)

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
