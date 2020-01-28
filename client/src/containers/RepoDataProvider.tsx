import React, {Component} from "react";
import RepoPage from "../pages/RepoPage";
import ApiClient from "../utils/apiClient";
import {IRepo} from "../utils/types";

interface IProps {
    apiClient: ApiClient
}

interface IState {
    isLoading: boolean
    repoList: Array<IRepo>
}

class RepoDataProvider extends Component<IProps, IState> {
    state = {
        repoList: [],
        isLoading: true
    };

    componentDidMount = (): void => {
        const {apiClient} = this.props;
        apiClient.fetchRepos().then(
            repoListing => {
                const newState: IState = {
                    repoList:repoListing,
                    isLoading: false
                };
                this.setState(newState);
            }
        ).catch(console.error);
    };

    public render(): JSX.Element {
        const { isLoading, repoList } = this.state;
        return (
            <RepoPage isLoading={isLoading} repos={repoList} />
        )
    }
}

export default RepoDataProvider;