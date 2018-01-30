import React from 'react';
import axios from 'axios';
import { Button, IconButton, Input } from 'react-toolbox';

import s from './styles.css';
import history from '../../history';
import Link from '../../components/Link';

class Login extends React.Component {
  state = { name: '', password: '' };

  componentDidMount() {
    document.title = 'Login';
  }

  login() {
    axios
      .post('/login', {
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
          <Input type="text" label="Name" name="name" value={this.state.name} maxLength={16} onChange={this.changeHandler.bind(this, 'name')} />
          <Input type="text" label="Password" name="password" value={this.state.password} maxLength={16} onChange={this.changeHandler.bind(this, 'password')} />
          <Button className={s.loginBtn} label="登陆" onClick={this.login.bind(this)} raised />
        </main>
      </div>
    );
  }
}

export default Login;
