package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Conversion struct {
	Amount float32
	From   string
	To     string
}

var tmpl *template.Template

func init() {
	var err error

	tmpl, err = template.ParseGlob("./templates/*.html")

	if err != nil {
		panic(err)
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/length", length).Methods("GET")
	router.HandleFunc("/weight", weight).Methods("GET")
	router.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "temperature", nil)
	}).Methods("GET")

	router.HandleFunc("/length", func(w http.ResponseWriter, r *http.Request) {
		amountStr := r.FormValue("amount")
		from := r.FormValue("from")
		to := r.FormValue("to")

		amount, _ := strconv.ParseFloat(amountStr, 32)

		conversion := Conversion{
			float32(amount), from, to,
		}

		fmt.Println(conversion)

	}).Methods("POST")
	// // r := chi.NewRouter()
	// // r.Use(middleware.Logger)
	// component := _templ.Header()
	// // r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// // 	w.Write([]byte(component.Render())
	// // })

	// http.Handle("/", templ.Handler(component))

	// // component.Render(context.Background(), os.Stdout)

	// http.ListenAndServe(":3000", nil)

	http.ListenAndServe(":3000", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func length(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "length", nil)
}

func weight(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "weight", nil)
}
