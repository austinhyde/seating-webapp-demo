package types

import (
	"database/sql/driver"
	"time"

	"github.com/cridenour/go-postgis"
	"github.com/google/uuid"
)

// TOOD: These are coupled to JSON and datbase representations
// when meant to be a generic type. need to clean these up

type Location struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Location    LatLon    `json:"location"`
}

type LatLon struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// Scan provides a custom implementation of database/sql.Scanner
func (self *LatLon) Scan(val interface{}) error {
	p := postgis.PointS{}
	// note: pgx appears to make `val` a string, whereas `pq` (the old driver)
	// used `[]byte`. this postgis library only supports `[]byte`, so we need to do some casting
	// https://github.com/cridenour/go-postgis/blob/master/decode.go#L21
	err := p.Scan([]byte(val.(string)))
	if err != nil {
		return err
	}
	self.Latitude = p.Y
	self.Longitude = p.X
	return nil
}

// Value provides a custom implementation of database/sql/driver.Valuer
func (self *LatLon) Value() (driver.Value, error) {
	p := postgis.PointS{X: self.Longitude, Y: self.Latitude}
	v, err := p.Value()
	// note: pgx appears to need this to be a string, whereas `pq` is fine
	// with the `[]byte` that this postgis library returns
	return string(v.([]byte)), err
}

type UpdatedLocation struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Location    LatLon     `json:"location"`
}
