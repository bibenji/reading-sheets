package meander

// APIKey Google APIKey for Google Places API
var APIKey string

// Place a place
type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

// Public Public fn for Place
func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":      p.Name,
		"icon":      p.Icon,
		"photos":    p.Photos,
		"vicintity": p.Vicinity,
		"lat":       p.Lat,
		"lng":       p.Lng,
	}
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}
