package dictionary

import "sort"
type Entry struct {
	Definition string
}

func (e Entry) String() string {

	return e.Definition
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {

	return &Dictionary{entries: make(map[string]Entry)}
}

func (d *Dictionary) Add(word string, definition string) {
	d.entries[word] = Entry{Definition: definition}
}

func (d *Dictionary) Get(word string) (Entry, bool) {
	entry, exists := d.entries[word]
	return entry, exists
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}



func (d *Dictionary) List() ([]string) {
	var words []string
	for word := range d.entries {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}
