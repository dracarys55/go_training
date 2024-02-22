package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users []User
var nextID int = 1

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/users", getUsers).Methods("GET")
    router.HandleFunc("/user", createUser).Methods("POST")
    router.HandleFunc("/user/{id}", getUser).Methods("GET")
    router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
    router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

    fmt.Println("Starting server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var newUser User
    json.NewDecoder(r.Body).Decode(&newUser)
    newUser.ID = nextID
    nextID++
    users = append(users, newUser)
    json.NewEncoder(w).Encode(newUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, user := range users {
        if user.ID == id {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    http.NotFound(w, r)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for i, user := range users {
        if user.ID == id {
            users[i].Name = r.FormValue("name")
            json.NewEncoder(w).Encode(users[i])
            return
        }
    }
    http.NotFound(w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            fmt.Fprintf(w, "User with ID %v deleted successfully", id)
            return
        }
    }
    http.NotFound(w, r)
}
