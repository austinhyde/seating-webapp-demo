import React, {useState} from 'react';
import ReactMapGL from 'react-map-gl';

export default function Map() {
  const [viewport, setViewport] = useState({
    latitude: 37.7577,
    longitude: -122.4376,
    zoom: 8
  });
  return (
    <ReactMapGL
      {...viewport}
      width="100%" height="100%"
      onViewportChange={setViewport}
      mapboxApiAccessToken={MAPBOX_TOKEN}
    />
  );
}