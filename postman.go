package main

type Postman struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Order        []string      `json:"order"`
	Folders      []interface{} `json:"folders"`
	FoldersOrder []interface{} `json:"folders_order"`
	Timestamp    int64         `json:"timestamp"`
	Owner        int           `json:"owner"`
	Public       bool          `json:"public"`
	Requests     []struct {
		ID               string        `json:"id"`
		Headers          string        `json:"headers"`
		HeaderData       []interface{} `json:"headerData"`
		URL              string        `json:"url"`
		QueryParams      []interface{} `json:"queryParams"`
		PreRequestScript interface{}   `json:"preRequestScript"`
		PathVariables    struct {
		} `json:"pathVariables"`
		PathVariableData []interface{} `json:"pathVariableData"`
		Method           string        `json:"method"`
		Data             []struct {
			Key         string `json:"key"`
			Value       string `json:"value"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Enabled     bool   `json:"enabled"`
		} `json:"data"`
		DataMode         string      `json:"dataMode"`
		Version          int         `json:"version,omitempty"`
		Tests            interface{} `json:"tests"`
		CurrentHelper    string      `json:"currentHelper"`
		HelperAttributes struct {
		} `json:"helperAttributes"`
		Time         int64         `json:"time"`
		Name         string        `json:"name"`
		Description  string        `json:"description"`
		CollectionID string        `json:"collectionId"`
		Responses    []interface{} `json:"responses"`
	} `json:"requests"`
}
