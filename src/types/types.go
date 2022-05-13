package types

// ZarfFile defines a file to deploy.
type ZarfFile struct {
	Source     string   `yaml:"source"`
	Shasum     string   `yaml:"shasum,omitempty"`
	Target     string   `yaml:"target"`
	Executable bool     `yaml:"executable,omitempty"`
	Symlinks   []string `yaml:"symlinks,omitempty"`
}

// ZarfChart defines a helm chart to be deployed.
type ZarfChart struct {
	Name        string   `yaml:"name"`
	Url         string   `yaml:"url"`
	Version     string   `yaml:"version"`
	Namespace   string   `yaml:"namespace"`
	ValuesFiles []string `yaml:"valuesFiles,omitempty"`
	GitPath     string   `yaml:"gitPath,omitempty"`
}

// ZarfComponent is the primary functional grouping of assets to deploy by zarf.
type ZarfComponent struct {
	// Name is the unique identifier for this component
	Name string `yaml:"name"`

	// Description is a message given to a user when deciding to enable this componenent or not
	Description string `yaml:"description,omitempty"`

	// Default changes the default option when deploying this component
	Default bool `yaml:"default,omitempty"`

	// Required makes this component mandatory for package deployment
	Required bool `yaml:"required,omitempty"`

	// SecretName is the secret zarf will use for the registry, the default is "zarf-registry"
	SecretName string `yaml:"secretName,omitempty"`

	// Files are files to place on disk during deploy
	Files []ZarfFile `yaml:"files,omitempty"`

	// Charts are helm charts to install during package deploy
	Charts []ZarfChart `yaml:"charts,omitempty"`

	// Manifests are raw manifests that get converted into zarf-generated helm charts during deploy
	Manifests []ZarfManifest `yaml:"manifests,omitempty"`

	// Images are the online images needed to be included in the zarf package
	Images []string `yaml:"images,omitempty"`

	// Repos are any git repos that need to be pushed into the gitea server
	Repos []string `yaml:"repos,omitempty"`

	// Data pacakges to push into a running cluster
	DataInjections []ZarfDataInjection `yaml:"dataInjections,omitempty"`

	// Scripts are custom commands that run before or after package deployment
	Scripts ZarfComponentScripts `yaml:"scripts,omitempty"`

	// Import refers to another zarf.yaml package component.
	Import ZarfComponentImport `yaml:"import,omitempty"`

	//Path to cosign publickey for signed online resources
	CosignKeyPath string `yaml:"cosignKeyPath,omitempty"`

	// Dynamic template values for K8s resources
	Variables ZarfComponentVariables `yaml:"variables,omitempty"`
}

// ZarfComponentVariables are variables that can be used to dynaically template K8s resources
type ZarfComponentVariables map[string]string

// ZarfManifest defines raw manifests Zarf will deploy as a helm chart
type ZarfManifest struct {
	Name                       string   `yaml:"name"`
	DefaultNamespace           string   `yaml:"namespace,omitempty"`
	Files                      []string `yaml:"files,omitempty"`
	KustomizeAllowAnyDirectory bool     `yaml:"kustomizeAllowAnyDirectory,omitempty"`
	Kustomizations             []string `yaml:"kustomizations,omitempty"`
}

// ZarfComponentScripts are scripts that run before or after a component is deployed
type ZarfComponentScripts struct {
	ShowOutput     bool     `yaml:"showOutput,omitempty"`
	TimeoutSeconds int      `yaml:"timeoutSeconds,omitempty"`
	Retry          bool     `yaml:"retry,omitempty"`
	Before         []string `yaml:"before,omitempty"`
	After          []string `yaml:"after,omitempty"`
}

// ZarfMetadata lists information about the current ZarfPackage
type ZarfMetadata struct {
	Name         string `yaml:"name,omitempty"`
	Description  string `yaml:"description,omitempty"`
	Version      string `yaml:"version,omitempty"`
	Url          string `yaml:"url,omitempty"`
	Image        string `yaml:"image,omitempty"`
	Uncompressed bool   `yaml:"uncompressed,omitempty"`
	Architecture string `yaml:"architecture,omitempty"`
}

// ZarfContainerTarget defines the destination info for a ZarfData target
type ZarfContainerTarget struct {
	Namespace string `yaml:"namespace"`
	Selector  string `yaml:"selector"`
	Container string `yaml:"container,omitempty"`
	Path      string `yaml:"path"`
}

// ZarfDataInjection is a data-injection definition
type ZarfDataInjection struct {
	Source string              `yaml:"source"`
	Target ZarfContainerTarget `yaml:"target"`
}

// ZarfBuildData is written during the packager.Create() operation to track details of the created package
type ZarfBuildData struct {
	Terminal     string `yaml:"terminal"`
	User         string `yaml:"user"`
	Architecture string `yaml:"architecture"`
	Timestamp    string `yaml:"timestamp"`
	Version      string `yaml:"string"`
}

// ZarfPackage the top-level structure of a Zarf config file
type ZarfPackage struct {
	Kind       string          `yaml:"kind,omitempty"`
	Metadata   ZarfMetadata    `yaml:"metadata,omitempty"`
	Build      ZarfBuildData   `yaml:"build,omitempty"`
	Components []ZarfComponent `yaml:"components,omitempty"`
	Seed       string          `yaml:"seed,omitempty"`
}

// ZarfState is maintained as a secret in the Zarf namespace to track Zarf init data
type ZarfState struct {
	ZarfAppliance bool   `json:"zarfAppliance"`
	Distro        string `json:"distro"`
	Architecture  string `json:"architecture"`
	StorageClass  string `json:"storageClass"`
	Secret        string `json:"secret"`
	NodePort      string `json:"nodePort"`
}

// TLSConfig tracks the user-defined options for TLS cert generation
type TLSConfig struct {
	CertPublicPath  string `yaml:"certPublicPath"`
	CertPrivatePath string `yaml:"certPrivatePath"`
	Host            string `yaml:"host"`
}

// ZarfDeployOptions tracks the user-defined preferences during a package deployment
type ZarfDeployOptions struct {
	PackagePath string
	Confirm     bool
	Components  string

	// Zarf init is installing the k3s component
	ApplianceMode bool

	// Zarf init override options
	StorageClass string
	Secret       string
	NodePort     string
}

// ZarfImport structure for including imported zarf components
type ZarfComponentImport struct {
	ComponentName string `yaml:"name,omitempty"`
	Path          string `yaml:"path"`
}
