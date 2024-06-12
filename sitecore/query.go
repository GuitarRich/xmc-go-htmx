package sitecore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func RunQuery(query string) map[string]interface{} {

	jsonMapInstance := map[string]string{
		"query": query,
	}

	jsonData, err := json.Marshal(jsonMapInstance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running query")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", GetEnvVar("GRAPHQL_ENDPOINT"), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sc_apikey", GetEnvVar("SITECORE_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
