import React from 'react';
import {Menu, Icon, Input, Dropdown} from 'semantic-ui-react';

export default function AppMenu({locations, selectedLocation, onSelectLocation}) {
  const locationOptions = locations.map(loc => ({
    key: loc.id,
    value: loc.id,
    text: loc.name,
  }));

  const locationDropdownChanged = (ev, {value}) => onSelectLocation(locations.find(l => l.id === value));

  const users = [
    {key: '1', value: '1', text: 'Jack O\'Neill'},
    {key: '2', value: '2', text: 'Samantha Carter'},
    {key: '3', value: '3', text: 'Daniel Jackson'},
    {key: '4', value: '4', text: 'Teal\'c'},
  ];

  return (
    <div className="app-menu">
      <Menu attached="top" size="small">
        <Menu.Item header><Icon name="building"/> Seating App</Menu.Item>
        <Menu.Item className="inline-menu-search">
          <Icon name="map marker alternate"/>
          <Dropdown
            placeholder="Select Location"
            search
            selection
            options={locationOptions}
            onChange={locationDropdownChanged}
            value={selectedLocation && selectedLocation.id}
          />
        </Menu.Item>
        <Menu.Item className="inline-menu-search">
          <Icon name="user"/>
          <Dropdown
            placeholder="Search for user..."
            search
            selection
            options={users}
          />
        </Menu.Item>
      </Menu>
    </div>
  );
}