// go-login project main.go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("indexHandler...")
	tmpl, err := template.ParseFiles("index.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

//doLogin
func doLoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	log.Print("loginHandler...  " + name + "  " + password)
	if name == "admin" && password == "123456" {
		tmpl, err := template.ParseFiles("submit.html")
		checkErr(err)
		err = tmpl.Execute(w, nil)
		checkErr(err)
	} else {
		tmpl, err := template.ParseFiles("logined.html")
		checkErr(err)
		err = tmpl.Execute(w, nil)
		checkErr(err)
	}
}

//register
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("register.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

//doRegister
func doRegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("registered.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

//sumbit
func submitHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("submit.html")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
}

func Register() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/doLogin.do", doLoginHandler)
	http.HandleFunc("/register.do", registerHandler)
	http.HandleFunc("/doRegister.do", doRegisterHandler)
	http.HandleFunc("/submit.do", submitHandler)
}

func main() {
	log.Println("start...")
	Register()
	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		log.Println("ListenAndServe", err.Error())
	}
}
