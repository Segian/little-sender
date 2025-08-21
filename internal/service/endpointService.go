package service

import (
	"github.com/Segian/little-sender/internal/config"
	"github.com/Segian/little-sender/internal/model"
	"github.com/Segian/little-sender/internal/model/dto"
	"github.com/Segian/little-sender/internal/util"

	"github.com/google/uuid"
)

func fillEndpointModel(endpointModel *model.EndpointModel, endpointDto dto.EndpointDtoUpdate) {
	endpointModel.Name = util.StringIfNotEmpty(endpointDto.Name, endpointModel.Name)
	endpointModel.Headers = util.JsonIfNotEmpty(endpointDto.Headers, endpointModel.Headers)
	endpointModel.URL = util.StringIfNotEmpty(endpointDto.URL, endpointModel.URL)
	endpointModel.Method = util.StringIfNotEmpty(endpointDto.Method, endpointModel.Method)
	if endpointDto.Body != nil {
		endpointModel.Body = &model.BodyModel{Data: endpointDto.Body}
	} else {
		endpointModel.Body = nil
	}
}

func endpointDtoToModel(endpointDto dto.EndpointDto) model.EndpointModel {
	endpoint := model.EndpointModel{
		ID:      uuid.New(),
		Name:    endpointDto.Name,
		Method:  endpointDto.Method,
		Headers: endpointDto.Headers,
		URL:     endpointDto.URL,
		Body:    nil, // Initialize Body as nil
	}
	if endpointDto.Body != nil {
		endpoint.Body = &model.BodyModel{Data: endpointDto.Body}
	}
	return endpoint
}

func endpointModelToDtoResponse(endpoint model.EndpointModel) dto.EndpointResponseDto {
	if endpoint.Body != nil {
		return dto.EndpointResponseDto{
			ID:      endpoint.ID,
			Name:    endpoint.Name,
			Method:  endpoint.Method,
			Headers: endpoint.Headers,
			URL:     endpoint.URL,
			Body:    endpoint.Body.Data,
		}
	}
	return dto.EndpointResponseDto{
		ID:      endpoint.ID,
		Name:    endpoint.Name,
		Method:  endpoint.Method,
		Headers: endpoint.Headers,
		URL:     endpoint.URL,
	}
}

func NewEndpointService(EndpointDto dto.EndpointDto) (dto.EndpointResponseDto, error) {

	endpointToAdd := endpointDtoToModel(EndpointDto)

	util.MethodToUpper(&endpointToAdd.Method)

	if err := config.DB.Create(&endpointToAdd).Error; err != nil {
		return dto.EndpointResponseDto{}, err
	}

	endpointResponse := endpointModelToDtoResponse(endpointToAdd)

	return endpointResponse, nil
}

func GetEndpoints() ([]dto.EndpointResponseDto, error) {
	var endpoints []model.EndpointModel
	if err := config.DB.Find(&endpoints).Error; err != nil {
		return nil, err
	}

	var endpointResponses []dto.EndpointResponseDto
	for _, endpoint := range endpoints {
		endpointResponses = append(endpointResponses, endpointModelToDtoResponse(endpoint))
	}

	return endpointResponses, nil
}

func GetEndpointByID(id uuid.UUID) (dto.EndpointResponseDto, error) {
	var endpoint model.EndpointModel
	if err := config.DB.First(&endpoint, id).Error; err != nil {
		return dto.EndpointResponseDto{}, err
	}
	if err := config.DB.Preload("Body").First(&endpoint, id).Error; err != nil {
		return dto.EndpointResponseDto{}, err
	}

	return endpointModelToDtoResponse(endpoint), nil
}

func UpdateEndpoint(id uuid.UUID, endpointDto dto.EndpointDtoUpdate) (dto.EndpointResponseDto, error) {
	var endpoint model.EndpointModel
	if err := config.DB.First(&endpoint, id).Error; err != nil {
		return dto.EndpointResponseDto{}, err
	}
	if endpointDto.Method != "" {
		util.MethodToUpper(&endpointDto.Method)
	}
	fillEndpointModel(&endpoint, endpointDto)

	if err := config.DB.Save(&endpoint).Error; err != nil {
		return dto.EndpointResponseDto{}, err
	}

	return endpointModelToDtoResponse(endpoint), nil
}

func DeleteEndpoint(id uuid.UUID) error {
	var endpoint model.EndpointModel
	if err := config.DB.First(&endpoint, id).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&endpoint).Error; err != nil {
		return err
	}
	return nil
}
