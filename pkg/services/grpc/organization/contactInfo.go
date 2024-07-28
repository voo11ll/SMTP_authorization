package organization

import (
	"auth/auth_back/pkg/globalvars"
	context "context"

	uuid "github.com/google/uuid"
)

/**
* Company contact info
**/

func (s *GrpcServer) AddCompanyContactInfo(ctx context.Context, in *AddCompanyContactInfoRequest) (*CompanyResponse, error) {
	tmpl := toCompanyContactInfoCreateModel(in)

	_, err, message := s.CompanyRepo.CreateContactInfoItem(context.TODO(), tmpl)

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
		company, _ := s.GetCompany(ctx, &GetRequest{Id: in.CompanyId})
		return &CompanyResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			Company: company.Company,
		}, nil
	}
}

func (s *GrpcServer) UpdateCompanyContactInfo(ctx context.Context, in *UpdateContactInfoRequest) (*CompanyResponse, error) {
	tmpl := toCompanyContactInfoUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, err, message := s.CompanyRepo.UpdateContactInfoItem(context.TODO(), tmpl, id)

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
		company, _ := s.GetCompany(ctx, &GetRequest{Id: item.CompanyID.String()})
		return &CompanyResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			Company: company.Company,
		}, nil
	}
}

func (s *GrpcServer) GetCompanyContactInfo(ctx context.Context, in *GetContactInfoRequest) (*GetContactInfoResponse, error) {

	id, _ := uuid.Parse(in.Id)

	item := s.CompanyRepo.FindContactInfoItemById(ctx, id)

	return &GetContactInfoResponse{
		ContactInfo: toCompanyContactInfoGet(item, s),
		Code:        globalvars.StatusOK,
		Message:     "OK",
	}, nil
}

func (s *GrpcServer) GetCompanyContactInfos(ctx context.Context, in *GetCompanyContactInfosRequest) (*GetContactInfosResponse, error) {

	companyId, _ := uuid.Parse(in.CompanyId)

	items := s.CompanyRepo.FindAllContactInfoItems(ctx, companyId)

	return &GetContactInfosResponse{
		ContactInfos: toCompanyContactInfosGet(items, s),
		Code:         globalvars.StatusOK,
		Message:      "OK",
	}, nil
}

func (s *GrpcServer) DeleteCompanyContactInfo(ctx context.Context, in *DeleteContactInfoRequest) (*CompanyResponse, error) {
	id, _ := uuid.Parse(in.Id)
	item, err, message := s.CompanyRepo.DeleteContactInfoItem(context.TODO(), id)

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
		company, _ := s.GetCompany(ctx, &GetRequest{Id: item.CompanyID.String()})
		return &CompanyResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			Company: company.Company,
		}, nil
	}
}

/**
* Customer contact info
**/

func (s *GrpcServer) AddCustomerContactInfo(ctx context.Context, in *AddCustomerContactInfoRequest) (*CustomerResponse, error) {
	tmpl := toCustomerContactInfoCreateModel(in)

	_, err, message := s.CustomerRepo.CreateContactInfoItem(context.TODO(), tmpl)

	if err != nil && message == "" {
		return &CustomerResponse{
			Code:     globalvars.ServerInternalError,
			Message:  "Internal server error on creating bank",
			Customer: nil,
		}, nil
	} else if err != nil && message != "" {
		return &CustomerResponse{
			Code:     globalvars.NotFound,
			Message:  message,
			Customer: nil,
		}, nil
	} else {
		customer, _ := s.GetCustomer(ctx, &GetRequest{Id: in.CustomerId})
		return &CustomerResponse{
			Code:     globalvars.StatusOK,
			Message:  message,
			Customer: customer.Customer,
		}, nil
	}
}

func (s *GrpcServer) UpdateCustomerContactInfo(ctx context.Context, in *UpdateContactInfoRequest) (*CustomerResponse, error) {
	tmpl := toCustomerContactInfoUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, err, message := s.CustomerRepo.UpdateContactInfoItem(context.TODO(), tmpl, id)

	if err != nil && message == "" {
		return &CustomerResponse{
			Code:     globalvars.ServerInternalError,
			Message:  "Internal server error on updating role",
			Customer: nil,
		}, nil
	} else if err != nil && message != "" {
		return &CustomerResponse{
			Code:     globalvars.NotFound,
			Message:  message,
			Customer: nil,
		}, nil
	} else {
		customer, _ := s.GetCustomer(ctx, &GetRequest{Id: item.CustomerID.String()})
		return &CustomerResponse{
			Code:     globalvars.StatusOK,
			Message:  message,
			Customer: customer.Customer,
		}, nil
	}
}

func (s *GrpcServer) GetCustomerContactInfo(ctx context.Context, in *GetContactInfoRequest) (*GetContactInfoResponse, error) {

	id, _ := uuid.Parse(in.Id)

	item := s.CustomerRepo.FindContactInfoItemById(ctx, id)

	return &GetContactInfoResponse{
		ContactInfo: toCustomerContactInfoGet(item, s),
		Code:        globalvars.StatusOK,
		Message:     "OK",
	}, nil
}

func (s *GrpcServer) GetCustomerContactInfos(ctx context.Context, in *GetCustomerContactInfosRequest) (*GetContactInfosResponse, error) {

	customerId, _ := uuid.Parse(in.CustomerId)

	items := s.CustomerRepo.FindAllContactInfoItems(ctx, customerId)

	return &GetContactInfosResponse{
		ContactInfos: toCustomerContactInfosGet(items, s),
		Code:         globalvars.StatusOK,
		Message:      "OK",
	}, nil
}

func (s *GrpcServer) DeleteCustomerContactInfo(ctx context.Context, in *DeleteContactInfoRequest) (*CustomerResponse, error) {
	id, _ := uuid.Parse(in.Id)
	item, err, message := s.CustomerRepo.DeleteContactInfoItem(context.TODO(), id)

	if err != nil && message == "" {
		return &CustomerResponse{
			Code:     globalvars.ServerInternalError,
			Message:  "Internal server error on updating role",
			Customer: nil,
		}, nil
	} else if err != nil && message != "" {
		return &CustomerResponse{
			Code:     globalvars.NotFound,
			Message:  message,
			Customer: nil,
		}, nil
	} else {
		customer, _ := s.GetCustomer(ctx, &GetRequest{Id: item.CustomerID.String()})
		return &CustomerResponse{
			Code:     globalvars.StatusOK,
			Message:  message,
			Customer: customer.Customer,
		}, nil
	}
}
