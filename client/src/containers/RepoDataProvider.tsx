import React, {Component} from "react";
import RepoPage from "../pages/RepoPage";
import {IBuildStatus} from "../utils/types";
import ApiClient from "../utils/apiClient";

interface IState {
    buildStatus: IBuildStatus
    isLoading: boolean
}

class RepoDataProvider extends Component {
    state = {
        buildStatus: {
            jobs: []
        },
        isLoading: true
    };

    public componentDidMount(): void {
        const apiClient = new ApiClient("");
        apiClient.fetchJobs().then(
            jobStatus => {
                const newState: IState = {
                    buildStatus: jobStatus,
                    isLoading: false
                };
                this.setState(newState as any);
            }
        ).catch(console.error);
    }

    public render(): JSX.Element {
        const { isLoading, buildStatus } = this.state;
        return (
            <RepoPage isLoading={isLoading} buildJobsStatus={buildStatus} />
        )
    }
}

export default RepoDataProvider;