package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Segian/little-sender/internal/config"
	"github.com/Segian/little-sender/internal/model"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func ExecuteEndpoint(endpointID uuid.UUID) (model.HistoryModel, error) {
	var endpoint model.EndpointModel
	if err := config.DB.First(&endpoint, endpointID).Error; err != nil {
		return model.HistoryModel{}, err
	}
	var body model.BodyModel
	if err := config.DB.Preload("Body").First(&endpoint, endpointID).Error; err != nil {
		return model.HistoryModel{}, err
	}

	client := &http.Client{}
	var reqBody *bytes.Reader

	if (endpoint.Method == http.MethodPost || endpoint.Method == http.MethodPut) && body.Data != nil {
		reqBody = bytes.NewReader(body.Data)
	} else {
		reqBody = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(endpoint.Method, endpoint.URL, reqBody)
	if err != nil {
		return model.HistoryModel{}, err
	}

	var headers map[string]string
	if err := json.Unmarshal(endpoint.Headers, &headers); err != nil {
		return model.HistoryModel{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	timeStart := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return model.HistoryModel{}, err
	}
	defer resp.Body.Close()
	timeTaken := time.Since(timeStart)
	responseTime := timeTaken.Milliseconds()

	if resp.StatusCode != http.StatusOK {
		return model.HistoryModel{}, fmt.Errorf("failed to execute endpoint: %s", resp.Status)
	}

	var responseBody datatypes.JSON
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return model.HistoryModel{}, err
	}

	var addHistory = model.HistoryModel{
		ID:           uuid.New(),
		EndpointID:   endpoint.ID,
		StatusCode:   resp.StatusCode,
		Response:     responseBody,
		Timestamp:    time.Now(),
		ResponseTime: responseTime,
	}

	if config.DB.Create(&addHistory).Error != nil {
		return model.HistoryModel{}, fmt.Errorf("failed to create history record: %w", err)
	}

	return addHistory, nil
}
