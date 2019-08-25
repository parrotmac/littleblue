package source

type GitClient interface {
	CloneDefaultTo(targetDirectory string) error
	CloneBranchTo(branchRef string, targetDirectory string) error
	CloneRevisionTo(revision string, targetDirectory string) error
}
