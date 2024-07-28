package businessUniverseHandler

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/logger"
	org "auth/auth_back/pkg/services/grpc/organization"
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var l = logger.Logger{}

func CreateBusinessUniverse(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.CreateBusinessUniverse")
	}
	defer conn.Close()

	c := org.NewB24OrganizationServiceClient(conn)

	response, err := c.CreateBusinessUniverse(context.Background(), &org.BusinessUniverseCreateRequest{})

	createRes := toCreateResponse(response)

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

func GetBusinessUniverse(w http.ResponseWriter, r *http.Request) {
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

	response, err := c.GetBusinessUniverse(context.Background(), _getReq)

	getResp := toGetResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &GetResponse{
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

func UpdateBusinessUniverse(w http.ResponseWriter, r *http.Request) {
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

	response, _ := c.UpdateBusinessUniverse(context.Background(), &org.BusinessUniverseUpdateRequest{
		Id:     updateReq.Id,
		Name:   updateReq.Name,
		Domain: updateReq.Name + viper.GetString("grpc.domain.templateUrl"),
	})

	updateRes := toUpdateResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := UpdateResponse{
			Message: globalvars.ErrServerInternal.Error.Error(),
			Code:    globalvars.ServerInternalError,
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.ServerInternalError))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/businessuniverse/businessuniversehttp/handler.UpdateBusinessUniverse")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := UpdateResponse{
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
