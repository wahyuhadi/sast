package helpers

import (
	"log"
	"os"

	"github.com/wahyuhadi/supply-chain/models"
	"gopkg.in/yaml.v3"
)

func ConfParser(file string) (models.DB, error) {
	yfile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	conf := &models.DB{}
	err = yaml.Unmarshal(yfile, conf)
	if err != nil {
		log.Fatal(err)
	}

	return *conf, err
}
