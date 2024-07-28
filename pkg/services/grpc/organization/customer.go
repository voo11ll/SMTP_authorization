package organization

import (
	"auth/auth_back/pkg/globalvars"
	context "context"

	uuid "github.com/google/uuid"
)

func (s *GrpcServer) CreateCustomer(ctx context.Context, in *CreateRequest) (*CustomerResponse, error) {
	tmpl := toCustomerCreateModel(in)

	item, err, message := s.CustomerRepo.CreateItem(context.TODO(), tmpl)

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
		return &CustomerResponse{
			Code:     globalvars.StatusOK,
			Message:  message,
			Customer: toCustomerGet(item, s),
		}, nil
	}
}

func (s *GrpcServer) UpdateCustomer(ctx context.Context, in *UpdateRequest) (*CustomerResponse, error) {
	tmpl := toCustomerUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, err, message := s.CustomerRepo.UpdateItem(context.TODO(), tmpl, id)

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
		return &CustomerResponse{
			Code:     globalvars.StatusOK,
			Message:  message,
			Customer: toCustomerGet(item, s),
		}, nil
	}
}

func (s *GrpcServer) GetCustomer(ctx context.Context, in *GetRequest) (*CustomerResponse, error) {

	id, _ := uuid.Parse(in.Id)

	item := s.CustomerRepo.FindItemById(ctx, id)

	return &CustomerResponse{
		Customer: toCustomerGet(item, s),
		Code:     globalvars.StatusOK,
		Message:  "OK",
	}, nil
}

func (s *GrpcServer) GetCustomers(ctx context.Context, in *GetsRequest) (*CustomersGetResponse, error) {

	businessUniverseId, _ := uuid.Parse(in.BusinessUniverseId)

	items := s.CustomerRepo.FindAllItems(ctx, businessUniverseId)

	return &CustomersGetResponse{
		Customers: toCustomersGet(items, s),
		Code:      globalvars.StatusOK,
		Message:   "OK",
	}, nil
}
