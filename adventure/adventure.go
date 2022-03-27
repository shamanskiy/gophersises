package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type storyArc struct {
	Title      string
	StoryLines []string
}

func main() {

	story := readStoryFromJson("gopher.json")

	intro := story["intro"].(map[string]interface{})
	introArc := storyArc{
		Title:      intro["title"].(string),
		StoryLines: make([]string, len(intro["story"].([]interface{}))),
	}
	for i, line := range intro["story"].([]interface{}) {
		introArc.StoryLines[i] = line.(string)
	}

	pageTemplate, err := template.New("Story page").Parse(
		`<html>
			<body>
			<h1>{{.Title}}</h1>
			{{ range .StoryLines}}
			<p> {{ . }}
			{{ end }}
			</body>
		</html>`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageTemplate.Execute(w, introArc)
	})
	http.ListenAndServe(":80", nil)
}

func readStoryFromJson(filename string) map[string]interface{} {
	storyJson, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var story map[string]interface{}
	err = json.Unmarshal(storyJson, &story)
	if err != nil {
		log.Fatal(err)
	}

	return story
}
