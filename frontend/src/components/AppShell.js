import React, {useState} from 'react';
import {useAsync} from 'react-async-hook';

import client from '../client';
import AppMenu from './AppMenu';
import Map from './Map';

export default function AppShell() {
  const asyncLocations = useAsync(() => client.getLocations(), []);
  const [centeredLocation, setCenteredLocation] = useState(null);

  if (asyncLocations.error) {
    return (<strong>Could not load locations: {asyncLocations.error}</strong>);
  }

  const locations = asyncLocations.loading ? [] : asyncLocations.result;

  return (
    <div className="app-shell">
      <AppMenu
        locations={locations}
        centeredLocation={centeredLocation}
        onSelectLocation={setCenteredLocation}
      />
      <Map
        locations={locations}
        centeredLocation={centeredLocation}
        onNewLocation={loc => console.log('new location', loc)}
        onSelectLocation={loc => console.log('select location', loc)}
      />
    </div>
  );
};