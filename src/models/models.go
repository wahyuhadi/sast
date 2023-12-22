package models

type Optional struct {
	Dir             string
	Rules           string
	Lang            string
	Mode            string
	Conf            string
	ConfScan        *ConfigScan
	GitConf         string
	FileSaved       string
	RepoURI         string
	Username        string
	Password        string
	Private         bool
	ElasticUser     string
	ElasticPassword string
	ElasticIndex    string
	ElasticPort     int
	ElasticHost     string
	UniqueID        string
	Key             string
	Branch          string
}

type ConfigScan struct {
	Lang     string   `yaml:"lang"`
	Rules    string   `yaml:"rules"`
	ScanType string   `yaml:"scan_type"`
	Projects []string `yaml:"projects"`
}

type DB struct {
	Elasticsearch struct {
		IP        string `yaml:"ip"`
		Port      int    `yaml:"port"`
		IndexName string `yaml:"index-name"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
	} `yaml:"elasticsearch"`
}
