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

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var createReq CreateCompanyRequest

	err := httpServerHelper.ExtractBody(r.Body, &createReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/company/companyhttp/handler.CreateCompany")
	}
	defer conn.Close()

	_createReq := toGrpcCreateCompanyRequest(&createReq)

	c := org.NewB24OrganizationServiceClient(conn)

	response, err := c.CreateCompany(context.Background(), _createReq)

	createRes := companyHandler.ToCompanyResponse(response)

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

func GetCompany(w http.ResponseWriter, r *http.Request) {
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

	response, err := c.GetCompany(context.Background(), _getReq)

	getResp := companyHandler.ToCompanyResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &companyHandler.CompanyResponse{
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

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	var getReq GetCompaniesRequest

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

	_getReq := toGetCompaniesRequest(&getReq)

	response, err := c.GetCompanies(context.Background(), _getReq)

	getResp := companyHandler.ToCompaniesResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &companyHandler.CompaniesResponse{
			Message:   globalvars.ErrTokenAccess.Error.Error(),
			Code:      globalvars.Unauthorized,
			Companies: nil,
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

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
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

	response, _ := c.UpdateCompany(context.Background(), &org.UpdateRequest{
		Id:           updateReq.Id,
		Name:         updateReq.Name,
		FullName:     updateReq.FullName,
		Inn:          updateReq.INN,
		Kpp:          updateReq.KPP,
		LegalAddress: updateReq.LegalAddress,
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
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.UpdateBusinessUniverse")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := companyHandler.CompanyResponse{
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
