import React, { Component } from 'react';
import { Route, Switch } from 'react-router';

import LoginPage from '../pages/LoginPage';
import NotFound from '../pages/NotFound';
import SettingsPage from '../pages/SettingsPage';
import RepoDataProvider from "./RepoDataProvider";
import ApiClient from "../utils/apiClient";

class Routes extends Component<{apiClient: ApiClient}> {

    InjectedRepoDataProvider = (): JSX.Element => {
        const { apiClient } = this.props;
        return <RepoDataProvider apiClient={apiClient} />
    };

    public render(): JSX.Element {
        return (
            <>
                <Switch>
                    <Route path={'/login'} component={LoginPage as any} />
                    <Route path={'/repos'} component={this.InjectedRepoDataProvider} />
                    {/*<Route path={'/repos/{repo_name}/builds/'} component={RepoDataProvider as any} />*/}
                    {/*<Route path={'/repos/{repo_name}/builds/{build_id}'} component={RepoDataProvider as any} />*/}
                    <Route path={'/settings'} component={SettingsPage as any} />
                    <Route component={NotFound as any} />
                </Switch>
            </>
        );
    }
}

export default Routes;
