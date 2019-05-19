import { IBuildStatus } from "./types";

class ApiClient {
    private baseUrl: string;
    constructor(baseUrl: string) {
        this.baseUrl = baseUrl;
    }
    
    public fetchJobs = (): Promise<IBuildStatus> => {
        return new Promise((resolve, reject) => {
            fetch(`${this.baseUrl}/api/jobs`).then(
                res => {
                    if (res.status === 200) {
                        res.json().then(
                            jobs => {
                                const buildStats: IBuildStatus = {
                                  jobs,
                                };
                                resolve(buildStats);
                            }
                        ).catch(reject);
                    } else {
                        reject(res.statusText)
                    }
                }
            ).catch(reject)
        })
    }
}

export default ApiClient;
