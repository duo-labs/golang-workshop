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
		t.Fatalf("incorrect url received from datastore when requesting a shortcode that should exist. expected %s got %s", got.URL, exampleURL)
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
		t.Fatalf("incorrect hit count received after incrementing it. expected %d got %d", expectedHitCount, got.Hits)
	}
	got, _ = store.GetURL(exampleShortcode)
	if got.Hits != expectedHitCount {
		t.Fatalf("incorrect hit count received when re-requesting it after incrementing it earlier. expected %d got %d", expectedHitCount, got.Hits)
	}
}

func TestListURLs(t *testing.T) {
	store := setupTest(t)
	urls, err := store.ListURLs()
	if err != nil {
		t.Fatalf("error getting list of urls from the datastore: %v", err)
	}
	if len(urls) != 1 {
		t.Fatalf("invalid number of urls received from the datastore. expected %d got %d", 1, len(urls))
	}
	got := urls[0]
	if got.Shortcode != exampleShortcode {
		t.Fatalf("invalid UrlEntry received. expected shortcode %s got %s", exampleShortcode, got.Shortcode)
	}
}
