package goutil

import "testing"

func TestURL(t *testing.T) {

	urls := []string{}
	urls = append(urls, URLJoin("http://google.com/", "a", "b/", "c/d/e/", "f/g").Format())
	urls = append(urls, URLJoin("http://google.com", "a", "b/", "/c/d/e/", "f/g").Format())
	urls = append(urls, URLJoin("http://google.com", "a", "", "b/", "", "c/d/e/", "/f/g/").Format())

	expectedURL := "http://google.com/a/b/c/d/e/f/g"
	for i, url := range urls {
		if url != expectedURL {
			t.Errorf("Expected %s - Got %s - For #%d", expectedURL, url, i)
		}
	}

	urlWithParameter := URLJoin("http://google.com/", "a", "b/", "c/d/e/", "f/g").
		SetURLParameter("hello", "world").
		SetURLParameter("k", nil).
		SetURLParameter("number", 10.2)

	expectedURLs := []string{
		"http://google.com/a/b/c/d/e/f/g?hello=world&number=10.2&k",
		"http://google.com/a/b/c/d/e/f/g?hello=world&k&number=10.2",
		"http://google.com/a/b/c/d/e/f/g?k&number=10.2&hello=world",
		"http://google.com/a/b/c/d/e/f/g?k&hello=world&number=10.2",
		"http://google.com/a/b/c/d/e/f/g?number=10.2&k&hello=world",
		"http://google.com/a/b/c/d/e/f/g?number=10.2&hello=world&k",
	}

	if !IsIn(urlWithParameter.Format(), expectedURLs...) {
		t.Errorf("%s - Expected sth like %s - %s",
			"URL Parameter parsing",
			expectedURLs[0],
			urlWithParameter.Format(),
		)
	}

}
