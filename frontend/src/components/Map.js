import React, {useEffect, useState} from 'react';
import ReactMapGL, {Marker} from 'react-map-gl';
import {Button, Icon, Input} from 'semantic-ui-react';

export default function Map({locations, centeredLocation, onNewLocation, onSelectLocation}) {
  const [viewport, setViewport] = useState({
    latitude: 40.4499,
    longitude: -79.9871,
    zoom: 12,
  });

  // when the centered location changes, zoom into that location
  useEffect(() => {
    if (centeredLocation) {
      setViewport(vp => ({
        ...vp,
        latitude: centeredLocation.location.lat,
        longitude: centeredLocation.location.lon,
        zoom: 16,
      }));
    }
  }, [centeredLocation && centeredLocation.location]);

  // a "new location" is a temporary one with provisional details. it'll need to be saved, which
  // will invoke the `onNewLocation` callback
  const [newLocation, setNewLocation] = useState(null);

  // when the map is clicked, we create a new location at the clicked coordinates
  const startNewLocation = ({lngLat:[lon, lat], target}) => {
    // HACK: for some reason, Marker captureClick is not capturing the click, and no combination of preventDefault/stopPropagation is
    // preventing this handler from firing, so we need to filter out anything that isn't a straight up map click.
    // a debug session shows that clicking on the map has a target of <div class="overlays" ...>, and other clicks do not
    // this could break on a whim from react-map-gl though, so this is pretty brittle
    if (target.className !== "overlays") {
      return;
    }
    const location = {
      name: 'New Location',
      location: {lon, lat},
    };
    setNewLocation(location);
  };
  const onUpdateNewLocationName = name => setNewLocation(loc => ({...loc, name}));
  const onNewLocationSave = () => onNewLocation(newLocation);
  const onNewLocationCancel = () => setNewLocation(null);

  return (
    <ReactMapGL
      {...viewport}
      width="100%" height="100%"
      onViewportChange={setViewport}
      mapboxApiAccessToken={MAPBOX_TOKEN}
      onClick={startNewLocation}
    >
      {locations.map(loc => <LocationMarker key={loc.id} location={loc} onSelect={onSelectLocation}/>)}
      {newLocation && <LocationMarker location={newLocation} onChangeName={onUpdateNewLocationName} onSave={onNewLocationSave} onCancel={onNewLocationCancel}/>}
    </ReactMapGL>
  );
}

function LocationMarker({location, onSelect, onChangeName, onSave, onCancel}) {
  // a location is "provisional" if there's no id for it yet - those only come from the database
  const provisional = !location.id;

  // provisional locations get a little form for filling out a name - rest of the management will come later
  let nameDisplay = location.name;
  if (provisional) {
    const canSave = location.name.trim().length > 0;
    nameDisplay = (
      <Input size="mini" action value={location.name} onChange={(e, {value}) => onChangeName(value)}>
        <input />
        <Button icon="cancel" size="mini" onClick={() => onCancel()}/>
        <Button primary icon="check" size="mini" disabled={!canSave} onClick={() => onSave(location)}/>
      </Input>
    );
  }

  const onClick = e => {
    if (!provisional) {
      onSelect(location);
    }
  };

  return (
    <Marker
      latitude={location.location.lat}
      longitude={location.location.lon}
      offsetLeft={-12}
      offsetTop={-10.5}
      captureClick // TODO: this doesn't seem to work?
      onClick={onClick}
    >
      <Icon name={provisional ? 'building outline' : 'building'} size="large"/> {nameDisplay}
    </Marker>
  );
}