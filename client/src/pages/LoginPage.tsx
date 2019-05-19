import React from 'react';

import LoginForm from '../components/LoginForm';

import './LoginPage.scss';

class LoginPage extends React.Component {

  public render(): JSX.Element {
    return (
      <div className={'LoginPage'}>
        <h1>The Little Blue Container Builder</h1>
        <hr style={{ width: '35%' }} />
        <p>Please sign in</p>
        <LoginForm />
      </div>
    );
  }
}

export default LoginPage;
