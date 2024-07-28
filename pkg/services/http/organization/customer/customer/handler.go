package customerHandler

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/logger"
	org "auth/auth_back/pkg/services/grpc/organization"
	customerHandler "auth/auth_back/pkg/services/http/organization/customer"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var l = logger.Logger{}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var createReq CreateCustomerRequest

	err := httpServerHelper.ExtractBody(r.Body, &createReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/Customer/Customerhttp/handler.CreateCustomer")
	}
	defer conn.Close()

	_createReq := toGrpcCreateCustomerRequest(&createReq)

	c := org.NewB24OrganizationServiceClient(conn)

	response, err := c.CreateCustomer(context.Background(), _createReq)

	if err != nil {
		l.LogError(err.Error(), "pkg/Customer/Customerhttp/handler.CreateCustomer")
	}

	createRes := customerHandler.ToCustomerResponse(response)

	if response.Code == globalvars.NotFound {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.NotFound))
		if err := json.NewEncoder(w).Encode(createRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.CreateBusinessUniverse")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(createRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.CreateBusinessUniverse")
		}
	}
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	var getReq GetRequest

	err := httpServerHelper.ExtractBody(r.Body, &getReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.GetBusinessUniverse")
	}
	defer conn.Close()

	c := org.NewB24OrganizationServiceClient(conn)

	_getReq := toGetRequest(&getReq)

	response, err := c.GetCustomer(context.Background(), _getReq)

	fmt.Println("Customer", response.String())

	getResp := customerHandler.ToCustomerResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &customerHandler.CustomerResponse{
			Message: globalvars.ErrTokenAccess.Error.Error(),
			Code:    globalvars.Unauthorized,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.Unauthorized))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.GetBusinessUniverse")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getResp); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.GetBusinessUniverse")
		}
	}
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var getReq GetCustomersRequest

	err := httpServerHelper.ExtractBody(r.Body, &getReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.GetBusinessUniverse")
	}
	defer conn.Close()

	c := org.NewB24OrganizationServiceClient(conn)

	_getReq := toGetCustomersRequest(&getReq)

	response, err := c.GetCustomers(context.Background(), _getReq)

	getResp := customerHandler.ToCustomersResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &customerHandler.CustomersResponse{
			Message:   globalvars.ErrTokenAccess.Error.Error(),
			Code:      globalvars.Unauthorized,
			Customers: nil,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.Unauthorized))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.GetBusinessUniverse")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getResp); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.GetBusinessUniverse")
		}
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var updateReq UpdateRequest

	err := httpServerHelper.ExtractBody(r.Body, &updateReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.UpdateBusinessUniverse")
	}
	defer conn.Close()

	c := org.NewB24OrganizationServiceClient(conn)

	response, _ := c.UpdateCustomer(context.Background(), &org.UpdateRequest{
		Id:           updateReq.Id,
		Name:         updateReq.Name,
		FullName:     updateReq.FullName,
		Inn:          updateReq.INN,
		Kpp:          updateReq.KPP,
		LegalAddress: updateReq.LegalAddress,
	})

	updateRes := customerHandler.ToCustomerResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := customerHandler.CustomerResponse{
			Message: globalvars.ErrServerInternal.Error.Error(),
			Code:    globalvars.ServerInternalError,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.ServerInternalError))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.UpdateBusinessUniverse")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := customerHandler.CustomerResponse{
			Message: response.Message,
			Code:    globalvars.Error,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.UpdateBusinessUniverse")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(updateRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.UpdateBusinessUniverse")
		}
	}
}
