package service

import (
	"github.com/Segian/little-sender/internal/config"
	"github.com/Segian/little-sender/internal/dto"
	"github.com/Segian/little-sender/internal/model"

	"github.com/google/uuid"
)

func NewEndpointService(EndpointDto dto.EndpointDto) (dto.EndpointResponseDto, error) {

	var endpointToAdd model.EndpointModel

	endpointToAdd.ID = uuid.New()
	endpointToAdd.Name = EndpointDto.Name
	endpointToAdd.Headers = EndpointDto.Headers
	endpointToAdd.URL = EndpointDto.URL

	if EndpointDto.Body != nil {
		endpointToAdd.Body = &model.BodyModel{
			Data: EndpointDto.Body,
		}
	}

	if err := config.DB.Create(&endpointToAdd).Error; err != nil {
		return dto.EndpointResponseDto{}, err
	}

	var endpointResponse dto.EndpointResponseDto
	endpointResponse.ID = endpointToAdd.ID
	endpointResponse.Name = endpointToAdd.Name
	endpointResponse.Headers = endpointToAdd.Headers
	endpointResponse.URL = endpointToAdd.URL
	endpointResponse.Body = endpointToAdd.Body.Data

	return endpointResponse, nil
}
