import 'babel-polyfill';
import 'whatwg-fetch';

import React from 'react';
import ReactDOM from 'react-dom';
import FastClick from 'fastclick';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import { Router } from 'react-router-dom';

import Routers from './routers.jsx';
import history from './history';
import reducers from './reducers';

const store = createStore(reducers, applyMiddleware(thunk));

FastClick.attach(document.body);

ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <Routers />
    </Router>
  </Provider>,
  document.getElementById('container'),
);
