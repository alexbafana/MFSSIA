package aggregator

type SystemHashResponse struct {
	Success    bool   `json:"success"`
	StatusCode int64  `json:"statusCode"`
	Message    string `json:"message"`
	Data       Data   `json:"data"`
	Links      any    `json:"linkw"`
}

type Data struct {
	Hash string `json:"hash"`
}
