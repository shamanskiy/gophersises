package main

import (
	"encoding/json"
	"flag"
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
	port := flag.Int("port", 80, "Port to serve the adventure at")
	flag.Parse()

	pageHtml, err := readPageHtmlTemplate("storyPage.html.template")
	checkError(err)

	pageTemplate := template.Must(template.New("Story page").Parse(pageHtml))

	story, err := readStoryFromJson("gopher.json")
	checkError(err)

	mux := http.NewServeMux()
	mux.HandleFunc("/", storyHandler(pageTemplate, story))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Listening on port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func storyHandler(pageTemplate *template.Template, story Story) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Access to %s\n", r.URL.Path)

		chapter := strings.ReplaceAll(r.URL.Path, "/", "")
		if _, ok := story[chapter]; !ok {
			chapter = "intro"
		}
		err := pageTemplate.Execute(w, story[chapter])
		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readPageHtmlTemplate(filename string) (string, error) {
	pageHtml, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(pageHtml), nil
}

func readStoryFromJson(filename string) (Story, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
