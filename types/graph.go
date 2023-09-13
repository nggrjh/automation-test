package types

type GraphError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

type GraphResponse struct {
	Errors []GraphError `json:"errors"`
	Data   any          `json:"data"`
}
