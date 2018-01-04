import React from 'react';
import axios from 'axios';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import { Button, IconButton, Input } from 'react-toolbox';

import * as actions from '../../actions/signup';
import history from '../../history';
import Link from '../../components/Link';
import s from './styles.css';

class Signup extends React.Component {
  constructor(props) {
    super(props);
  }
  state = {
    name: '',
    password: '',
  };
  static propTypes = {
    error: PropTypes.object,
  };
  componentDidMount() {
    document.title = 'Login';
  }
  goBack = event => {
    event.preventDefault();
    history.goBack();
  };

  login() {
    axios
      .post('http://127.0.0.1:8080/register', {
        name: this.state.name,
        password: this.state.password,
      })
      .then(res => {
        if (res.status == 200) console.log(res.data);
      });
  }

  changeHandler(name, value) {
    this.setState({ ...this.state, [name]: value });
  }

  handleSubmit() {
    this.props.register(this.state);
  }

  render() {
    return (
      <div className={s.container}>
        <main className={s.loginForm}>
          <Input type="text" label="Name" name="name" value={this.state.name} maxLength={16} onChange={this.changeHandler.bind(this, 'name')} />
          <Input type="text" label="Password" name="password" value={this.state.password} maxLength={16} onChange={this.changeHandler.bind(this, 'password')} />
          <Button className={s.loginBtn} label="注册" raised onClick={this.handleSubmit.bind(this)} />
        </main>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { ...state };
}

export default connect(mapStateToProps, actions)(Signup);
