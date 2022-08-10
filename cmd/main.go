package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gradity/go-portainer-action/service"
)

func main() {
	jwt := service.GetJWTToken()
	log.Println(*jwt)
	lp := jwt.ListEndpoints()
	log.Println("ListEndpoint: ", *lp)

	lpMarshal, _ := json.Marshal(*lp)
	log.Println("lpMarshal: ", lpMarshal)
	lpString := string(lpMarshal)
	log.Println("lpString: ", lpString)

	// use this as action set-output
	lpStringEscaped := fmt.Sprintf("%q", lpString)

	log.Println("lpStringEscaped: ", lpStringEscaped)

	// var lpUnmarshal []map[string]interface{}
	// _ = json.Unmarshal([]byte(lpString), &lpUnmarshal)
	// log.Println("lpUnmarshal: ", lpUnmarshal)
}
