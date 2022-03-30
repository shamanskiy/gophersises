package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Story map[string]Chapter

type StoryOption struct {
	Chapter     string `json:"arc"`
	Description string `json:"text"`
}

type Chapter struct {
	Title      string        `json:"title"`
	Paragraphs []string      `json:"story"`
	Options    []StoryOption `json:"options"`
}

func main() {
	pageHtml := readPageHtmlTemplate("storyPage.html.template")
	pageTemplate, err := template.New("Story page").Parse(pageHtml)
	if err != nil {
		log.Fatal(err)
	}

	story := readStoryFromJson("gopher.json")

	mux := http.NewServeMux()
	mux.HandleFunc("/", storyHandler(pageTemplate, story))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := 80
	fmt.Printf("Listening on port %d:\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func storyHandler(pageTemplate *template.Template, story Story) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Access to %s\n", r.URL.Path)

		arcLabel := strings.ReplaceAll(r.URL.Path, "/", "")
		if _, ok := story[arcLabel]; !ok {
			arcLabel = "intro"
		}
		pageTemplate.Execute(w, story[arcLabel])
	}
}

func readPageHtmlTemplate(filename string) string {
	pageHtml, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(pageHtml)
}

func readStoryFromJson(filename string) Story {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		log.Fatal(err)
	}

	return story
}
