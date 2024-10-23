package sitecore

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func RunQuery(query string) map[string]interface{} {
	var headers map[string]string
	var params map[string]string
	return RunQueryWithParameters(query, "", headers, params)
}

func RunQueryWithParameters(query string, queryName string, headers map[string]string, params map[string]string) map[string]interface{} {

	jsonMapInstance := map[string]interface{}{
		"query":         query,
		"operationName": queryName,
		"variables":     params,
	}

	jsonData, err := json.Marshal(jsonMapInstance)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", GetEnvVar("GRAPHQL_ENDPOINT"), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sc_apikey", GetEnvVar("SITECORE_API_KEY"))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

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
