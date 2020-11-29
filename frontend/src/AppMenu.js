import React from 'react';
import {Menu, Icon, Input, Dropdown} from 'semantic-ui-react';

export default function AppMenu() {
  const locations = [
    {key: '1', value: '1', text: 'Pittsburgh'},
    {key: '2', value: '2', text: 'San Francisco'},
  ];

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
            options={locations}
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