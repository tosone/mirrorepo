import React from 'react';
import PropTypes from 'prop-types';
import Layout from '../../components/Layout';
import s from './styles.css';
import { title, html } from './index.md';
import { Button, IconButton } from 'react-toolbox/lib/button';

class HomePage extends React.Component {
  static propTypes = {
    articles: PropTypes.arrayOf(
      PropTypes.shape({
        url: PropTypes.string.isRequired,
        title: PropTypes.string.isRequired,
        author: PropTypes.string.isRequired,
      }).isRequired,
    ).isRequired,
  };

  componentDidMount() {
    document.title = title;
  }

  render() {
    return (
      <Layout className={s.content}>
        <div // eslint-disable-next-line react/no-danger
          dangerouslySetInnerHTML={{
            __html: html,
          }}
        />
        <h4>Articles</h4>
        <Button icon="add" floating />
        <ul>
          {this.props.articles.map(article => (
            <li key={article.url}>
              <a href={article.url}>{article.title}</a>
              by {article.author}
            </li>
          ))}
        </ul>
        <p>
          <br />
          <br />
        </p>
      </Layout>
    );
  }
}

export default HomePage;
