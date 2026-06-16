package vo

import "testing"

func TestURL(t *testing.T) {
	// any scheme allowed
	valid := []string{
		"https://example.com",
		"http://example.com/path?q=1#frag",
		"ftp://host/file.txt",
		"git://host/repo.git",
	}
	for _, v := range valid {
		if err := URL("uri", v); err != nil {
			t.Errorf("expected %q to be valid, got: %s", v, err)
		}
	}

	invalid := []string{
		"",
		"example.com",
		"/just/a/path",
		"not a url",
		"http://", // no host
	}
	for _, v := range invalid {
		if err := URL("uri", v); err == nil {
			t.Errorf("expected %q to be invalid", v)
		}
	}
}

func TestWebURL(t *testing.T) {
	valid := []string{
		"http://example.com",
		"https://example.com/path",
	}
	for _, v := range valid {
		if err := WebURL("homepage", v); err != nil {
			t.Errorf("expected %q to be valid, got: %s", v, err)
		}
	}

	invalid := []string{
		"ftp://host/file.txt",
		"git://host/repo.git",
		"mailto:admin@example.com",
	}
	for _, v := range invalid {
		if err := WebURL("homepage", v); err == nil {
			t.Errorf("expected %q to be invalid for http/https", v)
		}
	}
}
