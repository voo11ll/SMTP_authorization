package organization

import (
	"auth/auth_back/pkg/globalvars"
	context "context"

	uuid "github.com/google/uuid"
)

func (s *GrpcServer) CreateBusinessUniverse(ctx context.Context, in *BusinessUniverseCreateRequest) (*BusinessUniverseCreateResponse, error) {

	item, message, err := s.BusinessUniverseRepo.CreateItem(context.TODO())

	if err != nil && message == "" {
		return &BusinessUniverseCreateResponse{
			Code:             globalvars.ServerInternalError,
			Message:          "Internal server error on creating business universe",
			BusinessUniverse: &BusinessUniverse{},
		}, nil
	} else if err != nil && message != "" {
		return &BusinessUniverseCreateResponse{
			Code:             globalvars.NotFound,
			Message:          message,
			BusinessUniverse: &BusinessUniverse{},
		}, nil
	} else {
		return &BusinessUniverseCreateResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			BusinessUniverse: &BusinessUniverse{
				Id:     item.ID.String(),
				Name:   item.Name,
				Domain: item.Domain,
			},
		}, nil
	}
}

func (s *GrpcServer) UpdateBusinessUniverse(ctx context.Context, in *BusinessUniverseUpdateRequest) (*BusinessUniverseUpdateResponse, error) {
	tmpl := toItemUpdateModel(in)
	id, _ := uuid.Parse(in.Id)
	item, message, err := s.BusinessUniverseRepo.UpdateItem(context.TODO(), tmpl, id)

	if err != nil && message == "" {
		return &BusinessUniverseUpdateResponse{
			Code:             globalvars.ServerInternalError,
			Message:          "Internal server error on updating business universe",
			BusinessUniverse: &BusinessUniverse{},
		}, nil
	} else if err != nil && message != "" {
		return &BusinessUniverseUpdateResponse{
			Code:             globalvars.NotFound,
			Message:          message,
			BusinessUniverse: &BusinessUniverse{},
		}, nil
	} else {
		return &BusinessUniverseUpdateResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			BusinessUniverse: &BusinessUniverse{
				Id:     item.ID.String(),
				Name:   item.Name,
				Domain: item.Domain,
			},
		}, nil
	}
}

func (s *GrpcServer) GetBusinessUniverse(ctx context.Context, in *BusinessUniverseGetRequest) (*BusinessUniverseGetResponse, error) {

	id, _ := uuid.Parse(in.Id)

	bu := s.BusinessUniverseRepo.FindBusinessUniverseById(ctx, id)

	return &BusinessUniverseGetResponse{
		BusinessUniverse: &BusinessUniverse{
			Id:     in.Id,
			Name:   bu.Name,
			Domain: bu.Domain,
		},
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}
