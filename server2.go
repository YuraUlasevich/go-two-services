package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// UsersContainer contains the list of users.
type UsersContainer struct {
	Users []User `json:"users"`
}

type User struct {
	ID   int
	Name string
}

func (u *User) IsEmpty() bool {
	return u.ID == 0
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users UsersContainer
	jsonFile, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &users); err != nil {
		// TODO: handle error
	}
	if err != nil {
		// TODO: handle error
	}

	result := users
	if err := json.NewEncoder(w).Encode(result); err != nil {
		// TODO: handle error
	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	var users UsersContainer
	if param != "" {
		jsonFile, err := os.Open("data.json")
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		if err := json.Unmarshal(byteValue, &users); err != nil {
			// TODO: handle error
		}

		var result User
		lookup, err := strconv.ParseInt(param, 10, 32)

		if err != nil {
			// TODO: handle error
		}
		for _, u := range users.Users {
			if int64(u.ID) == lookup {
				result = u
			}
		}

		if result.IsEmpty() {
			w.Write([]byte("{}"))
			return
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			// TODO: handle error
		}
	}
}

func main() {
	http.HandleFunc("/getusers", getUsers)
	http.HandleFunc("/getuser", getUser)

	fmt.Println("Starting server at :8081")
	http.ListenAndServe(":8081", nil)
}
