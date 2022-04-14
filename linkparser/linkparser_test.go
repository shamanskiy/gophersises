package linkparser

import (
	"os"
	"strings"
	"testing"

	"github.com/Shamanskiy/gophercises/base"
)

func TestExtractLinks_NakedA(t *testing.T) {
	html := `<a href="/link">Link text</a>`

	got, err := Parse(strings.NewReader(html))
	want := []Link{
		{"/link", "Link text"},
	}

	base.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex1(t *testing.T) {
	html, err := os.Open("testdata/ex1.html")
	base.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"/other-page", "A link to another page"},
	}

	base.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex2(t *testing.T) {
	html, err := os.Open("testdata/ex2.html")
	base.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
		{"https://github.com/gophercises", "Gophercises is on Github!"},
	}

	base.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex3(t *testing.T) {
	html, err := os.Open("testdata/ex3.html")
	base.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"#", "Login"},
		{"/lost", "Lost? Need help?"},
		{"https://twitter.com/marcusolsson", "@marcusolsson"},
	}

	base.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex4(t *testing.T) {
	html, err := os.Open("testdata/ex4.html")
	base.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"/dog-cat", "dog cat"},
	}

	base.CheckError(err, t)
	checkLinks(got, want, t)
}

func checkLinks(got, want []Link, t *testing.T) {
	t.Helper()
	if !base.SameElements(got, want) {
		base.ReportDifferentSlices(got, want, "Different links!", t)
	}
}
