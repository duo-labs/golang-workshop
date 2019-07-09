package main

import "testing"

const exampleShortcode = "example"
const exampleURL = "http://example.com"

func setupTest(t *testing.T) DataStore {
	store := NewMapStore()
	expected := NewURLEntry(exampleShortcode, exampleURL)
	store.urls[exampleShortcode] = expected
	return store
}

func TestGetURL(t *testing.T) {
	store := setupTest(t)
	got, err := store.GetURL(exampleShortcode)
	if err != nil {
		t.Fatalf("error getting url: %v", err)
	}
	if got.URL != exampleURL {
		t.Fatalf("incorrect url received. expected %s got %s", got.URL, exampleURL)
	}
}

func TestHitURL(t *testing.T) {
	store := setupTest(t)
	expectedHitCount := 1
	got, err := store.HitURL(exampleShortcode)
	if err != nil {
		t.Fatalf("error increasing url hit count: %v", err)
	}
	if got.Hits != expectedHitCount {
		t.Fatalf("incorrect hit count received. expected %d got %d", expectedHitCount, got.Hits)
	}
	got, _ = store.GetURL(exampleShortcode)
	if got.Hits != expectedHitCount {
		t.Fatalf("incorrect hit count received. expected %d got %d", expectedHitCount, got.Hits)
	}
}

func TestListURLs(t *testing.T) {
	store := setupTest(t)
	urls, err := store.ListURLs()
	if err != nil {
		t.Fatalf("error getting list of urls: %v", err)
	}
	if len(urls) != 1 {
		t.Fatalf("invalid number of urls received. expected %d got %d", 1, len(urls))
	}
	got := urls[0]
	if got.Shortcode != exampleShortcode {
		t.Fatalf("invalid UrlEntry received. expected shortcode %s got %s", exampleShortcode, got.Shortcode)
	}
}
