package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var tpl *template.Template
var db *sql.DB
var store = sessions.NewCookieStore([]byte("super-secret"))

/*
	type GatorDex struct {
		quetion string
		answer  string
	}
*/
type User struct {
	ID       string
	username string
	password string
}

func main() {
	tpl, _ = template.ParseGlob("Templates/*.html")
	var err error
	db, err = sql.Open("mysql", "root:012002Pw0539004*@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/createGatorDex", createHandler)
	http.HandleFunc("/createGatorDexauth", createHandlerauth)
	http.HandleFunc("/profile", profileHandlerAuth)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.ListenAndServe("localhost:8080", context.ClearHandler(http.DefaultServeMux))
}

// registerHandler serves form for registring new users
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****RegisterHandler Running*****")
	tpl.ExecuteTemplate(w, "register.html", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****LoginPage Running*****")
	tpl.ExecuteTemplate(w, "login.html", nil)
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****Create GatorDex Running*****")
	tpl.ExecuteTemplate(w, "createGatorDex.html", nil)
}
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****LoginAuthHandle Running*****")
	r.ParseForm()
	var u User
	u.username = r.FormValue("username")
	u.password = r.FormValue("password")
	var userID string

	fmt.Println("username:", u.username, "password:", u.password)

	stmt := "SELECT username FROM students WHERE username = ?"
	row := db.QueryRow(stmt, u.username)
	err := row.Scan(&u.password)

	if err != nil {
		fmt.Println("error selecting Password in db by Username")
		tpl.ExecuteTemplate(w, "login.html", "check username and password")
		return
	}

	if err == nil {
		session, _ := store.Get(r, "session")
		// session struct has field Values map[interface{}]interface{}
		session.Values["userID"] = userID
		// save before writing to response/return from handler
		session.Save(r, w)
		tpl.ExecuteTemplate(w, "profile.html", u)
		return
	}

	fmt.Println("incorrect password")
	tpl.ExecuteTemplate(w, "login.html", "check username and password")

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****logoutHandler running*****")
	session, _ := store.Get(r, "session")
	// The delete built-in function deletes the element with the specified key (m[key]) from the map.
	// If m is nil or there is no such element, delete is a no-op.
	delete(session.Values, "userID")
	session.Save(r, w)
	tpl.ExecuteTemplate(w, "login.html", "Logged Out")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****aboutHandler running*****")
	session, _ := store.Get(r, "session")
	_, ok := session.Values["userID"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound) // http.StatusFound is 302
		return
	}
	tpl.ExecuteTemplate(w, "about.html", "Logged In")
}

// registerAuthHandler creates new user in database
func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	/*
		1. check username criteria
		2. check password criteria
		3. check if username is already exists in database
		4. create bcrypt hash from password
		5. insert username and password hash in database
		(email validation will be in another video)
	*/
	fmt.Println("*****registerAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	// check username for only alphaNumeric characters
	var nameAlphaNumeric = true
	for _, char := range username {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	// check username length
	var nameLength bool
	if 5 <= len(username) && len(username) <= 50 {
		nameLength = true
	}
	// check password criteria
	password := r.FormValue("password")
	fmt.Println("password:", password, "\npswdLength:", len(password))
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	}
	// check if username already exists for availability
	stmt := "SELECT username FROM students WHERE username = ?"
	fmt.Println(username)
	row := db.QueryRow(stmt, username)
	var uID string
	err := row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		tpl.ExecuteTemplate(w, "register.html", "username already taken")
		return
	}

	db, err := sql.Open("mysql", "root:012002Pw0539004*@tcp(localhost:3306)/testdb")

	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	insert, err := db.Query("INSERT INTO `testdb`.`students` (`username`, `password`) VALUES (?,?)", (username), (password))
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	if err != nil {
		fmt.Println("error inserting new user")
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}

	fmt.Println("Successful Connection to Database!")
	fmt.Fprint(w, "congrats, your account has been successfully created")
}

func createHandlerauth(w http.ResponseWriter, r *http.Request) {
	/*
		1. check username criteria
		2. check password criteria
		3. check if username is already exists in database
		4. create bcrypt hash from password
		5. insert username and password hash in database
		(email validation will be in another video)
	*/
	fmt.Println("*****CreateGatorDexAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	question := r.FormValue("question")
	// check username length
	var nameLength bool
	if 1 <= len(question) {
		nameLength = true
		print(nameLength)
	}
	// check password criteria
	answer := r.FormValue("answer")
	fmt.Println("answer:", answer, "\nanswerLength:", len(answer))

	// check if username already exists for availability

	db, err := sql.Open("mysql", "root:012002Pw0539004*@tcp(localhost:3306)/testdb")

	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	insert, err := db.Query("INSERT INTO `testdb`.`gatordex` (`username`, `question`, `answer`) VALUES (?,?,?)", (username), (question), (answer))
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	if err != nil {
		fmt.Println("error inserting new GatorDex")
		tpl.ExecuteTemplate(w, "register.html", "there was a problem creating GatorDex")
		return
	}

	fmt.Println("Successful Connection to Database!")
	fmt.Fprint(w, "congrats, your Gator Dex has been successfully created")
}
func profileHandlerAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****ProfileHandler running*****")
	var User User
	session, _ := store.Get(r, "session")
	User.ID, _ = session.Values["userID"].(string)
	/*err := User.SelectByID()
	if err != nil {
		td := map[string]string{
			"UserMessage": "There was an issue displaying profile information",
		}
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}*/
	fmt.Println("User:", User)
	tpl.ExecuteTemplate(w, "profile.html", User)
}

func (u *User) SelectByID() error {
	stmt := "SELECT username FROM students WHERE username = ?"
	row := db.QueryRow(stmt, u.username)
	err := row.Scan(&u.password)
	if err != nil {
		fmt.Println("error selecting user by id, err:", err)
		return err
	}
	return err
}
