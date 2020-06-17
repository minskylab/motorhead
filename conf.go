package motorhead

// ConfigFile represent one config file that describe a resource (aka Object)
type ConfigFile struct {
	APIVersion string `yaml:"apiVersion"`
	Kind Kind `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
}
