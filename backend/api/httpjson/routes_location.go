package httpjson

import (
	"encoding/json"
	"net/http"

	"github.com/austinhyde/seating/types"
	"github.com/google/uuid"
)

func (self *HttpJsonApi) HandleGetLocations(r *http.Request) Response {
	locs, err := self.Service.GetLocations(r.Context())
	if err != nil {
		return fail(500, "Unknown Error", err.Error())
	}
	return ok(locs)
}

func (self *HttpJsonApi) HandleGetLocationById(r *http.Request) Response {
	var body struct {
		ID uuid.UUID `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return fail(400, "Invalid json body", err.Error())
	}

	loc, err := self.Service.GetLocationById(r.Context(), body.ID)
	if err != nil {
		return fail(500, "Unknown Error", err.Error())
	}
	if loc == nil {
		return fail(404, "Location not found", "")
	}
	return ok(loc)
}

func (self *HttpJsonApi) HandleGetLocationNearest(r *http.Request) Response {
	var body types.LatLon
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return fail(400, "Invalid json body", err.Error())
	}

	loc, err := self.Service.GetLocationNearest(r.Context(), body)
	if err != nil {
		return fail(500, "Unknown Error", err.Error())
	}
	if loc == nil {
		return fail(404, "Location not found", "")
	}
	return ok(loc)
}

func (self *HttpJsonApi) HandleSaveLocation(r *http.Request) Response {
	var body types.UpdatedLocation
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return fail(400, "Invalid json body", err.Error())
	}

	loc, err := self.Service.SaveLocation(r.Context(), body)
	if err != nil {
		return fail(500, "Unknown Error", err.Error())
	}
	return ok(loc)
}

func (self *HttpJsonApi) HandleRemoveLocation(r *http.Request) Response {
	var body struct {
		ID uuid.UUID `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return fail(400, "Invalid json body", err.Error())
	}

	loc, err := self.Service.RemoveLocation(r.Context(), body.ID)
	if err != nil {
		return fail(500, "Unknown Error", err.Error())
	}
	return ok(loc)
}
