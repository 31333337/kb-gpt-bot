package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

var Client *elasticsearch.Client

func Initialize() {
	var err error
	Client, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error initializing Elasticsearch client: %s", err)
	}
}

// Add more functions to interact with Elasticsearch.

