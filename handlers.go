package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sajari/word2vec"
)

func ClosestWordsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	count, _ := strconv.ParseInt(r.PostFormValue("count"), 10, 32)
	matches, err := app.Model.CosN(word2vec.Expr{r.PostFormValue("word"): 1}, int(count))
	if err != nil {
		fmt.Fprintf(w, "This word is not recognized by the model")
		return
	}

	for i := range matches {
		matches[i].Score *= 100
	}

	app.Templates.ExecuteTemplate(w, "_tpl_closest.html", matches)
}

func SimilarityHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	word1 := r.PostFormValue("word1")
	word2 := r.PostFormValue("word2")

	similarity, err := app.Model.Cos(word2vec.Expr{word1: 1}, word2vec.Expr{word2: 1})
	if err != nil {
		fmt.Fprintf(w, "At least one word is not recognized by the model")
		return
	}

	app.Templates.ExecuteTemplate(w, "_tpl_similarity.html", similarity*100)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	app.Templates.ExecuteTemplate(w, "index.html", app.ModelName)
}
