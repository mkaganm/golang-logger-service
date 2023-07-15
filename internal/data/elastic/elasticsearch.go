package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"logger_service/internal/config"
	"logger_service/internal/utils"
	"net/http"
)

type LogData struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	App       string `json:"app"`
	Host      string `json:"host"`
	Env       string `json:"env"`
}

// SendToElasticsearch sends log to Elasticsearch
func SendToElasticsearch(logData LogData) {
	logJSON, err := json.Marshal(logData)
	if err != nil {
		fmt.Println("Failed to marshal log data:", err)
		return
	}

	resp, err := http.Post(config.EnvConfigs.ElasticUrl, "application/json", bytes.NewBuffer(logJSON))
	utils.CheckErr("Failed to send log to Elasticsearch:", err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.CheckErr("Failed to close response body:", err)
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Failed to send log to Elasticsearch. Status code:", resp.StatusCode)
	}

}
