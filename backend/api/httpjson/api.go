package httpjson

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/austinhyde/seating/service"
)

type HttpJsonApi struct {
	Log     *zerolog.Logger
	Service *service.SeatingService
}

func (self *HttpJsonApi) GetHttpHandler() http.Handler {
	m := mux.NewRouter()

	m.Path("/get_locations").Handler(self.wrapHandler(self.HandleGetLocations))
	m.Methods("POST").Path("/get_location_by_id").Handler(self.wrapHandler(self.HandleGetLocationById))
	m.Methods("POST").Path("/get_location_nearest").Handler(self.wrapHandler(self.HandleGetLocationNearest))
	m.Methods("POST").Path("/save_location").Handler(self.wrapHandler(self.HandleSaveLocation))
	m.Methods("POST").Path("/remove_location").Handler(self.wrapHandler(self.HandleRemoveLocation))

	return m
}

func (self *HttpJsonApi) wrapHandler(h HandlerFunc) http.Handler {
	return WrapHandler(self.Log, h)
}
