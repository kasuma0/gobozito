package controller

import "github.com/kasuma0/gobozito/model"

func PingController(request *model.InteractionRequest) interface{} {
	if request.Type == 1 {
		return &model.PingResponse{Type: 1}
	}
	return nil
}
