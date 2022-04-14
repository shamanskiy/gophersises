package sitemap

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/Shamanskiy/gophercises/base"
)

var htmlTemplate string = `
<html>
  <body>
    {{ range .}}
      <a href="{{.}}">Some link text</a>
    {{ end }}
  </body>
</html>
`

var responseWithLinks *template.Template = template.Must(template.New("").Parse(htmlTemplate))

func TestMapBuilder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Server: GET request %v\n", r.URL)
		links := []string{"/home", "/about", "/about", "https://other-domain.com"}
		responseWithLinks.Execute(w, links)
	}))
	defer server.Close()

	got, err := NewSiteMapBuilder(server.URL).Parse()
	want := []string{
		server.URL,
		server.URL + "/home",
		server.URL + "/about",
	}

	base.CheckError(err, t)
	compareSiteMaps(got, want, t)
}

func compareSiteMaps(got, want []string, t *testing.T) {
	t.Helper()
	if !base.SameElements(got, want) {
		base.ReportDifferentSlices(got, want, "Different sitemaps!", t)
	}
}
