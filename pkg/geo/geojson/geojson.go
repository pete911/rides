package geojson

// TODO - gpx to GEOJson (also KML?)

type GeoJson struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
	Geometry Geometry  `json:"geometry"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Links []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
}

type Geometry struct {
	Type string `json:"type"`
	// longitude, latitude, elevation
	Coordinates [3]int `json:"coordinates"`
}
