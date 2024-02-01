package aggregator

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
)

type DkgClient struct {
	aggregator *Aggregator
}

func (client *DkgClient) getContractHash(contractId string) string {

	dkgUrl := "https://ec2-3-123-233-195.eu-central-1.compute.amazonaws.com/api/consensus/contractHash/" + contractId

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	resp, err := httpClient.Get(dkgUrl)

	if err != nil {
		client.aggregator.logger.Error("Error calling dkg api", err)
		os.Exit(1)
	}

	dkgHash, err := io.ReadAll(resp.Body)

	if err != nil {
		client.aggregator.logger.Error("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return string(dkgHash)
}
