package goconda

type verb string

const (
	Install verb = "install"
	Download
)

type conda struct {
	verb         verb
	packageInfo  PackageInfo
	dryRun       bool
	downloadPath string // used in CONDA_PKGS_DIRS.
	condaPath    string // used in CONDA_ENVS_PATH.
}

type command struct {
	binaryName      *string
	binaryArguments []string
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
func New(condaPath string) *conda {
	return &conda{
		condaPath: condaPath,
	}
}

// Build the Conda command so it can be run.
func (p *conda) Seal() (*command, error) {
	return nil, nil
}
