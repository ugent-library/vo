package vo

import (
	"net/url"
	"slices"
)

var (
	RuleURL = "url"

	MessageURL = "must be a valid URL"
)

// URL checks if the given string is a valid absolute URL. If one or more schemes
// are given, the URL's scheme must be one of them; otherwise any scheme is allowed.
//
//	err := vo.URL("uri", "ftp://host/file")                     // any scheme
//	err := vo.URL("homepage", "https://example.com", "http", "https") // web case
func URL(key, val string, schemes ...string) *Error {
	u, err := url.Parse(val)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return NewError(key, RuleURL).WithMessage(MessageURL)
	}
	if len(schemes) > 0 && !slices.Contains(schemes, u.Scheme) {
		return NewError(key, RuleURL, schemes).WithMessage(MessageURL)
	}
	return nil
}

// WebURL checks if the given string is a valid http or https URL.
//
//	err := vo.WebURL("homepage", "https://example.com")
func WebURL(key, val string) *Error {
	return URL(key, val, "http", "https")
}
