package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type availableArc struct {
	Label       string
	Description string
}

type storyArc struct {
	Title        string
	StoryLines   []string
	StoryOptions []availableArc
}

type story map[string]storyArc

func main() {

	s := readStoryFromJson("gopher.json")

	pageTemplate, err := template.New("Story page").Parse(
		`<html>
			<body>
			<h1>{{.Title}}</h1>
			{{ range .StoryLines}}
			<p> {{ . }}
			{{ end }}
			<ul>
			{{ range .StoryOptions}}
			<li><a href="/{{.Label}}/">{{.Description}}</a></li>
			{{end}}
			</ul>
			</body>
		</html>`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Access to %s\n", r.URL.Path)
		pageTemplate.Execute(w, s["intro"])
	})

	port := 80
	fmt.Printf("Listening on port %d:\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func readStoryFromJson(filename string) story {
	storyJson, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var storyMap map[string]interface{}
	err = json.Unmarshal(storyJson, &storyMap)
	if err != nil {
		log.Fatal(err)
	}

	s := make(map[string]storyArc)
	for key, value := range storyMap {
		s[key] = buildStoryArc(value)
	}

	return s
}

func buildStoryArc(arcUnstructured interface{}) storyArc {
	arcMap := arcUnstructured.(map[string]interface{})
	arcStory := arcMap["story"].([]interface{})
	arcOptions := arcMap["options"].([]interface{})

	arc := storyArc{
		Title:        arcMap["title"].(string),
		StoryLines:   make([]string, len(arcStory)),
		StoryOptions: make([]availableArc, len(arcOptions)),
	}

	for i, line := range arcStory {
		arc.StoryLines[i] = line.(string)
	}

	for i, optionUnstructured := range arcOptions {
		optionMap := optionUnstructured.(map[string]interface{})
		arc.StoryOptions[i] = availableArc{
			Label:       optionMap["arc"].(string),
			Description: optionMap["text"].(string),
		}
	}

	return arc
}
