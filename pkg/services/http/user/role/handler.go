package roleHandler

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/logger"

	"auth/auth_back/pkg/services/grpc/user"
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var l = logger.Logger{}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	var createReq CreateRequest

	err := httpServerHelper.ExtractBody(r.Body, &createReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.CreateRole")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_createReq := toCreateRequest(&createReq)

	response, err := c.CreateRole(context.Background(), _createReq)

	createRes := toRoleResponse(response)

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

func GetRole(w http.ResponseWriter, r *http.Request) {
	var getReq GetRequest

	err := httpServerHelper.ExtractBody(r.Body, &getReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRole")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_getReq := toGetRequest(&getReq)

	response, err := c.GetRole(context.Background(), _getReq)

	getResp := toRoleResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &RoleResponse{
			Message: globalvars.ErrTokenAccess.Error.Error(),
			Code:    globalvars.Unauthorized,
			Role:    RoleType{},
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

func GetRoleByName(w http.ResponseWriter, r *http.Request) {
	var getReq GetByNameRequest

	err := httpServerHelper.ExtractBody(r.Body, &getReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRole")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_getReq := toGetByNameRequest(&getReq)

	response, err := c.GetRoleByName(context.Background(), _getReq)

	getResp := toRoleResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &RoleResponse{
			Message: globalvars.ErrTokenAccess.Error.Error(),
			Code:    globalvars.Unauthorized,
			Role:    RoleType{},
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

func GetRoles(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.GetRoles")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	response, err := c.GetRoles(context.Background(), &user.RolesGetRequest{})

	getItemsResp := toGetsResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &GetsResponse{
			Message: globalvars.ErrTokenAccess.Error.Error(),
			Code:    response.Code,
			Roles:   []RoleType{},
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

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	var updateReq UpdateRequest

	err := httpServerHelper.ExtractBody(r.Body, &updateReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/role/roleHttp/handler.UpdateRole")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_userUpdReq := toUpdateRequest(&updateReq)

	response, _ := c.UpdateRole(context.Background(), _userUpdReq)

	updateRes := toRoleResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := RoleResponse{
			Message: globalvars.ErrServerInternal.Error.Error(),
			Code:    globalvars.ServerInternalError,
			Role:    RoleType{},
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.ServerInternalError))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/role/roleHttp/handler.UpdateRole")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := RoleResponse{
			Message: globalvars.GetUpdateErrors("Role").Error.Error(),
			Code:    globalvars.NotFound,
			Role:    RoleType{},
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.NotFound))
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
