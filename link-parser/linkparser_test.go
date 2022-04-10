package linkparser

import "testing"

func TestExtractLinks_SingleLink(t *testing.T) {
	html := `<html>
	<body>
	  <h1>Hello!</h1>
	  <a href="/other-page">A link to another page</a>
	</body>
	</html>
	links := ExtractLinks(html)`

	got, err := ExtractLinks([]byte(html))
	want := map[string]string{
		"/other-page": "A link to another page",
	}

	if err != nil {
		t.Errorf("Failed with error: %s\n", err)
	}
	if !linkMapsEqual(got, want) {
		t.Errorf("Wrong links! Got %v, want %v\n", got, want)
	}
}

func TestExtractLinks_TwoLinks(t *testing.T) {
	html := `<html>
	<body>
	  <h1>Hello!</h1>
	  <a href="/other-page">A link to another page</a>
	  <a href="/another-link">Click me!</a>
	</body>
	</html>
	links := ExtractLinks(html)`

	got, err := ExtractLinks([]byte(html))
	want := map[string]string{
		"/other-page":   "A link to another page",
		"/another-link": "Click me!",
	}

	if err != nil {
		t.Errorf("Failed with error: %s\n", err)
	}
	if !linkMapsEqual(got, want) {
		t.Errorf("Wrong links! Got %v, want %v\n", got, want)
	}
}

func linkMapsEqual(got, want map[string]string) bool {
	if len(got) != len(want) {
		return false
	}

	for key, got_value := range got {
		want_value, ok := want[key]
		if !ok || want_value != got_value {
			return false
		}
	}

	return true
}
