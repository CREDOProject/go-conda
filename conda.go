package goconda

import "errors"

type verb string

const (
	Install  verb = "install"
	Download      = "download"
)

var (
	ErrNoPackageName = errors.New("Package name not specified.")
	ErrNoVerb        = errors.New("Verb not specified.")
)

type conda struct {
	binaryName   string
	condaPath    string // used in CONDA_ENVS_PATH.
	downloadPath string // used in CONDA_PKGS_DIRS.
	dryRun       bool
	packageInfo  PackageInfo
	verb         verb
}

type command struct {
	binaryName      *string
	binaryArguments []string
	env             map[string]string
}

type PackageInfo struct {
	PackageName string
	Channel     string
}

// Installs a package.
// https://docs.anaconda.com/free/working-with-conda/packages/install-packages/
func (c *conda) Install(info *PackageInfo) *conda {
	c.packageInfo = *info
	c.verb = Install
	return c
}

// Downloads a package.
// https://docs.anaconda.com/free/working-with-conda/packages/shared-pkg-cache/
func (c *conda) Download(info *PackageInfo, downloadPath string) *conda {
	c.packageInfo = *info
	c.verb = Download
	c.downloadPath = downloadPath
	return c
}

// Enables dry-run.
// https://docs.conda.io/projects/conda/en/latest/commands/install.html#output,-prompt,-and-flow-control-options
func (c *conda) DryRun() *conda {
	c.dryRun = true
	return c
}

// Start a new Conda command.
// https://docs.anaconda.com/free/working-with-conda/
func New(binaryName string, downloadPath string, condaPath string) *conda {
	return &conda{
		binaryName:   binaryName,
		condaPath:    condaPath,
		downloadPath: downloadPath,
	}
}

// Build the Conda command so it can be run.
func (p *conda) Seal() (*command, error) {
	args := []string{}

	switch p.verb {
	case Install:
		args = append(args, "install")
	case Download:
		args = append(args, "install", "--download-only")
	default:
		return nil, errors.New("No verb specified.")
	}

	if p.packageInfo.PackageName == "" {
		return nil, errors.New("No package name specified.")
	}

	args = append(args, p.packageInfo.PackageName)

	if p.packageInfo.Channel != "" {
		args = append(args, "--channel", p.packageInfo.Channel)
	}

	if p.dryRun {
		args = append(args, "--dry-run")
	}

	env := make(map[string]string)

	if p.downloadPath != "" {
		env["CONDA_PKGS_DIRS"] = p.downloadPath
	}
	if p.condaPath != "" {
		env["CONDA_ENVS_PATH"] = p.condaPath
	}

	return &command{
		binaryName:      &p.binaryName,
		binaryArguments: args,
		env:             env,
	}, nil
}
