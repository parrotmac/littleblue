import React, {Component} from 'react';
import { BrowserRouter } from 'react-router-dom';

import NavBar from './components/NavBar';
import Routes from "./containers/Routes";

import './App.scss';
import ApiClient from "./utils/apiClient";

class App extends Component {

    public render(): JSX.Element {
    return (
        <BrowserRouter>
          <>
            <NavBar/>
            <div className="App">
              <Routes apiClient={new ApiClient("")}/>
            </div>
          </>
        </BrowserRouter>
    );
  }
}

export default App;
