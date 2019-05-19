export interface IBuildMessageBody_Stream {
    stream: string;
}

export interface IBuildMessageBody_ErrorDetail {
    code: number;
    message: string;
}

export interface IBuildMessageBody_Error {
    error: string;
    errorDetail: IBuildMessageBody_ErrorDetail;
}


export interface IBuildMessage {
    level: string;
    message_source: string; // Not yet available
    repo_full_name: string;
    build_identifier: string;
    body: string; // Possibly JSON
}

export interface IBuildJob {
    repo_name: string;
    messages: Array<IBuildMessage>;
}

export interface IBuildStatus {
    jobs: Array<IBuildJob>;
}
