package organization

import (
	"auth/auth_back/pkg/globalvars"
	context "context"

	uuid "github.com/google/uuid"
)

/**
* Company bank
**/
func (s *GrpcServer) AddCompanyBank(ctx context.Context, in *AddCompanyBankRequest) (*CompanyResponse, error) {
	tmpl := toBankCreateModel(in)

	_, err, message := s.CompanyRepo.CreateBankItem(context.TODO(), tmpl)

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

func (s *GrpcServer) UpdateCompanyBank(ctx context.Context, in *UpdateBankRequest) (*CompanyResponse, error) {
	tmpl := toBankUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, err, message := s.CompanyRepo.UpdateBankItem(context.TODO(), tmpl, id)

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

func (s *GrpcServer) GetCompanyBank(ctx context.Context, in *GetBankRequest) (*GetBankResponse, error) {

	id, _ := uuid.Parse(in.Id)

	item := s.CompanyRepo.FindBankItemById(ctx, id)

	return &GetBankResponse{
		Bank:    toBankGet(item),
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}

func (s *GrpcServer) GetCompanyBanks(ctx context.Context, in *GetCompanyBanksRequest) (*GetBanksResponse, error) {

	companyId, _ := uuid.Parse(in.CompanyId)

	items := s.CompanyRepo.FindAllBankItems(ctx, companyId)

	return &GetBanksResponse{
		Banks:   toBanksGet(items),
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}

/**
* Customer bank
**/
func (s *GrpcServer) AddCustomerBank(ctx context.Context, in *AddCustomerBankRequest) (*CustomerResponse, error) {
	tmpl := toCustomerBankCreateModel(in)

	_, err, message := s.CustomerRepo.CreateBankItem(context.TODO(), tmpl)

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

func (s *GrpcServer) UpdateCustomerBank(ctx context.Context, in *UpdateBankRequest) (*CustomerResponse, error) {
	tmpl := toCustomerBankUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, err, message := s.CustomerRepo.UpdateBankItem(context.TODO(), tmpl, id)

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

func (s *GrpcServer) GetCustomerBank(ctx context.Context, in *GetBankRequest) (*GetBankResponse, error) {

	id, _ := uuid.Parse(in.Id)

	item := s.CustomerRepo.FindBankItemById(ctx, id)

	return &GetBankResponse{
		Bank:    toCustomerBankGet(item),
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}

func (s *GrpcServer) GetCustomerBanks(ctx context.Context, in *GetCustomerBanksRequest) (*GetBanksResponse, error) {

	customerId, _ := uuid.Parse(in.CustomerId)

	items := s.CustomerRepo.FindAllBankItems(ctx, customerId)

	return &GetBanksResponse{
		Banks:   toCustomerBanksGet(items),
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}
