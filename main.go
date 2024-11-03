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

type Result struct {
	Old  float32
	New  float32
	From string
	To   string
}

var Lengths = map[string]float32{
	"mm":  1000.00,
	"cm":  100.00,
	"dm":  10.00,
	"m":   1.00,
	"dam": 0.1,
	"hm":  0.01,
	"km":  0.001,
}

var Weights = map[string]float32{
	"mg": 0.1,
	"g":  1,
	"cg": 10,
	"dg": 100,
	"kg": 1000,
}

var Temperatures = []string{"C", "F"}

// var tmpl *template.Template
var tmpl = template.Must(template.ParseFiles("./templates/index.html", "./templates/length.html"))

func init() {
	var err error

	tmpl, err = template.ParseGlob("./templates/*.html")

	if err != nil {
		panic(err)
	}
}

func convertLength(c Conversion) Result {

	from := Lengths[c.From]
	to := Lengths[c.To]

	ratio := to / from

	return Result{c.Amount, c.Amount * ratio, c.From, c.To}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/length", length).Methods("GET")
	router.HandleFunc("/weight", weight).Methods("GET")
	router.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "temperature", Temperatures)
	}).Methods("GET")

	router.HandleFunc("/length", func(w http.ResponseWriter, r *http.Request) {
		amountStr := r.FormValue("amount")
		from := r.FormValue("from")
		to := r.FormValue("to")

		amount, _ := strconv.ParseFloat(amountStr, 32)

		toConvert := Conversion{
			float32(amount), from, to,
		}

		fmt.Println(toConvert)

		tmpl.ExecuteTemplate(w, "result", convertLength(toConvert))

	}).Methods("POST")

	router.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		amountStr := r.FormValue("amount")
		from := r.FormValue("from")
		to := r.FormValue("to")

		amount, _ := strconv.ParseFloat(amountStr, 32)

		var result Result

		result.Old = float32(amount)

		if from == "C" {
			result.From = "C"
			if to == "F" {
				result.To = "F"
				result.New = float32((amount * 1.8) + 32)
			}
		} else if from == "F" {
			result.From = "F"
			if to == "C" {
				result.To = "C"
				result.New = float32((amount - 32) / 1.8)
			}
		}

		tmpl.ExecuteTemplate(w, "result", result)
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
	tmpl.ExecuteTemplate(w, "index.html", Lengths)
}

func length(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "length", Lengths)
}

func weight(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "weight", Weights)
}
