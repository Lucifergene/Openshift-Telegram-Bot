package login

type Config struct {
	APIVersion     string      `yaml:"apiVersion"`
	Clusters       []Cluster   `yaml:"clusters"`
	Contexts       []Context   `yaml:"contexts"`
	CurrentContext string      `yaml:"current-context"`
	Kind           string      `yaml:"kind"`
	Preferences    interface{} `yaml:"preferences"`
	Users          []User      `yaml:"users"`
}

type Cluster struct {
	Cluster ClusterData `yaml:"cluster"`
	Name    string      `yaml:"name"`
}

type ClusterData struct {
	InsecureSkipTLSVerify bool   `yaml:"insecure-skip-tls-verify"`
	Server                string `yaml:"server"`
}

type Context struct {
	Context ContextData `yaml:"context"`
	Name    string      `yaml:"name"`
}

type ContextData struct {
	Cluster   string `yaml:"cluster"`
	Namespace string `yaml:"namespace"`
	User      string `yaml:"user"`
}

type User struct {
	Name string   `yaml:"name"`
	User UserData `yaml:"user"`
}

type UserData struct {
	Token string `yaml:"token"`
}
