package contactTypeHandler

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/logger"

	"auth/auth_back/pkg/services/grpc/organization"
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var l = logger.Logger{}

func CreateContactType(w http.ResponseWriter, r *http.Request) {
	var createReq CreateRequest

	err := httpServerHelper.ExtractBody(r.Body, &createReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/organization/contactTypeHttp/handler.CreateContactType")
	}
	defer conn.Close()

	c := organization.NewB24OrganizationServiceClient(conn)

	_createReq := toCreateRequest(&createReq)

	response, err := c.CreateContactType(context.Background(), _createReq)

	createRes := toContactTypeResponse(response)

	if response.Code == globalvars.NotFound {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(createRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.CreateRole")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(createRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.CreateRole")
		}
	}
}

func GetContactType(w http.ResponseWriter, r *http.Request) {
	var getReq GetRequest

	err := httpServerHelper.ExtractBody(r.Body, &getReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRole")
	}
	defer conn.Close()

	c := organization.NewB24OrganizationServiceClient(conn)

	_getReq := toGetRequest(&getReq)

	response, err := c.GetContactType(context.Background(), _getReq)

	getResp := toContactTypeResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &ContactTypeResponse{
			Message:     globalvars.ErrTokenAccess.Error.Error(),
			Code:        globalvars.Unauthorized,
			ContactType: ContactType{},
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.Unauthorized))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRole")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getResp); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRole")
		}
	}
}

func GetContactTypes(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRoles")
	}
	defer conn.Close()

	c := organization.NewB24OrganizationServiceClient(conn)

	response, err := c.GetContactTypes(context.Background(), &organization.ContactTypesGetRequest{})

	getItemsResp := toGetsResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &GetsResponse{
			Message:      globalvars.ErrTokenAccess.Error.Error(),
			Code:         response.Code,
			ContactTypes: []ContactType{},
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRoles")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getItemsResp); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetItems")
		}
	}
}

func UpdateContactType(w http.ResponseWriter, r *http.Request) {
	var updateReq UpdateRequest

	err := httpServerHelper.ExtractBody(r.Body, &updateReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.UpdateRole")
	}
	defer conn.Close()

	c := organization.NewB24OrganizationServiceClient(conn)

	_userUpdReq := toUpdateRequest(&updateReq)

	response, _ := c.UpdateContactType(context.Background(), _userUpdReq)

	updateRes := toContactTypeResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := ContactTypeResponse{
			Message:     globalvars.ErrServerInternal.Error.Error(),
			Code:        globalvars.ServerInternalError,
			ContactType: ContactType{},
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.ServerInternalError))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.UpdateRole")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := ContactTypeResponse{
			Message:     globalvars.GetUpdateErrors("Role").Error.Error(),
			Code:        globalvars.StatusOK,
			ContactType: ContactType{},
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.UpdateRole")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(updateRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.UpdateRole")
		}
	}
}
