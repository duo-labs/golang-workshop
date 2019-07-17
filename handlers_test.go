package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestRestrictedShortcodes(t *testing.T) {
	blacklist := []string{"add", "list"}
	for _, shortcode := range blacklist {
		form := url.Values{
			"shortcode": {shortcode},
			"url":       {"http://example.com"},
		}
		r := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// Make the request, recording the response in our recorder
		w := httptest.NewRecorder()
		AddURL(w, r)

		// Verify that we got an HTTP 400 response, since these shouldn't be allowed!
		if w.Code != http.StatusBadRequest {
			t.Fatalf("unexpected status code received when requesting the blacklist entry \"%s\". expected %d got %d", shortcode, http.StatusBadRequest, w.Code)
		}
	}
}

func TestNonExistentRedirect(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/bogus", nil)
	w := httptest.NewRecorder()

	// Make the request, recording the response in our recorder
	RedirectHandler(w, r)

	if w.Code != http.StatusNotFound {
		t.Fatalf("unexpected status code received when requesting a shortcode that shouldn't exist. expected %d got %d", http.StatusNotFound, w.Code)
	}
}

func TestValidRedirect(t *testing.T) {
	expectedShortcode := "example"
	expectedURL := "http://example.com"

	// Manually add a shortcode -> URL mapping
	urls[expectedShortcode] = expectedURL

	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", expectedShortcode), nil)
	w := httptest.NewRecorder()

	// Make the request, recording the response in our recorder
	RedirectHandler(w, r)

	if w.Code != http.StatusFound {
		t.Fatalf("unexpected redirect status code received when requesting a shortcode that should exist. expected %d got %d", http.StatusFound, w.Code)
	}

	got := w.Header().Get("Location")
	if got != expectedURL {
		t.Fatalf("unexpected redirect location received. expected %s got %s", expectedURL, got)
	}
}
