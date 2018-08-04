package ride

import "github.com/pete911/rides/pkg/geo"

type Ride struct {
	User  User
	Group string
	GEO   geo.GEO
}
