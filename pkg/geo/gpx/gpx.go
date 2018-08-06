package gpx

// GPX is the root element in the XML file
type Gpx struct {
	Version    string   `xml:"version,attr"`         // You must include the version number in your GPX document
	Creator    string   `xml:"creator,attr"`         // You must include the name or URL of the software that created your GPX document. This allows others to inform the creator of a GPX instance document that fails to validate
	Metadata   Metadata `xml:"metadata,omitempty"`   // Metadata about the file
	Wpt        []Wpt    `xml:"wpt,omitempty"`        // A list of waypoints
	Rte        []Rte    `xml:"rte,omitempty"`        // A list of routes
	Trk        []Trk    `xml:"trk,omitempty"`        // A list of tracks
	Extensions []byte   `xml:"extensions,omitempty"` // You can add extend GPX by adding your own elements from another schema here
}

// Information about the GPX file, author, and copyright restrictions goes in the metadata section. Providing rich,
// meaningful information about your GPX files allows others to search for and use your GPS data.
type Metadata struct {
	Name       string    `xml:"name,omitempty"`       // The name of the GPX file
	Desc       string    `xml:"desc,omitempty"`       // A description of the contents of the GPX file
	Author     Author    `xml:"author,omitempty"`     // The person or organization who created the GPX file
	Copyright  Copyright `xml:"copyright,omitempty"`  // Copyright and license information governing use of the file
	Links      []Link    `xml:"link,omitempty"`       // URLs associated with the location described in the file
	Time       string    `xml:"time,omitempty"`       // The creation date of the file
	Keywords   string    `xml:"keywords"`             // Keywords associated with the file. Search engines or databases can use this information to classify the data
	Bounds     string    `xml:"bounds"`               // Minimum and maximum coordinates which describe the extent of the coordinates in the file
	Extensions []byte    `xml:"extensions,omitempty"` // You can add extend GPX by adding your own elements from another schema here.
}

type Link struct {
	Href string `xml:"href,attr"` // URL of hyperlink
	Text string `xml:"text"`      // Text of hyperlink
	Type string `xml:"type"`      // Mime type of content (image/jpeg)
}

type Copyright struct {
	Author  string `xml:"author,attr"`       // Copyright holder
	Year    int    `xml:"year,omitempty"`    // Year of copyright
	License string `xml:"license,omitempty"` // Link to external file containing license text
}

// A person or organization
type Author struct {
	Name  string `xml:"name,omitempty"`  // Name of person or organization
	Email string `xml:"email,omitempty"` // Email address
	Link  Link   `xml:"link,omitempty"`  // Link to Web site or other external information about person
}

// rte represents route - an ordered list of waypoints representing a series of turn points leading to a destination
type Rte struct {
	Name       string `xml:"name,omitempty"`       // GPS name of route
	Cmt        string `xml:"cmt,omitempty"`        // GPS comment for route
	Desc       string `xml:"desc,omitempty"`       // Text description of route for user. Not sent to GPS
	Src        string `xml:"src,omitempty"`        // Source of data. Included to give user some idea of reliability and accuracy of data
	Links      []Link `xml:"link,omitempty"`       // Links to external information about the route
	Number     int    `xml:"number,omitempty"`     // GPS route number
	Type       string `xml:"type,omitempty"`       // Type (classification) of route
	Extensions []byte `xml:"extensions,omitempty"` // You can add extend GPX by adding your own elements from another schema here
	Rtept      []Wpt  `xml:"rtept,omitempty"`      // A list of route points
}

// trk represents a track - an ordered list of points describing a path
type Trk struct {
	Name   string   `xml:"name,omitempty"`   // GPS name of track
	Cmt    string   `xml:"cmt,omitempty"`    // GPS comment for track
	Desc   string   `xml:"desc,omitempty"`   // User description of track
	Src    string   `xml:"src,omitempty"`    // Source of data. Included to give user some idea of reliability and accuracy of data
	Links  []Link   `xml:"link,omitempty"`   // Links to external information about track
	Type   string   `xml:"type,omitempty"`   // Type (classification) of track
	Trkseg []Trkseg `xml:"trkseg,omitempty"` // A Track Segment holds a list of Track Points which are logically connected in order. To represent a single GPS track where GPS reception was lost, or the GPS receiver was turned off, start a new Track Segment for each continuous span of track data
}

// A Track Segment holds a list of Track Points which are logically connected in order. To represent a single GPS track
// where GPS reception was lost, or the GPS receiver was turned off, start a new Track Segment for each continuous span
// of track data
type Trkseg struct {
	Trkpt      []Wpt  `xml:"trkpt,omitempty"`      // A Track Point holds the coordinates, elevation, timestamp, and metadata for a single point in a track
	Extensions []byte `xml:"extensions,omitempty"` // You can add extend GPX by adding your own elements from another schema here
}

// trkpt represents a waypoint, point of interest, or named feature on a map
type Wpt struct {
	Lat           float64 `xml:"lat,attr"`                // The latitude of the point. Decimal degrees, WGS84 datum
	Lon           float64 `xml:"lon,attr"`                // The longitude of the point. Decimal degrees, WGS84 datum
	Ele           float64 `xml:"ele,omitempty"`           // Elevation (in meters) of the point
	Time          string  `xml:"time,omitempty"`          // Creation/modification timestamp for element. Date and time in are in Univeral Coordinated Time (UTC), not local time! Conforms to ISO 8601 specification for date/time representation. Fractional seconds are allowed for millisecond timing in tracklogs
	Magvar        float64 `xml:"magvar,omitempty"`        // Magnetic variation (in degrees) at the point
	Geoidheight   float64 `xml:"geoidheight"`             // Height (in meters) of geoid (mean sea level) above WGS84 earth ellipsoid. As defined in NMEA GGA message
	Name          string  `xml:"name,omitempty"`          // The GPS name of the waypoint
	Cmt           string  `xml:"cmt,omitempty"`           // GPS waypoint comment
	Desc          string  `xml:"desc,omitempty"`          // A text description of the element. Holds additional information about the element intended for the user, not the GPS
	Src           string  `xml:"src,omitempty"`           // Source of data. Included to give user some idea of reliability and accuracy of data
	Links         []Link  `xml:"link,omitempty"`          // Link to additional information about the waypoint
	Sym           string  `xml:"sym,omitempty"`           // Text of GPS symbol name. For interchange with other programs, use the exact spelling of the symbol as displayed on the GPS. If the GPS abbreviates words, spell them out
	Type          string  `xml:"type,omitempty"`          // Type (classification) of the waypoint
	Fix           string  `xml:"fix,omitempty"`           // Type of GPX fix, {'none'|'2d'|'3d'|'dgps'|'pps'}. none means GPS had no fix. To signify "the fix info is unknown, leave out fixType entirely. pps = military signal used
	Sat           int     `xml:"sat,omitempty"`           // Number of satellites used to calculate the GPX fix
	Hdop          float64 `xml:"hdop,omitempty"`          // Horizontal dilution of precision
	Vdop          float64 `xml:"vdop,omitempty"`          // Vertical dilution of precision
	Pdop          float64 `xml:"pdop,omitempty"`          // Position dilution of precision
	Ageofdgpsdata float64 `xml:"ageofdgpsdata,omitempty"` // Number of seconds since last DGPS update
	Dgpsid        float64 `xml:"dgpsid,omitempty"`        // ID of DGPS station used in differential correction
	Extensions    []byte  `xml:"extensions,omitempty"`    // You can add extend GPX by adding your own elements from another schema here
}
