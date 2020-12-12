package thesaurus

// Thesaurus a thesaurus interface,to add interchangeable implementations for other services
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
