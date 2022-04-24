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

// Test that duplicated links and external links don't end up in the site map.
func TestMapBuilder_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Server: GET request at %v\n", r.URL)
		links := []string{"/posts", "/about", "/about", "https://other-domain.com"}
		responseWithLinks.Execute(w, links)
	}))
	defer server.Close()

	got, err := BuildMap(server.URL, nil)
	want := []string{
		server.URL,
		server.URL + "/posts",
		server.URL + "/about",
	}

	base.CheckError(err, t)
	compareSiteMaps(got, want, t)
}

// Test that circular links don't lead to a hang.
func TestMapBuilder_MultiplePages(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Server: GET request at %v\n", r.URL)

		var links []string
		switch r.URL.Path {
		case "/":
			links = []string{"/posts", "/about"}
		case "/posts":
			links = []string{"/", "/about"}
		case "/about":
			links = []string{"/", "/posts"}
		}

		responseWithLinks.Execute(w, links)
	}))
	defer server.Close()

	got, err := BuildMap(server.URL, nil)
	want := []string{
		server.URL,
		server.URL + "/posts",
		server.URL + "/about",
	}

	base.CheckError(err, t)
	compareSiteMaps(got, want, t)
}

func TestSameDomainLink(t *testing.T) {
	domain := "https://example.com"

	t.Run("domain url", func(t *testing.T) {
		url := "https://example.com"
		if !sameDomainLink(url, domain) {
			t.Fatalf("%s is on domain %s\n", url, domain)
		}
	})

	t.Run("empty", func(t *testing.T) {
		url := ""
		if sameDomainLink(url, domain) {
			t.Fatalf("%s is not on domain %s\n", url, domain)
		}
	})

	t.Run("domain/home url", func(t *testing.T) {
		url := "https://example.com/home"
		if !sameDomainLink(url, domain) {
			t.Fatalf("%s is on domain %s\n", url, domain)
		}
	})

	t.Run("relative url", func(t *testing.T) {
		url := "/home"
		if !sameDomainLink(url, domain) {
			t.Fatalf("%s is on domain %s\n", url, domain)
		}
	})

	t.Run("other domain", func(t *testing.T) {
		url := "https://google.com/home"
		if sameDomainLink(url, domain) {
			t.Fatalf("%s is not on domain %s\n", url, domain)
		}
	})
}

func TestFormalHRef(t *testing.T) {
	domain := "https://example.com"

	t.Run("remove trailing slash", func(t *testing.T) {
		url := "https://example.com/"

		got := formatHRef(url, domain)
		want := "https://example.com"
		base.CheckEqual(got, want, t)
	})

	t.Run("relative url", func(t *testing.T) {
		url := "/home"

		got := formatHRef(url, domain)
		want := "https://example.com/home"
		base.CheckEqual(got, want, t)
	})
}

func compareSiteMaps(got, want []string, t *testing.T) {
	t.Helper()
	if !base.SameElements(got, want) {
		base.ReportDifferentSlices(got, want, "Different sitemaps!", t)
	}
}
