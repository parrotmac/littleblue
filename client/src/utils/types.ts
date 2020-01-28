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

export interface IRepo {
    id: number
    source_provider_id: number
    repo_uuid: string
    name: string
}

export interface ILogs {
    source: string[] | null;
    build: string[] | null;
    push: string[] | null;
}

export interface IBuildJob {
        id: number;
        build_identifier: string;
        build_config_id: number;
        start_time: Date;
        end_time: Date | null;
        status: string;
        failure: boolean;
        failure_detail: string | null;
        build_host: string | null;
        source_reference: string;
        source_revision: string | null;
        source_uri: string;
        artifact_uri: string;
        logs: ILogs;
    }

