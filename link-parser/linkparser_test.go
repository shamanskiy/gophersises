package linkparser

import (
	"os"
	"testing"
)

func TestExtractLinks_NakedA(t *testing.T) {
	html := `<a href="/link">Link text</a>`

	got, err := ExtractLinks([]byte(html))
	want := map[string]string{
		"/link": "Link text",
	}

	checkError(err, t)
	checkLinkMaps(got, want, t)
}

func TestExtractLinks_Ex1(t *testing.T) {

	html, err := os.ReadFile("testdata/ex1.html")
	checkError(err, t)

	got, err := ExtractLinks(html)
	want := map[string]string{
		"/other-page": "A link to another page",
	}

	checkError(err, t)
	checkLinkMaps(got, want, t)
}

func TestExtractLinks_Ex2(t *testing.T) {
	html, err := os.ReadFile("testdata/ex2.html")
	checkError(err, t)

	got, err := ExtractLinks([]byte(html))
	want := map[string]string{
		"https://www.twitter.com/joncalhoun": "Check me out on twitter",
		"https://github.com/gophercises":     "Gophercises is on Github!",
	}

	checkError(err, t)
	checkLinkMaps(got, want, t)
}

func TestExtractLinks_Ex3(t *testing.T) {
	html, err := os.ReadFile("testdata/ex3.html")
	checkError(err, t)

	got, err := ExtractLinks([]byte(html))
	want := map[string]string{
		"#":                                "Login",
		"/lost":                            "Lost? Need help?",
		"https://twitter.com/marcusolsson": "@marcusolsson",
	}

	checkError(err, t)
	checkLinkMaps(got, want, t)
}

func TestExtractLinks_Ex4(t *testing.T) {
	html, err := os.ReadFile("testdata/ex4.html")
	checkError(err, t)

	got, err := ExtractLinks([]byte(html))
	want := map[string]string{
		"/dog-cat": "dog cat",
	}

	checkError(err, t)
	checkLinkMaps(got, want, t)
}

func checkError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Failed with error: %s\n", err)
	}
}

func checkLinkMaps(got, want map[string]string, t *testing.T) {
	if len(got) != len(want) {
		t.Errorf("Wrong links!\nGot:\n%v\nWant:\n%v\n", got, want)
	}

	for key, got_value := range got {
		want_value, ok := want[key]
		if !ok || want_value != got_value {
			t.Errorf("Wrong links!\nGot:\n%v\nWant:\n%v\n", got, want)
		}
	}
}
