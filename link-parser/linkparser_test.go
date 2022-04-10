package linkparser

import (
	"os"
	"strings"
	"testing"
)

func TestExtractLinks_NakedA(t *testing.T) {
	html := `<a href="/link">Link text</a>`

	got, err := Parse(strings.NewReader(html))
	want := []Link{
		{"/link", "Link text"},
	}

	checkError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex1(t *testing.T) {
	html, err := os.Open("testdata/ex1.html")
	checkError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"/other-page", "A link to another page"},
	}

	checkError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex2(t *testing.T) {
	html, err := os.Open("testdata/ex2.html")
	checkError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
		{"https://github.com/gophercises", "Gophercises is on Github!"},
	}

	checkError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex3(t *testing.T) {
	html, err := os.Open("testdata/ex3.html")
	checkError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"#", "Login"},
		{"/lost", "Lost? Need help?"},
		{"https://twitter.com/marcusolsson", "@marcusolsson"},
	}

	checkError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex4(t *testing.T) {
	html, err := os.Open("testdata/ex4.html")
	checkError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"/dog-cat", "dog cat"},
	}

	checkError(err, t)
	checkLinks(got, want, t)
}

// Helper functions ////

func checkError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Failed with error: %s\n", err)
	}
}

func checkLinks(got, want []Link, t *testing.T) {
	t.Helper()
	if len(got) != len(want) {
		reportLinksError(got, want, t)
	}

	for i, link := range got {
		if want[i] != link {
			reportLinksError(got, want, t)
		}
	}
}

func reportLinksError(got, want []Link, t *testing.T) {
	t.Helper()
	t.Logf("Got:\n")
	for _, link := range got {
		t.Logf("\t%+v\n", link)
	}
	t.Logf("Want:\n")
	for _, link := range want {
		t.Logf("\t%+v\n", link)
	}
	t.Errorf("Wrong links!\n")
}
