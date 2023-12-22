package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/wahyuhadi/supply-chain/models"
)

var (
	API   = "http://210.247.248.219:8080/api/sca"
	TOKEN = "a"
)

func ExtractResult(conf models.Optional) {
	date := time.Now().Format("January 02, 2006 15:04:05")
	time := time.Now().String() // number of seconds since January 1, 1970 UTC

	content, err := os.ReadFile(conf.FileSaved)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	// Now let's unmarshall the data into `payload`
	var payload models.Results
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	if len(payload.Results) == 0 {
		var finding models.Finding
		finding.ScanAt = time
		finding.Branch = conf.Branch
		finding.Key = conf.Key
		finding.UniqueID = conf.UniqueID
		finding.IsFinding = false
		push_data_scanning(finding, conf)
	} else {
		for _, i := range payload.Results {
			var finding models.Finding
			finding.Branch = conf.Branch
			finding.IsFinding = true
			finding.ScanAt = time
			finding.Key = conf.Key
			finding.UniqueID = conf.UniqueID
			finding.CheckID = i.CheckID
			finding.Confidence = i.Extra.Metadata.Confidence
			finding.EndCol = i.End.Col
			finding.EndLine = i.End.Line
			finding.EndOffset = i.End.Offset
			finding.EngineKind = i.Extra.EngineKind
			finding.Fingerprint = i.Extra.Fingerprint
			finding.IsIgnored = i.Extra.IsIgnored
			finding.Lines = i.Extra.Lines
			finding.Message = i.Extra.Message
			finding.Path = i.Path
			finding.Severity = i.Extra.Severity
			finding.StartCol = i.Start.Col
			finding.StartLine = i.Start.Line
			finding.StartOffset = i.Start.Offset
			if len(i.Extra.Metadata.Technology) == 0 {
				finding.Technology = nil
			} else {
				finding.Technology = &i.Extra.Metadata.Technology[0]
			}
			finding.ValidationState = i.Extra.ValidationState
			finding.Project = conf.Dir
			finding.Date = date

			push_data_scanning(finding, conf)
		}
	}

	os.Remove(conf.FileSaved)
}

func push_data_patrol(data models.Finding) error {
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", API, bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", TOKEN)

	client := &http.Client{}
	res, err := client.Do(r)
	if res.StatusCode != 201 {
		fmt.Println(fmt.Sprintf("PUSH DATA WITH SEVERITY %s SUCCESS", data.Severity))
		return nil
	}
	return err
}

func push_data_scanning(data models.Finding, conf models.Optional) error {
	log.Println("Push data to elastic")
	URI := fmt.Sprintf("%s://%s:%v", "http", conf.ElasticHost, conf.ElasticPort)
	cfg := elasticsearch.Config{
		Addresses: []string{
			URI,
		},
		Username: conf.ElasticUser,
		Password: conf.ElasticPassword,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println(err)

		return errors.New("error connection")
	}

	res, err := es.Index(conf.ElasticIndex, esutil.NewJSONReader(&data))
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println(res.StatusCode)

	defer res.Body.Close()
	log.Println(res)

	return nil
}
