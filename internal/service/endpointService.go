package service

import (
	"github.com/Segian/little-sender/internal/model"
)

func NewEndpointService(endpointModel model.EndpointModel) model.EndpointModel {

	var endpointToAdd model.EndpointModel
	endpointToAdd.ID = endpointModel.ID
	endpointToAdd.Name = endpointModel.Name
	endpointToAdd.Headers = endpointModel.Headers
	endpointToAdd.URL = endpointModel.URL

	if endpointModel.Bodies != nil {
		endpointToAdd.Bodies = make([]model.BodyModel, len(endpointModel.Bodies))
		for i, body := range endpointModel.Bodies {
			endpointToAdd.Bodies[i] = model.BodyModel{
				ID:         body.ID,
				EndpointID: body.EndpointID,
				Body:       body.Body,
			}
		}
	}

	if endpointToAdd.Bodies == nil {
		endpointToAdd.Bodies = []model.BodyModel{} // Ensure it's initialized even if empty
	}

	return endpointToAdd
}
