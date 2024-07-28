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

func CreateContactInfo(w http.ResponseWriter, r *http.Request) {

	var createReq CreateContactInfoRequest

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

	_createReq := toGrpcCreateCompanyContactInfoRequest(&createReq)

	c := org.NewB24OrganizationServiceClient(conn)

	response, err := c.AddCompanyContactInfo(context.Background(), &org.AddCompanyContactInfoRequest{
		CompanyId:   createReq.CompanyId,
		ContactInfo: _createReq,
	})

	createRes := companyHandler.ToCompanyResponse(response)

	if response.Code == globalvars.NotFound {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
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

func GetContactInfo(w http.ResponseWriter, r *http.Request) {
	var getReq GetContactInfoRequest

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

	_getReq := toGetContactInfoRequest(&getReq)

	response, err := c.GetCompanyContactInfo(context.Background(), _getReq)

	getResp := companyHandler.ToGetContactInfoResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &companyHandler.GetContactInfoResponse{
			Message:     globalvars.ErrTokenAccess.Error.Error(),
			Code:        globalvars.Unauthorized,
			ContactInfo: companyHandler.ContactInfoType{},
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

func GetContactInfos(w http.ResponseWriter, r *http.Request) {
	var getReq GetContactInfosRequest

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

	_getReq := toGetContactInfosRequest(&getReq)

	response, err := c.GetCompanyContactInfos(context.Background(), _getReq)

	getResp := companyHandler.ToGetContactInfosResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &companyHandler.GetContactInfosResponse{
			Message:      globalvars.ErrTokenAccess.Error.Error(),
			Code:         globalvars.Unauthorized,
			ContactInfos: nil,
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

func UpdateContactInfo(w http.ResponseWriter, r *http.Request) {
	var updateReq UpdateContactInfoRequest

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

	response, _ := c.UpdateCompanyContactInfo(context.Background(), &org.UpdateContactInfoRequest{
		Id:            updateReq.Id,
		ContactTypeID: updateReq.ContactTypeId,
		Value:         updateReq.Value,
	})

	updateRes := companyHandler.ToCompanyResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := companyHandler.CompanyResponse{
			Message: globalvars.ErrServerInternal.Error.Error(),
			Code:    globalvars.ServerInternalError,
			Company: companyHandler.CompanyType{},
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
			Company: companyHandler.CompanyType{},
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
