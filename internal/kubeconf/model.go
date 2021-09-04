package kubeconf

type UserData struct {
	ClientCertificateData string `yaml:"client-certificate-data"`
	ClientKeyData         string `yaml:"client-key-data"`
	ClientCertificate     string `yaml:"client-certificate"`
	ClientKey             string `yaml:"client-key"`
	Token                 string `yaml:"token"`
}

type User struct {
	Name string   `yaml:"name"`
	User UserData `yaml:"user"`
}

type ClusterData struct {
	CertificateAuthority     string `yaml:"certificate-authority"`
	CertificateAuthorityData string `yaml:"certificate-authority-data"`
	Server                   string `yaml:"server"`
	InsecureSkipTLSVerify    bool   `yaml:"insecure-skip-tls-verify"`
}

type Cluster struct {
	Name    string      `yaml:"name"`
	Cluster ClusterData `yaml:"cluster"`
}

type ContextData struct {
	User    string `yaml:"user"`
	Cluster string `yaml:"cluster"`
}

type Context struct {
	Name    string      `yaml:"name"`
	Context ContextData `yaml:"context"`
}

type KubeConfig struct {
	APIVersion     string                 `yaml:"apiVersion"`
	Kind           string                 `yaml:"kind"`
	Preferences    map[string]interface{} `yaml:"preferences"`
	Clusters       []Cluster              `yaml:"clusters"`
	Users          []User                 `yaml:"users"`
	Contexts       []Context              `yaml:"contexts"`
	CurrentContext string                 `yaml:"current-context"`
}

type Config struct {
	Path       string
	KubeConfig KubeConfig
}
