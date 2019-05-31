// +build linux

package builder

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/cli/cli/command/image/build"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/idtools"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

func LocalPathToBuildContextTar(contextPath string, config *entities.BuildConfiguration) (io.ReadCloser, error) {
	/*
		Much of this is pulled from directly from Docker
	*/
	var (
		dockerfileCtx io.ReadCloser
		contextDir    string
		relDockerfile string
	)

	contextDir, relDockerfile, err := build.GetContextFromLocalDir(contextPath, config.DockerfileName)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(relDockerfile, ".."+string(filepath.Separator)) {
		// Dockerfile is outside of build-context; read the Dockerfile and pass it as dockerfileCtx
		dockerfileCtx, err = os.Open(config.DockerfileName)
		if err != nil {
			return nil, errors.Wrap(err, "unable to open Dockerfile")
		}
		defer func() {
			err = dockerfileCtx.Close()
			if err != nil {
				logrus.Warnln(err)
			}
		}()
	}

	excludes, err := build.ReadDockerignore(contextDir)
	if err != nil {
		return nil, err
	}

	if err := build.ValidateContextDirectory(contextDir, excludes); err != nil {
		return nil, errors.Errorf("error checking context: '%s'.", err)
	}

	// And canonicalize dockerfile name to a platform-independent one
	relDockerfile = archive.CanonicalTarNameForPath(relDockerfile)
	excludes = build.TrimBuildFilesFromExcludes(excludes, relDockerfile, false)

	return archive.TarWithOptions(contextDir, &archive.TarOptions{
		ExcludePatterns: excludes,
		ChownOpts:       &idtools.Identity{UID: 0, GID: 0},
	})
}
