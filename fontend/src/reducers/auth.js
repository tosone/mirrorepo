export default (state = {}, action) => {
  switch (action.type) {
    case 'signupSuccess':
      return { ...state, singup: true, error: {} };
    case 'signupError':
      return { ...state, singup: false, error: {} };
    default:
      return { ...state };
  }
};
