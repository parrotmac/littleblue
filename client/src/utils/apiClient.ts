import {IBuildJob, IRepo} from "./types";

class ApiClient {
    private baseUrl: string;
    constructor(baseUrl: string) {
        this.baseUrl = baseUrl;
    }

    public fetchRepos = (): Promise<Array<IRepo>> => {
        return new Promise<Array<IRepo>>(((resolve, reject) => {
            fetch(`${this.baseUrl}/api/repos/`).then(
                res => {
                    if (res.status === 200) {
                        res.json().then(resolve).catch(reject);
                    } else {
                        reject(res.statusText);
                    }
                }
            ).catch(reject);
        }));
    };

    public fetchJobsForRepo = (repoID: number): Promise<Array<IBuildJob>> => {
        return new Promise((resolve, reject) => {
            fetch(`${this.baseUrl}/api/repos/${repoID}/jobs/`).then(
                res => {
                    if (res.status === 200) {
                        res.json().then(resolve).catch(reject);
                    } else {
                        reject(res.statusText)
                    }
                }
            ).catch(reject)
        })
    };

    public fetchJobs = (): Promise<Array<IBuildJob>> => {
        return new Promise((resolve, reject) => {
            fetch(`${this.baseUrl}/api/jobs`).then(
                res => {
                    if (res.status === 200) {
                        res.json().then(resolve).catch(reject);
                    } else {
                        reject(res.statusText)
                    }
                }
            ).catch(reject)
        })
    };
}

export default ApiClient;
