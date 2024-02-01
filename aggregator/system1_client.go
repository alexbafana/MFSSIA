package aggregator

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type System1Client struct {
	aggregator *Aggregator
}

func (client *System1Client) getContractHash(contractId string) string {
	s1Url := "http://ec2-44-213-31-19.compute-1.amazonaws.com:4000/v1/api/business-contract/" + contractId + "/system_1-hash"

	resp, err := http.Get(s1Url)

	if err != nil {
		client.aggregator.logger.Error("Error calling system1 api", err)
		os.Exit(1)
	}

	jsonBody, err := io.ReadAll(resp.Body)

	if err != nil {
		client.aggregator.logger.Error("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var hashResponse SystemHashResponse

	errM := json.Unmarshal(jsonBody, &hashResponse)

	if errM != nil {
		client.aggregator.logger.Error("Error when unmarshalling response", errM)
	}

	return hashResponse.Data.Hash
}
