package main

import (
	"encoding/json"
	"net/http"
	"os"
	"portfolio_website/templates"
)

func main() {
	data, _ := os.ReadFile("content/hero.json")
	var heroData templates.HeroData
	json.Unmarshal(data, &heroData)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.Index(heroData).Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
