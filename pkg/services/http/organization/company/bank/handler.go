package companyHandler

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/logger"
	org "auth/auth_back/pkg/services/grpc/organization"
	companyHandler "auth/auth_back/pkg/services/http/organization/company"
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var l = logger.Logger{}

func CreateBank(w http.ResponseWriter, r *http.Request) {

	var createReq CreateBankRequest

	err := httpServerHelper.ExtractBody(r.Body, &createReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.CreateBank")
	}
	defer conn.Close()

	_createReq := toGrpcCreateCompanyBankRequest(&createReq)

	c := org.NewB24OrganizationServiceClient(conn)

	response, err := c.AddCompanyBank(context.Background(), &org.AddCompanyBankRequest{
		CompanyId: createReq.CompanyId,
		Bank:      _createReq,
	})

	createRes := companyHandler.ToCompanyResponse(response)

	if response.Code == globalvars.NotFound {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(createRes); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.CreateBank")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(createRes); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.CreateBank")
		}
	}
}

func GetBank(w http.ResponseWriter, r *http.Request) {
	var getReq GetRequest

	err := httpServerHelper.ExtractBody(r.Body, &getReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.GetBank")
	}
	defer conn.Close()

	c := org.NewB24OrganizationServiceClient(conn)

	_getReq := toGetRequest(&getReq)

	response, err := c.GetCompanyBank(context.Background(), _getReq)

	getResp := companyHandler.ToGetBankResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &companyHandler.GetBankResponse{
			Message: globalvars.ErrTokenAccess.Error.Error(),
			Code:    globalvars.Unauthorized,
			Bank:    companyHandler.BankType{},
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.Unauthorized))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.GetBank")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getResp); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.GetBank")
		}
	}
}

func GetBanks(w http.ResponseWriter, r *http.Request) {
	var getReq GetBanksRequest

	err := httpServerHelper.ExtractBody(r.Body, &getReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.GetBank")
	}
	defer conn.Close()

	c := org.NewB24OrganizationServiceClient(conn)

	_getReq := toGetBanksRequest(&getReq)

	response, err := c.GetCompanyBanks(context.Background(), _getReq)

	getResp := companyHandler.ToGetBanksResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &companyHandler.GetBanksResponse{
			Message: globalvars.ErrTokenAccess.Error.Error(),
			Code:    globalvars.Unauthorized,
			Banks:   nil,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.Unauthorized))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.GetBank")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getResp); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.GetBank")
		}
	}
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	var updateReq UpdateRequest

	err := httpServerHelper.ExtractBody(r.Body, &updateReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.UpdateBank")
	}
	defer conn.Close()

	c := org.NewB24OrganizationServiceClient(conn)

	response, _ := c.UpdateCompanyBank(context.Background(), &org.UpdateBankRequest{
		Id:                updateReq.Id,
		Name:              updateReq.Name,
		AccountNumber:     updateReq.AccountNumber,
		Bik:               updateReq.Bik,
		CorrAccountNumber: updateReq.CorrAccountNumber,
	})

	updateRes := companyHandler.ToCompanyResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := companyHandler.CompanyResponse{
			Message: globalvars.ErrServerInternal.Error.Error(),
			Code:    globalvars.ServerInternalError,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.ServerInternalError))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.UpdateBank")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := companyHandler.CompanyResponse{
			Message: response.Message,
			Code:    globalvars.Error,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.UpdateBank")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(updateRes); err != nil {
			l.LogError(err.Error(), "pkg/Bank/Bankhttp/handler.UpdateBank")
		}
	}
}
