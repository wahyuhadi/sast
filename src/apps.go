package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/google/uuid"
	"github.com/wahyuhadi/supply-chain/helpers"
	"github.com/wahyuhadi/supply-chain/models"
)

var (
	branch = flag.String("branch", "main", "branch")
	rules  = flag.String("rules", "auto", "rules links")
	lang   = flag.String("lang", "", "lang")
	mode   = flag.String("mode", "-q", "mode to scaning ")
	key    = flag.String("key", "", "unique key")

	repo    = flag.String("repo", "", "mode to scaning ")
	private = flag.Bool("private", true, "is private repo ")
	gituser = flag.String("guser", "", "git user ")
	gpass   = flag.String("gpass", "", "git token ")

	config  = flag.String("config", "/root/config.yaml", "config file location")
	gitconf = flag.String("git-config", "", "use private repo to clone")
)

func parse_opt() *models.Optional {
	flag.Parse()
	return &models.Optional{
		Branch:    *branch,
		Key:       *key,
		Private:   *private,
		Rules:     *rules,
		Username:  *gituser,
		Password:  *gpass,
		Lang:      *lang,
		Mode:      *mode,
		Conf:      *config,
		ConfScan:  nil,
		GitConf:   *gitconf,
		FileSaved: "/tmp/finding",
		RepoURI:   *repo,
	}
}

func init() {
	if !check_semgrep() {
		log.Println("[!] Please install semgrep : pip3 install semgrep")
		os.Exit(1)
	}

}

func main() {
	config := parse_opt()
	if config.Conf != "" {
		conf, err := helpers.ConfParser(config.Conf)
		if err != nil {
			log.Println("[!] Error read config file ")
			os.Exit(1)
		}
		config.ElasticHost = conf.Elasticsearch.IP
		config.ElasticIndex = "code-scanning"
		config.ElasticPort = conf.Elasticsearch.Port
		config.ElasticPassword = conf.Elasticsearch.Password
		config.ElasticUser = conf.Elasticsearch.Username
		config.UniqueID = uuid.New().String()
	}
	helpers.Scan(*config)
}

func check_semgrep() bool {
	_, err1 := exec.LookPath("semgrep")
	if err1 != nil {
		return false
	}
	return true
}
