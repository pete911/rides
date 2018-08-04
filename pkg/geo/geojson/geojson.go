package geojson

import (
	"errors"
	"github.com/pete911/rides/pkg/geo/gpx"
)

// for simplicity only line string geometry is supported
// this should be OK for any ride
type GeoJson struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
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
	Type        string     `json:"type"` // only 'LineString' is supported
	Coordinates LineString `json:"coordinates"`
}

type LineString [][3]int

func FromGPX(g gpx.Gpx) (GeoJson, error) {

	gj := GeoJson{Type: "FeatureCollection"}

	// basic validation
	if len(g.Trk) == 0 {
		return gj, errors.New("gpx does not contain any tracks")
	}
	if len(g.Trk) != 1 {
		return gj, errors.New("gpx contains more than one track")
	}

	// properties
	var links []Link
	for _, link := range g.Trk[0].Links {
		links = append(links, Link{Href: link.Href})
	}

	properties := Properties{Name: g.Trk[0].Name, Links: links, Type: g.Trk[0].Type}
	feature := Feature{Type: "Feature", Properties: properties}

	// coordinates
	geometry := Geometry{Type: "LineString"}
	for _, trkseg := range g.Trk[0].Trkseg {
		for _, wpt := range trkseg.Trkpt {
			geometry.Coordinates = append(geometry.Coordinates, [3]int{int(wpt.Lon), int(wpt.Lat), int(wpt.Ele)})
		}
	}

	feature.Geometry = geometry
	gj.Features = []Feature{feature}
	return gj, nil
}
