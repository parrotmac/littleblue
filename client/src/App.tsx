import React, {Component} from 'react';
import { BrowserRouter } from 'react-router-dom';

import NavBar from './components/NavBar';
import Routes from "./containers/Routes";

import './App.css';

class App extends Component {

  public render(): JSX.Element {
    return (
        <BrowserRouter>
          <>
            <NavBar/>
            <div className="App">
              <Routes/>
            </div>
          </>
        </BrowserRouter>
    );
  }
}

export default App;
