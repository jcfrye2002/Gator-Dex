// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	Routers()
// }

// func Routers() {
// 	InitDB()
// 	defer db.Close()
// 	log.Println("Starting the HTTP server on port 9080")
// 	router := mux.NewRouter()
// 	router.HandleFunc("/users",
// 		GetUsers).Methods("GET")
// 	router.HandleFunc("/users",
// 		CreateUser).Methods("POST")
// 	router.HandleFunc("/users/{id}",
// 		GetUser).Methods("GET")
// 	router.HandleFunc("/users/{id}",
// 		UpdateUser).Methods("PUT")
// 	router.HandleFunc("/users/{id}",
// 		DeleteUser).Methods("DELETE")
// 	http.ListenAndServe(":9080",
// 		&CORSRouterDecorator{router})
// }

// /***************************************************/

// // Get all users
// func GetUsers(w http.ResponseWriter, r *http.Request) {  // not working
// 	w.Header().Set("Content-Type", "application/json")
// 	var users []User
// 	result, err := db.Query("SELECT id, first_name," +
// 		"last_name,email from users")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	for result.Next() {
// 		var user User
// 		err := result.Scan(&user.ID, &user.FirstName,
// 			&user.LastName, &user.Email)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		users = append(users, user)
// 	}
// 	json.NewEncoder(w).Encode(users)
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {  // not working
// 	w.Header().Set("Content-Type", "application/json")
// 	var users []User
// 	result, err := db.Query("SELECT email from users where email = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	for result.Next() {
// 		var user User
// 		err := result.Scan(&user.Password)
// 		if err != nil {
// 			fmt.Fprintf(w, "Wrong Password Please Try Again")
// 			panic(err.Error())
// 		}
// 		users = append(users, user)
// 	}
// 	json.NewEncoder(w).Encode(users)
// }

// // Create user
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	stmt, err := db.Prepare("INSERT INTO users(first_name," +
// 		"last_name,email,password) VALUES(?,?,?,?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	first_name := keyVal["firstName"]
// 	last_name := keyVal["lastName"]
// 	email := keyVal["email"]
// 	password := keyVal["password"]
// 	_, err = stmt.Exec(first_name, last_name, email,password)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "New user was created")
// }

// // Get user by ID
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	result, err := db.Query("SELECT id, first_name,"+
// 		"last_name,email,password from users WHERE id = ?", params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	var user User
// 	for result.Next() {
// 		err := result.Scan(&user.ID, &user.FirstName,
// 			&user.LastName, &user.Email)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

// // Update user
// func UpdateUser(w http.ResponseWriter, r *http.Request) {  // not working
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	stmt, err := db.Prepare("UPDATE users SET first_name = ?," +
// 		"last_name= ?, email=?, password = ? WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	first_name := keyVal["firstName"]
// 	last_name := keyVal["lastName"]
// 	email := keyVal["email"]
// 	password := keyVal["password"]
// 	_, err = stmt.Exec(first_name, last_name, email,password,
// 		params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "User with ID = %s was updated",
// 		params["id"])
// }

// // Delete User
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	_, err = stmt.Exec(params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "User with ID = %s was deleted",
// 		params["id"])
// }

// /***************************************************/

// type User struct {
// 	ID        string `json:"id"`
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// }

// // Db configuration
// var db *sql.DB
// var err error

// func InitDB() {
// 	db, err = sql.Open("mysql",
// 		"root:Gatordex#8867@tcp(127.0.0.1:3306)/userdb")
// 		//"root:012002Pw0539004*@tcp(127.0.0.1:3306)/userdb")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// /***************************************************/

// // CORSRouterDecorator applies CORS headers to a mux.Router
// type CORSRouterDecorator struct {
// 	R *mux.Router
// }

// func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
// 	req *http.Request) {
// 	if origin := req.Header.Get("Origin"); origin != "" {
// 		rw.Header().Set("Access-Control-Allow-Origin", origin)
// 		rw.Header().Set("Access-Control-Allow-Methods",
// 			"POST, GET, OPTIONS, PUT, DELETE")
// 		rw.Header().Set("Access-Control-Allow-Headers",
// 			"Accept, Accept-Language,"+
// 				" Content-Type, YourOwnHeader")
// 	}
// 	// Stop here if its Preflighted OPTIONS request
// 	if req.Method == "OPTIONS" {
// 		return
// 	}

// 	c.R.ServeHTTP(rw, req)
// }

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Routers()
}

func Routers() {
	InitDB()
	defer db.Close()
	log.Println("Starting the HTTP server on port 9080")
	router := mux.NewRouter()
	router.HandleFunc("/users",
		GetUsers).Methods("GET")
	router.HandleFunc("/users",
		CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}",
		GetUser).Methods("GET")
	router.HandleFunc("/users/{id}",
		UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",
		DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{gatorDeckID}",
		DeleteGatorDeck).Methods("DELETE")
	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})
	router.HandleFunc("/decks", CreateDeck).Methods("POST")
}

/***************************************************/

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) { // not working
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT id, first_name," +
		"last_name,email from users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.FirstName,
			&user.LastName, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) { // not working
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT email from users where email = ?")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user User
		err := result.Scan(&user.Password)
		if err != nil {
			fmt.Fprintf(w, "Wrong Password Please Try Again")
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO users(first_name," +
		"last_name,email,password) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	email := keyVal["email"]
	password := keyVal["password"]
	_, err = stmt.Exec(first_name, last_name, email, password)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New user was created")
}

func CreateDeck(w http.ResponseWriter, r *http.Request) { // new function
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO decks(gatorDeck_Name, class_code) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	gatorDeck_Name := keyVal["gatorDeckName"]
	class_code := keyVal["classcode"]
	_, err = stmt.Exec(gatorDeck_Name, class_code)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New deck was created")
}
func CreateCard(w http.ResponseWriter, r *http.Request) { // new function
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO cards(question1, answer1) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	question1 := keyVal["question"]
	answer1 := keyVal["answer"]
	_, err = stmt.Exec(question1, answer1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New Card was created")
}

// Get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, first_name,"+
		"last_name,email,password from users WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.ID, &user.FirstName,
			&user.LastName, &user.Email)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}

func GetDeck(w http.ResponseWriter, r *http.Request) { // new deck function
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT gatorDeck_ID, gatorDeck_Name,"+
		"class_code from decks WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var deck GatorDeck
	for result.Next() {
		err := result.Scan(&deck.GatorDeckID,
			&deck.GatorDeckName, &deck.ClassCode)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(deck)
}

// Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE users SET first_name = ?," +
		"last_name= ?, email=?, password = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	email := keyVal["email"]
	password := keyVal["password"]
	_, err = stmt.Exec(first_name, last_name, email, password,
		params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was updated",
		params["id"])
}

// Delete User
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with ID = %s was deleted",
		params["id"])
}

func DeleteGatorDeck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM decks WHERE gatorDeckID = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["gatorDeckID"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with DeckID = %s was deleted",
		params["gatorDeckID"])

}

/***************************************************/

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type GatorDeck struct {
	GatorDeckID   string `json:"gatorDeckID"`
	GatorDeckName string `json:"gatorDeckName"`
	ClassCode     string `json:"classcode"`
}

type GatorCard struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// Db configuration
var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		//"root:Gatordex#8867@tcp(127.0.0.1:3306)/userdb")
		"root:012002Pw0539004*@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		panic(err.Error())
	}
}

/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
