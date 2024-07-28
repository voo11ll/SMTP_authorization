package organization

import (
	"auth/auth_back/pkg/globalvars"
	context "context"

	uuid "github.com/google/uuid"
)

func (s *GrpcServer) CreateContactType(ctx context.Context, in *ContactTypeCreateRequest) (*ContactTypeResponse, error) {
	tmpl := toContactTypeCreateModel(in)

	item, err, message := s.ContactTypeRepo.CreateItem(context.TODO(), tmpl)

	if err != nil && message == "" {
		return &ContactTypeResponse{
			Code:        globalvars.ServerInternalError,
			Message:     "Internal server error on creating role",
			ContactType: nil,
		}, nil
	} else if err != nil && message != "" {
		return &ContactTypeResponse{
			Code:        globalvars.NotFound,
			Message:     message,
			ContactType: nil,
		}, nil
	} else {
		return &ContactTypeResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			ContactType: &ContactType{
				Id:     item.ID.String(),
				Name:   item.Name,
				Notion: item.Notion,
			},
		}, nil
	}
}

func (s *GrpcServer) UpdateContactType(ctx context.Context, in *ContactTypeUpdateRequest) (*ContactTypeResponse, error) {
	tmpl := toContactTypeUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, err, message := s.ContactTypeRepo.UpdateItem(context.TODO(), tmpl, id)

	if err != nil && message == "" {
		return &ContactTypeResponse{
			Code:        globalvars.ServerInternalError,
			Message:     "Internal server error on updating role",
			ContactType: nil,
		}, nil
	} else if err != nil && message != "" {
		return &ContactTypeResponse{
			Code:        globalvars.NotFound,
			Message:     message,
			ContactType: nil,
		}, nil
	} else {
		return &ContactTypeResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			ContactType: &ContactType{
				Id:     item.ID.String(),
				Name:   item.Name,
				Notion: item.Notion,
			},
		}, nil
	}
}

func (s *GrpcServer) GetContactType(ctx context.Context, in *ContactTypeGetRequest) (*ContactTypeResponse, error) {

	id, _ := uuid.Parse(in.Id)

	item := s.ContactTypeRepo.FindItemById(ctx, id)

	return &ContactTypeResponse{
		ContactType: &ContactType{
			Id:     in.Id,
			Name:   item.Name,
			Notion: item.Notion,
		},
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}

func (s *GrpcServer) GetContactTypes(ctx context.Context, in *ContactTypesGetRequest) (*ContactTypesResponse, error) {

	items := s.ContactTypeRepo.FindAllItems(ctx)

	return &ContactTypesResponse{
		ContactTypes: toContactTypesGet(items),
		Code:         globalvars.StatusOK,
		Message:      "OK",
	}, nil
}
