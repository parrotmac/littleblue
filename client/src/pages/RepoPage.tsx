import React, {Component} from "react";
import { Card, InputGroup, Tab, Tabs } from "@blueprintjs/core";
import {Elevation} from "@blueprintjs/core/lib/esm/common/elevation";

import BuildLog from "../components/BuildLog";
import {IBuildJob, IRepo} from "../utils/types";

import "flexboxgrid-sass/flexboxgrid.scss";
import "./RepoPage.scss";
import ApiClient from "../utils/apiClient";

interface RepoBuilds {
    repo: IRepo
    builds: Array<IBuildJob>
}

interface IState {
    repoBuilds: Array<RepoBuilds>
}

interface IProps {
    isLoading: boolean
    repos: Array<IRepo>
}

class RepoPage extends Component<IProps> {
    state = {
        repoBuilds: [],
    };

    componentDidMount(): void {
        this.updateBuilds();
    }

    componentDidUpdate(prevProps: Readonly<IProps>, prevState: Readonly<{}>, snapshot?: any): void {
        if (prevProps.repos.length !== this.props.repos.length) {
            this.updateBuilds();
        }
    }

    updateBuilds = () => {
        const { repos } = this.props;
        repos.forEach((repo) => {
            new ApiClient("").fetchJobsForRepo(repo.id).then((jobs => {
                this.setState({
                   repoBuilds: [
                       ...this.state.repoBuilds,
                       {
                           repo: repo,
                           builds: jobs,
                       }
                   ],
                });
            })).catch(console.error);
        });
    };

    private wrapRepoComponent = (repo: IRepo, builds: Array<IBuildJob>): JSX.Element => {
        return (
            <Card elevation={Elevation.TWO}>
                <h3>{repo.name}</h3>
                {builds.map(job =>
                    <>
                    <BuildLog isLoading={false} job={job} />
                    <hr />
                    </>
                )}
            </Card>
        )
    };

    public render(): JSX.Element {
        const { isLoading } = this.props;
        const { repoBuilds } = this.state;
        if (isLoading) {
            return <div>Loading...</div>
        }
        return (
            <div className={"row center-xs"}>
                <Card elevation={Elevation.TWO} className={"col-xs-12"}>
                    <Tabs
                        animate={true}
                        id="TabsExample"
                        renderActiveTabPanelOnly={true}
                        vertical={true}
                        defaultSelectedTabId={'repo-id-0'}
                    >
                        <h3>Repositories</h3>
                        <InputGroup type="text" placeholder="Search..." />
                        {repoBuilds.map((repoBuild: RepoBuilds, index) => <Tab
                            id={`repo-id-${index}`}
                            title={repoBuild.repo.name}
                            panel={
                                this.wrapRepoComponent(repoBuild.repo, repoBuild.builds)
                            }
                            key={`repo-tab-item-${repoBuild.repo.name}`}
                            className={"repo-panel-tab-item"}
                            panelClassName={"repo-panel-detail-card"}
                            />)}
                        <Tabs.Expander />
                    </Tabs>
                </Card>
            </div>
        )
    }
}

export default RepoPage;
