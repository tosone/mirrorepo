import axios from 'axios';
import { push } from 'react-router-redux';

import history from '../history';
import { APIURL } from '../definations';

const signup = params => {
  return function(dispatch) {
    axios
      .post(`${APIURL}/signup`, params)
      .then(() => {
        dispatch({ type: 'signupSuccess' });

        history.push(`/reduxauth/signup/verify-email?email=${params.email}`);
      })
      .catch(response => history.push('/'));
  };
};

export { signup };
