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
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{gatorDeckID}", DeleteGatorDeck).Methods("DELETE")
	router.HandleFunc("/decks", CreateDeck).Methods("POST")
	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})
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
	stmt, err := db.Prepare("INSERT INTO decks(gatorDeckName, classcode) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	gatorDeckName := keyVal["gatorDeck_Name"]
	classcode := keyVal["class_code"]
	_, err = stmt.Exec(gatorDeckName, classcode)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New deck was created")
}
func CreateCard(w http.ResponseWriter, r *http.Request) { // new function
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO cards(question, answer) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	question := keyVal["question1"]
	answer := keyVal["answer1"]
	_, err = stmt.Exec(question, answer)
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
	result, err := db.Query("SELECT gatordexID, gatorDeck_Name,"+
		"class_code from decks WHERE gatordexID = ?", params["gatordexID"])
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

func GetCard(w http.ResponseWriter, r *http.Request) { // new deck function
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT idcards, question, answer from cards WHERE idcards = ?", params["idcards"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var card GatorCard
	for result.Next() {
		err := result.Scan(&card.CardID,
			&card.Question, &card.Answer)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(card)
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

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE cards SET idcards = ?," +
		"question= ?, answer=?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idcards := keyVal["IDcards"]
	question := keyVal["Question"]
	answer := keyVal["Answer"]
	_, err = stmt.Exec(idcards, question, answer)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "card with ID = %s was updated",
		params["idcards"])
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
	stmt, err := db.Prepare("DELETE FROM decks WHERE gatordexID = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["gatordexID"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with DeckID = %s was deleted",
		params["gatordexID"])

}
func DeleteGatorCard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM cards WHERE idcards = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["idcards"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "User with DeckID = %s was deleted",
		params["idcards"])

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
	GatorDeckID   string `json:"gatordexID"`
	GatorDeckName string `json:"gatorDeck_Name"`
	ClassCode     string `json:"class_code"`
}

type GatorCard struct {
	CardID   string `json:"idcards"`
	Question string `json:"question1"`
	Answer   string `json:"answer1"`
}

// Db configuration
var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		"root:Gatordex#8867@tcp(127.0.0.1:3306)/userdb")
	//"root:012002Pw0539004*@tcp(127.0.0.1:3306)/userdb")
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
