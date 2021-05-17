package config

type varOptions struct {
	VarsFile     string       // Required to initialize the sdk global config object
	Verbose      bool         // Recommended for flag handling
	ServicePacks servicePacks `yaml:"ServicePacks"` // Optional
}

// servicePacks is only required if this pack accepts custom vars
type servicePacks struct {
	Wireframe wireframe `yaml:"Wireframe"`
}

// wireframe defines the custom vars for this service pack
type wireframe struct {
	Pass          string   `yaml:"Pass"`
	TagInclusions []string `yaml:"TagInclusions"`
	TagExclusions []string `yaml:"TagExclusions"`
}
