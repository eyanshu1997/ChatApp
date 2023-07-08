package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eyanshu1997/ChatApp/config"
)

func Registeration(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Registration Server Endpoint!")
	fmt.Println("Endpoint Hit: Registration")
}

func handleRequests(host string, port int) error {
	fmt.Printf("Starting server %s:%d", host, port)
	http.HandleFunc("/", Registeration)
	return http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
}

func main() {
	rcfg := config.RegistrationConfig{Host: "localhost", Port: 9000}
	rcfg.ReadConfigFile("filename")
	err := handleRequests(rcfg.Host, rcfg.Port)
	if err != nil {
		fmt.Println("unable to start server ", err)
	}
}
