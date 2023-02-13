package vo

type Config struct {
	NginxTemplate NginxTemplate `yaml:"template"`
	SshConfigs    []SshConfig   `yaml:"ssh-configs"`
}

type SshConfig struct {
	Host       string `yaml:"host"`
	Port       int64  `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Type       string `yaml:"type"`
	RsaPath    string `yaml:"rsa-path"`
	TargetPath string `yaml:"target-path"`
	NginxPath  string `yaml:"nginx-path"`
}

type NginxTemplate struct {
	Name   string `yaml:"name"`
	Header string `yaml:"header"`
	Footer string `yaml:"footer"`
}
