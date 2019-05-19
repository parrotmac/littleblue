import React, {Component} from "react";
import { default as AnsiUp } from 'ansi_up';

import "./BuildLog.scss";
import {Classes} from "@blueprintjs/core";
import {IBuildJob, IBuildMessageBody_Error, IBuildMessageBody_Stream} from "../utils/types";

export interface IBuildLog {
    isLoading: boolean;
    job: IBuildJob | null; // TODO
    
}

class BuildLog extends Component<IBuildLog> {
    state: IBuildLog = {
        isLoading: true,
        job: null,
    };

    private createHtmlRecord = (messages: Array<any>): string => {
        let messageStream = "";
        messages.forEach((msg) => {
            try {
                const parsedMsg = JSON.parse(msg.body);
                if ("error" in parsedMsg) {
                    const err: IBuildMessageBody_Error = parsedMsg;
                    messageStream += `\u001b[91m${err.error}\n\u001b[0m`; // TODO: Refactor color handling
                } else if ("stream" in parsedMsg) {
                    // Build messages
                    const msg: IBuildMessageBody_Stream = parsedMsg;
                    messageStream += msg.stream;
                } else if ("status" in parsedMsg) {
                    // Push messages
                    const statusText = parsedMsg["status"];

                    if ("id" in parsedMsg) {
                        const imgID = parsedMsg["id"];
                        if ("progress" in parsedMsg) {
                            // Progress bar (in progress)
                            console.log("parsedMsg", parsedMsg); // DEBUG
                            const progressBar = parsedMsg["progress"];
                            messageStream += `${imgID}: ${statusText}\t${progressBar}\n`;
                        } else {
                            // Done pushing
                            messageStream += `${imgID}: ${statusText}\n`;
                        }
                    } else {
                        // Other status
                        messageStream += `${statusText}\n`;
                    }
                } else {
                    console.warn("unable to handle non-stream message", msg);
                }
            } catch (e) {
                // not json
                console.warn("unable to parse JSON for message:", msg);
            }
        });
        
        // @ts-ignore
        const ansi_up = new AnsiUp();
        ansi_up.use_classes = true;
        return ansi_up.ansi_to_html(messageStream);
    };


    private loadingContent = () => {
        return (
            <pre className={`${Classes.SKELETON} build-log`}>loading...</pre>
        )
    };

    public render(): JSX.Element {
        const { isLoading, job } = this.props;
        if (isLoading || job === null) { // FIXME
            return this.loadingContent();
        }
        const buildLog: string = this.createHtmlRecord(job.messages);
        return <pre className={"build-log"} dangerouslySetInnerHTML={{ __html: buildLog }} />;
    }
}

export default BuildLog;
