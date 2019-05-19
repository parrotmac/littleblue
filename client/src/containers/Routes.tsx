import React, { Component } from 'react';
import { Route, Switch } from 'react-router';

import LoginPage from '../pages/LoginPage';
import NotFound from '../pages/NotFound';
import SettingsPage from '../pages/SettingsPage';
import RepoDataProvider from "./RepoDataProvider";

class Routes extends Component {
    public render(): JSX.Element {
        return (
            <>
                <Switch>
                    <Route path={'/login'} component={LoginPage as any} />
                    <Route path={'/repos'} component={RepoDataProvider as any} />
                    <Route path={'/settings'} component={SettingsPage as any} />
                    <Route component={NotFound as any} />
                </Switch>
            </>
        );
    }
}

export default Routes;
