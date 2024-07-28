package organization

import (
	"auth/auth_back/pkg/globalvars"
	context "context"

	uuid "github.com/google/uuid"
)

func (s *GrpcServer) CreateCompany(ctx context.Context, in *CreateRequest) (*CompanyResponse, error) {
	tmpl := toCompanyCreateModel(in)

	item, err, message := s.CompanyRepo.CreateItem(context.TODO(), tmpl)

	if err != nil && message == "" {
		return &CompanyResponse{
			Code:    globalvars.ServerInternalError,
			Message: "Internal server error on creating bank",
			Company: nil,
		}, nil
	} else if err != nil && message != "" {
		return &CompanyResponse{
			Code:    globalvars.NotFound,
			Message: message,
			Company: nil,
		}, nil
	} else {
		return &CompanyResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			Company: toCompanyGet(item, s),
		}, nil
	}
}

func (s *GrpcServer) UpdateCompany(ctx context.Context, in *UpdateRequest) (*CompanyResponse, error) {
	tmpl := toCompanyUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, err, message := s.CompanyRepo.UpdateItem(context.TODO(), tmpl, id)

	if err != nil && message == "" {
		return &CompanyResponse{
			Code:    globalvars.ServerInternalError,
			Message: "Internal server error on updating role",
			Company: nil,
		}, nil
	} else if err != nil && message != "" {
		return &CompanyResponse{
			Code:    globalvars.NotFound,
			Message: message,
			Company: nil,
		}, nil
	} else {
		return &CompanyResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			Company: toCompanyGet(item, s),
		}, nil
	}
}

func (s *GrpcServer) GetCompany(ctx context.Context, in *GetRequest) (*CompanyResponse, error) {

	id, _ := uuid.Parse(in.Id)

	item := s.CompanyRepo.FindItemById(ctx, id)

	return &CompanyResponse{
		Company: toCompanyGet(item, s),
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}

func (s *GrpcServer) GetCompanies(ctx context.Context, in *GetsRequest) (*CompaniesGetResponse, error) {

	businessUniverseId, _ := uuid.Parse(in.BusinessUniverseId)

	items := s.CompanyRepo.FindAllItems(ctx, businessUniverseId)

	return &CompaniesGetResponse{
		Companies: toCompaniesGet(items, s),
		Code:      globalvars.StatusOK,
		Message:   "OK",
	}, nil
}
