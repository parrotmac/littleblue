package entities

import "time"

type BuildJobService interface {
	CreateBuildJob(c *BuildJob) error
}

type BuildJob struct {
	ID uint `json:"id"`

	// UUID for build
	BuildIdentifier string `json:"build_identifier"`

	/*
		Reference to build configuration
	*/
	BuildConfigurationID uint `json:"build_config_id"`

	/*
		When job started (pulled from record's `created_at` column) and finished
	*/
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	/*
		Note: Status isn't cleared on failure
		unknown: Status is unknown
		created: Job created
		pending: Builder assigned
		cloning: Fetching source
		building: Build in progress
		pushing: Artifacts are being uploaded
		complete: Process finished
	*/
	Status string `json:"status"`

	/*
		Was an error encountered?
		If true, this is combined with the status field to create a failure type.
		If there are failure details, they can be found in the failure_detail field
	*/
	Failed bool `json:"failure"`

	// Details of encountered error
	FailureDetail string `json:"failure_detail"`

	// Name or address of machine performing build
	BuildHost string `json:"build_host"`

	// git (or other) ref
	SourceUri string `json:"source_uri"`

	// Docker image
	ArtifactUri string `json:"artifact_uri"`

	// Logs from different stages
	Logs BuildLogs `json:"logs"`
}

type BuildLogs struct {
	// Logs from getting source
	SetupLogs []string `json:"source"`

	// Build output
	BuildLogs []string `json:"build"`

	// Push log
	PushLogs []string `json:"push"`
}
