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

type availableArc struct {
	Label       string
	Description string
}

type storyArc struct {
	Title        string
	StoryLines   []string
	StoryOptions []availableArc
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

	port := 80
	fmt.Printf("Listening on port %d:\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func storyHandler(pageTemplate *template.Template, story map[string]storyArc) func(w http.ResponseWriter, r *http.Request) {
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

func readStoryFromJson(filename string) map[string]storyArc {
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
