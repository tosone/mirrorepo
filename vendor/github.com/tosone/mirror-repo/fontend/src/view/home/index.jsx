import React from 'react';
import PropTypes from 'prop-types';

import Layout from '../../components/Layout';
import s from './styles.css';
import { title, html } from './index.md';
import { Button, IconButton } from 'react-toolbox/lib/button';

class Home extends React.Component {
  componentDidMount() {
    document.title = title;
  }

  render() {
    return (
      <Layout className={s.content}>
        <h4>Articles</h4>
        <Button icon="add" floating />
      </Layout>
    );
  }
}

export default Home;
