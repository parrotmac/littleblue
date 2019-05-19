import React, {Component} from "react";
import { Card, InputGroup, Tab, Tabs } from "@blueprintjs/core";
import {Elevation} from "@blueprintjs/core/lib/esm/common/elevation";

import BuildLog from "../components/BuildLog";
import {IBuildJob, IBuildStatus} from "../utils/types";

import "flexboxgrid-sass/flexboxgrid.scss";
import "./RepoPage.scss";

interface IProps {
    isLoading: boolean
    buildJobsStatus: IBuildStatus
}

class RepoPage extends Component<IProps> {

    private wrapRepoComponent = (buildJob: IBuildJob): JSX.Element => {
        return (
            <Card elevation={Elevation.TWO}>
                <h3>{buildJob.repo_name}</h3>
                <BuildLog isLoading={this.props.isLoading} job={buildJob} />
            </Card>
        )
    };

    public render(): JSX.Element {
        const { buildJobsStatus: { jobs } } = this.props;
        const defaultTabID = jobs.length>0?`repo-id-${jobs[0].repo_name}`:undefined;
        return (
            <div className={"row center-xs"}>
                <Card elevation={Elevation.TWO} className={"col-xs-12"}>
                    <Tabs
                        animate={true}
                        id="TabsExample"
                        renderActiveTabPanelOnly={true}
                        vertical={true}
                        defaultSelectedTabId={defaultTabID}
                    >
                        <h3>Repositories</h3>
                        <InputGroup type="text" placeholder="Search..." />
                        {jobs.map(job => <Tab
                            id={`repo-id-${job.repo_name}`} /* FIXME */
                            title={job.repo_name}
                            panel={
                                this.wrapRepoComponent(job)
                            }
                            key={`repo-tab-item-${job.repo_name}`}
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
