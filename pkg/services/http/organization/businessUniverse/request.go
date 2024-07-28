package businessUniverseHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
)

type GetRequest struct {
	Id string `json:"id"`
}

type UpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func toGetRequest(r *GetRequest) *org.BusinessUniverseGetRequest {
	return &org.BusinessUniverseGetRequest{
		Id: r.Id,
	}
}
