package main

//https://tutorialedge.net/golang/creating-restful-api-with-golang/
//
// docker run --name some-cassandra --network some-network -e
// CASSANDRA_BROADCAST_ADDRESS=127.0.0.1 -p 9042:9042 -d cassandra:latest

//https://proselyte.net/tutorials/cassandra/tables/
//CREATE KEYSPACE restsoft WITH replication = {'class': 'SimpleStrategy', 'replication_factor':1};
//CREATE TABLE users (user_id UUID PRIMARY KEY, name text);
//cqlsh 127.0.0.1 9042


import (
	. "REST_soft/StructUser"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//HTTP requests
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/v1/users", returnAllUsers).Methods("GET")
	myRouter.HandleFunc("/api/v1/users/{id}", returnSingleUser).Methods("GET")
	myRouter.HandleFunc("/api/v1/post/users", createNewUser).Methods("POST")
	myRouter.HandleFunc("/api/v1/put/users/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/api/v1/del/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

//Function to create connection to Cassandra
func getCassandraSession() *gocql.Session{
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "restsoft"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return session
}

//Function return all users from Users
//method GET
func returnAllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllUsers")
	var user_id gocql.UUID
	var name string
	var Users []User
	var usr User

	session := getCassandraSession()
	defer session.Close()

	iter := session.Query(`SELECT user_id, name FROM users`).Iter()
	for iter.Scan(&user_id, &name) {
		fmt.Println("Tweet:", user_id, name)
		usr.Id = user_id
		usr.Name = name
		Users = append(Users, usr)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(Users)
}

//Function return user from Users by {Id}
//method GET
func returnSingleUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var usr User

	session := getCassandraSession()
	defer session.Close()

	if err := session.Query("SELECT user_id, name FROM users WHERE user_id = ?",
		id).Scan(&usr.Id, &usr.Name); err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(usr)
}

//Function create new user in Users
//method POST
func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	session := getCassandraSession()
	if err := session.Query(`INSERT INTO users (user_id, name) VALUES (?, ?)`,
		gocql.TimeUUID(), user.Name).Exec(); err != nil {
		log.Fatal(err)
	}
	defer session.Close()
}

//Function delete user from Users by {Id}
//method DELETE
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	session := getCassandraSession()
	defer session.Close()

	if err := session.Query("DELETE FROM users WHERE user_id = ?",
		id).Exec(); err != nil {
		fmt.Println(err)
	}
}

//Function update user in Users by {Id}
//method PUT
func updateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedEvent User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	session := getCassandraSession()
	defer session.Close()

	if err := session.Query("UPDATE users SET name = ? WHERE user_id = ?",
		updatedEvent.Name, id).Exec(); err != nil {
		fmt.Println(err)
	}
	//json.NewEncoder(w).Encode(user)
}

var Users []User

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}