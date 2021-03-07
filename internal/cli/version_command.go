package cli

import (
	"github.com/cbuschka/versioner/internal/build"
	"runtime"
)

// VersionCommandConfig is the version command config.
type VersionCommandConfig struct {
	Long bool
}

// Run executes the version command.
func (config *VersionCommandConfig) Run() error {
	buildInfo := build.GetBuildInfo()
	if !config.Long {
		Print("%s", buildInfo.Version)
	} else {
		Print(`versioner version is %s.
Built from commitish %s, on %s for %s/%s.
OS is %s, arch is %s.`,
			buildInfo.Version,
			buildInfo.Commitish, buildInfo.BuildTime, buildInfo.OS, buildInfo.Arch,
			runtime.GOOS, runtime.GOARCH)
	}

	return nil
}
