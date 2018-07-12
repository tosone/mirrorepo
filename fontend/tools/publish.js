const path = require('path');
const firebase = require('firebase-tools');

const build = require('./build');
const task = require('./task');
const config = require('./config');

module.exports = task('deploy', () =>
  Promise.resolve()
    .then(() => build())
    .then(() =>
      firebase.login({
        nonInteractive: false,
      }),
    )
    .then(() =>
      firebase.deploy({
        project: config.project,
        cwd: path.resolve(__dirname, '../'),
      }),
    )
    .then(() => {
      setTimeout(() => process.exit());
    }),
);
