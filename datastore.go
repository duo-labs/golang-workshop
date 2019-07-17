package main

import "errors"

// URLEntry represents a mapping between a shortcode and its URL.
// To support analytics, we've also introduced a `Hits` attribute that is
// expected to increment every time a user requests a certain shortcode.
type URLEntry struct {
	Shortcode string
	URL       string
	Hits      int
}

// NewURLEntry returns a new URLEntry
func NewURLEntry(shortcode, url string) *URLEntry {
	return &URLEntry{
		Shortcode: shortcode,
		URL:       url,
	}
}

// DataStore implements CRUD operations for a set of shortcode to URL mappings
type DataStore interface {
	AddURL(entry *URLEntry) error
	GetURL(shortcode string) (*URLEntry, error)
	ListURLs() ([]*URLEntry, error)
	HitURL(shortcode string) (*URLEntry, error)
}

// MapStore is a datastore implemented as an in-memory map.
//
// Note: This implementation isn't good for production, since write access
// to the map isn't concurrency-safe. Though, for this workshop it's good
// enough.
type MapStore struct {
	urls map[string]*URLEntry
}

// NewMapStore returns a new MapStore
func NewMapStore() *MapStore {
	store := &MapStore{
		urls: make(map[string]*URLEntry),
	}
	return store
}

// AddURL adds a new URLEntry to the MapStore
func (store *MapStore) AddURL(entry *URLEntry) error {
	short := entry.Shortcode

	// check for duplicate entry
	if _, ok := store.urls[short]; ok {
		return errors.New("Duplicate shortcode entry")
	}

	store.urls[short] = entry
	return nil
}

// GetURL returns the URLEntry for a given shortcode, if it exists
func (store *MapStore) GetURL(shortcode string) (*URLEntry, error) {
	entry, ok := store.urls[shortcode]
	if !ok {
		return entry, errors.New("Shortcode does not exist")
	}
	return entry, nil
}

// ListURLs returns a list of all URLEntry items in the datastore
func (store *MapStore) ListURLs() ([]*URLEntry, error) {
	entries := []*URLEntry{}
	// TODO 1: return a list of all URLEntrys in the MapStore
	//
	// HINT: To iterate over maps (and channels!) with a for loop, Go has a
	// special keyword called "range". You can use it in for loops, as shown
	// here: https://blog.golang.org/go-maps-in-action#TOC_3.
	//
	// HINT: To append a URLEntry to an array (actually a "slice" in Go terms)
	// you can use the "append" function like this:
	// entries = append(entries, some_url_entry)

	// Finally, you'll want to change this call to errors.New to be "nil" when
	// you've implemented the function.
	return entries, errors.New("not implemented")
}

// HitURL increments the number of hits for a given shortcode
func (store *MapStore) HitURL(shortcode string) (*URLEntry, error) {
	// TODO 2: Get the URLEntry from the shortcode provided, increment its Hits
	// property, and then return the URLEntry
	//
	// HINT: You can use the store.GetURL function from above to handle getting
	// the right URLEntry for you.
	return &URLEntry{}, errors.New("not implemented")
}
