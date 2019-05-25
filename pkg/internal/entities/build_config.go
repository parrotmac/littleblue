package entities

type BuildConfigurationService interface {
	CreateBuildConfiguration(c *BuildConfiguration) error
}

type BuildConfiguration struct {
	ID                 uint     `json:"id"`
	SourceRepositoryID uint     `json:"source_repo_id"`
	DockerfileName     string   `json:"dockerfile_name"`
	HostBuildOS        string   `json:"host_build_os"`   // e.g. "windows" or "linux"
	HostBuildArch      string   `json:"host_build_arch"` // e.g. "arm64" or "armhf"
	TaggingRules       []string `json:"tagging_rules"`   // set of rules to determine how builds are tagged
}
