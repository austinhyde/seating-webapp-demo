import React from 'react';
import { hot } from 'react-hot-loader/root';

import 'fomantic-ui-css/semantic.min.css';
import './styles.scss';

import AppMenu from './AppMenu';
import Map from './Map';

const Root = () => {
  return (
    <div className="app-shell">
      <AppMenu/>
      <Map/>
    </div>
  )
};

export default hot(Root);