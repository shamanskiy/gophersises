package linkparser

import (
	"os"
	"strings"
	"testing"

	"github.com/Shamanskiy/gophercises/testutils"
)

func TestExtractLinks_NakedA(t *testing.T) {
	html := `<a href="/link">Link text</a>`

	got, err := Parse(strings.NewReader(html))
	want := []Link{
		{"/link", "Link text"},
	}

	testutils.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex1(t *testing.T) {
	html, err := os.Open("testdata/ex1.html")
	testutils.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"/other-page", "A link to another page"},
	}

	testutils.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex2(t *testing.T) {
	html, err := os.Open("testdata/ex2.html")
	testutils.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
		{"https://github.com/gophercises", "Gophercises is on Github!"},
	}

	testutils.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex3(t *testing.T) {
	html, err := os.Open("testdata/ex3.html")
	testutils.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"#", "Login"},
		{"/lost", "Lost? Need help?"},
		{"https://twitter.com/marcusolsson", "@marcusolsson"},
	}

	testutils.CheckError(err, t)
	checkLinks(got, want, t)
}

func TestExtractLinks_Ex4(t *testing.T) {
	html, err := os.Open("testdata/ex4.html")
	testutils.CheckError(err, t)

	got, err := Parse(html)
	want := []Link{
		{"/dog-cat", "dog cat"},
	}

	testutils.CheckError(err, t)
	checkLinks(got, want, t)
}

func checkLinks(got, want []Link, t *testing.T) {
	t.Helper()
	if !testutils.SameElements(got, want) {
		testutils.ReportDifferentSlices(got, want, "Different links!", t)
	}
}
