package geojson

import (
	"github.com/pete911/rides/pkg/geo/gpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromGPX(t *testing.T) {

	trk := gpx.Trk{
		Type:  "Ride",
		Name:  "Short Bike Ride",
		Links: []gpx.Link{{Href: "https://one-ride.com"}, {Href: "https://two-ride.com"}},
		Trkseg: []gpx.Trkseg{
			{
				Trkpt: []gpx.Wpt{
					{Lat: 4, Lon: 5, Ele: 1},
					{Lat: 5, Lon: 6},
				},
			},
		},
	}

	g := gpx.Gpx{Version: "1.0", Creator: "test", Trk: []gpx.Trk{trk}}
	gj, err := FromGPX(g)
	require.NoError(t, err)

	require.Equal(t, 1, len(gj.Features))
	assert.Equal(t, "LineString", gj.Features[0].Geometry.Type)
	require.Equal(t, 2, len(gj.Features[0].Geometry.Coordinates))

	assert.Equal(t, "Short Bike Ride", gj.Features[0].Properties.Name)
	assert.Equal(t, "Ride", gj.Features[0].Properties.Type)
	assert.Equal(t, 2, len(gj.Features[0].Properties.Links))

	assert.Equal(t, 5, gj.Features[0].Geometry.Coordinates[0][0])
	assert.Equal(t, 4, gj.Features[0].Geometry.Coordinates[0][1])
	assert.Equal(t, 1, gj.Features[0].Geometry.Coordinates[0][2])
	assert.Equal(t, 6, gj.Features[0].Geometry.Coordinates[1][0])
	assert.Equal(t, 5, gj.Features[0].Geometry.Coordinates[1][1])
	assert.Equal(t, 0, gj.Features[0].Geometry.Coordinates[1][2])
}

func TestFromGPXWithMultipleTrksegElements(t *testing.T) {

	trk := gpx.Trk{
		Trkseg: []gpx.Trkseg{
			{
				Trkpt: []gpx.Wpt{
					{Lat: 4, Lon: 5, Ele: 1},
					{Lat: 5, Lon: 6, Ele: 0},
				},
			},
			{
				Trkpt: []gpx.Wpt{
					{Lat: 6, Lon: 7, Ele: 3},
				},
			},
		},
	}

	g := gpx.Gpx{Version: "1.0", Creator: "test", Trk: []gpx.Trk{trk}}
	gj, err := FromGPX(g)
	require.NoError(t, err)

	require.Equal(t, 1, len(gj.Features))
	assert.Equal(t, "LineString", gj.Features[0].Geometry.Type)
	require.Equal(t, 3, len(gj.Features[0].Geometry.Coordinates))

	assert.Equal(t, 5, gj.Features[0].Geometry.Coordinates[0][0])
	assert.Equal(t, 4, gj.Features[0].Geometry.Coordinates[0][1])
	assert.Equal(t, 1, gj.Features[0].Geometry.Coordinates[0][2])
	assert.Equal(t, 7, gj.Features[0].Geometry.Coordinates[2][0])
	assert.Equal(t, 6, gj.Features[0].Geometry.Coordinates[2][1])
	assert.Equal(t, 3, gj.Features[0].Geometry.Coordinates[2][2])
}

func TestFromGPXNoTracks(t *testing.T) {

	g := gpx.Gpx{Version: "1.0", Creator: "test"}
	_, err := FromGPX(g)

	require.Error(t, err)
}

func TestFromGPXMoreThanOneTrack(t *testing.T) {

	g := gpx.Gpx{Version: "1.0", Creator: "test", Trk: []gpx.Trk{{}, {}, {}}}
	_, err := FromGPX(g)

	require.Error(t, err)
}
