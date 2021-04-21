package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:8081/getusers"
	resp, err := http.Get(url)
	if err != nil {
		// TODO: handle error
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	//	fmt.Printf("%v", string(respBody))
	fmt.Fprintln(w, string(respBody))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	url := "http://127.0.0.1:8081/getuser?id=" + param
	resp, err := http.Get(url)
	if err != nil {
		// TODO: handle error
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	//      fmt.Printf("%v", string(respBody))
	fmt.Fprintln(w, string(respBody))
}

func main() {
	http.HandleFunc("/getusers", getUsers)
	http.HandleFunc("/getuser", getUser)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
