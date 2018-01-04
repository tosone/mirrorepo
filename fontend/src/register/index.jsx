import React from 'react';
import history from '../history';
import Link from '../../components/Link';
import s from './styles.css';
import axios from 'axios';
import { Button, IconButton, Input } from 'react-toolbox';

class Login extends React.Component {
  state = {
    name: '',
    password: '',
  };
  static propTypes = {
    error: React.PropTypes.object,
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

  render() {
    return (
      <div className={s.container}>
        <main className={s.loginForm}>
          <Input
            type="text"
            label="Name"
            name="name"
            value={this.state.name}
            maxLength={16}
            onChange={this.changeHandler.bind(this, 'name')}
          />
          <Input
            type="text"
            label="Password"
            name="password"
            value={this.state.password}
            maxLength={16}
            onChange={this.changeHandler.bind(this, 'password')}
          />
          <Button
            className={s.loginBtn}
            label="登陆"
            raised
            onClick={this.login.bind(this)}
          />
        </main>
      </div>
    );
  }
}

export default Login;
