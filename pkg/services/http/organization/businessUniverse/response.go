package businessUniverseHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
)

type CreateResponse struct {
	Code             int32                `json:"code"`
	Message          string               `json:"message"`
	BusinessUniverse BusinessUniverseType `json:"businessUniverse"`
}

type BusinessUniverseType struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type GetResponse struct {
	Code             int32                `json:"code"`
	Message          string               `json:"message"`
	BusinessUniverse BusinessUniverseType `json:"businessUniverse"`
}

type UpdateResponse struct {
	Code             int32                `json:"code"`
	Message          string               `json:"message"`
	BusinessUniverse BusinessUniverseType `json:"businessUniverse"`
}

func toCreateResponse(r *org.BusinessUniverseCreateResponse) *CreateResponse {
	return &CreateResponse{
		Code:    r.Code,
		Message: r.Message,
		BusinessUniverse: BusinessUniverseType{
			Name:   r.BusinessUniverse.Name,
			Domain: r.BusinessUniverse.Domain,
		},
	}
}

func toGetResponse(r *org.BusinessUniverseGetResponse) *GetResponse {
	return &GetResponse{
		Code:    r.Code,
		Message: r.Message,
		BusinessUniverse: BusinessUniverseType{
			Name:   r.BusinessUniverse.Name,
			Domain: r.BusinessUniverse.Domain,
		},
	}
}

func toUpdateResponse(r *org.BusinessUniverseUpdateResponse) *UpdateResponse {
	return &UpdateResponse{
		Code:    r.Code,
		Message: r.Message,
		BusinessUniverse: BusinessUniverseType{
			Name:   r.BusinessUniverse.Name,
			Domain: r.BusinessUniverse.Domain,
		},
	}
}
