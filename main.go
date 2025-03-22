package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/sajari/word2vec"
)

type App struct {
	Model     *word2vec.Model
	Templates *template.Template
	ModelName string
}

var app App

func main() {
	binaryPath := flag.String("b", "", "Path to pre-trained word embedding binary (word2vec format)")

	flag.Parse()

	if *binaryPath == "" {
		panic("Missing binary, specify it with '-b <binary_path>'.")
	}

	fmt.Println("Loading pre-trained binary... ⏳")

	binaryFile, err := os.Open(*binaryPath)
	if err != nil {
		panic(err)
	}

	defer binaryFile.Close()

	app.Model, err = word2vec.FromReader(binaryFile)
	if err != nil {
		panic(err)
	}

	app.ModelName = GetFileName(binaryFile.Name())

	fmt.Println("Binary loaded ✅")

	fmt.Println("Loading template(s)... ⏳")
	app.Templates = template.Must(template.ParseFiles("index.html", "templates/_tpl_closest.html", "templates/_tpl_similarity.html"))
	fmt.Println("Template(s) loaded... ✅")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/similar", SimilarityHandler)
	http.HandleFunc("/closest", ClosestWordsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":3333", nil)
}
