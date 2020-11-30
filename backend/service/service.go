package service

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/austinhyde/seating/db"
	"github.com/austinhyde/seating/types"
)

// SeatingService implements all the things you can do with this service
type SeatingService struct {
	DB db.Queryable
}

// GetLocations returns all known Locations
func (self *SeatingService) GetLocations(ctx context.Context) ([]*types.Location, error) {
	rows, err := self.DB.Query(ctx, `
		SELECT id, name, location, description, created_at, modified_at
		FROM location
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetLocations query")
	}

	out := []*types.Location{}
	for rows.Next() {
		loc := &types.Location{}
		err = rows.Scan(&loc.ID, &loc.Name, &loc.Location, &loc.Description, &loc.CreatedAt, &loc.ModifiedAt)
		if err != nil {
			return nil, errors.Wrap(err, "scanning GetLocations result")
		}
		out = append(out, loc)
	}

	return out, nil
}

// GetLocationNearest returns the Location nearest the given point,
// or nil if there are no locations
func (self *SeatingService) GetLocationNearest(ctx context.Context, ll types.LatLon) (*types.Location, error) {
	loc := &types.Location{}
	// https://postgis.net/docs/geometry_distance_knn.html
	err := self.DB.QueryRow(ctx, `
		SELECT id, name, location, description, created_at, modified_at
		FROM location
		ORDER BY location <-> $1 ASC
		LIMIT 1
	`, ll).Scan(&loc.ID, &loc.Name, &loc.Location, &loc.Description, &loc.CreatedAt, &loc.ModifiedAt)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, nil
		}
		return nil, errors.Wrap(err, "GetLocationNearest query")
	}
	return loc, nil
}

// GetLocationById returns the Location with the given id,
// or nil if there is none with that id
func (self *SeatingService) GetLocationById(ctx context.Context, id uuid.UUID) (*types.Location, error) {
	loc := &types.Location{}
	err := self.DB.QueryRow(ctx, `
		SELECT id, name, location, description, created_at, modified_at
		FROM location
		WHERE id = $1
		ORDER BY name ASC
	`, id.String()).Scan(&loc.ID, &loc.Name, &loc.Location, &loc.Description, &loc.CreatedAt, &loc.ModifiedAt)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, nil
		}
		return nil, errors.Wrap(err, "GetLocations query")
	}
	return loc, nil
}

// SaveLocation creates or updates a location with the provided information,
// returning the updated Location
func (self *SeatingService) SaveLocation(ctx context.Context, input types.UpdatedLocation) (*types.Location, error) {
	var sql string
	params := []interface{}{input.Name, input.Description, input.Location}
	if input.ID == nil {
		sql = `
			INSERT INTO location (name, description, location)
			VALUES ($1, $2, $3)
		`
	} else {
		sql = `
			UPDATE location
			SET name = $1,
					description = $2,
					location = $3,
					modified_at = NOW()
			WHERE id = $4
		`
		params = append(params, input.ID)
	}

	sql += `
			RETURNING id, name, description, location, created_at, modified_at
	`

	out := &types.Location{}
	err := self.DB.QueryRow(ctx, sql, params...).
		Scan(&out.ID, &out.Name, &out.Description, &out.Location, &out.CreatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "SaveLocation query")
	}
	return out, nil
}

// RemoveLocation removes the Location with the given id, returning the now-absent Location.
// It is an error to remove a Location that does not exist
func (self *SeatingService) RemoveLocation(ctx context.Context, id uuid.UUID) (*types.Location, error) {
	out := &types.Location{}
	err := self.DB.QueryRow(ctx, `
		DELETE FROM location
		WHERE id = $1
		RETURNING id, name, description, location, created_at, modified_at
	`, id).Scan(&out.ID, &out.Name, &out.Description, &out.Location, &out.CreatedAt, &out.ModifiedAt)
	if err != nil {
		return nil, errors.Wrap(err, "RemoveLocation query")
	}
	return out, nil
}
