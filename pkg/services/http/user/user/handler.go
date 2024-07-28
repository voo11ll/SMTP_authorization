package userHandler

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
	"google.golang.org/grpc/credentials/insecure"
)

var l = logger.Logger{}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var changePassReq ChangePasswordRequest

	userId := r.Header.Get("userId")

	err := httpServerHelper.ExtractBody(r.Body, &changePassReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.ChangePassword")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	changePassReq.Id = userId

	_userUpdReq := toChangePasswordRequest(&changePassReq)

	response, _ := c.ChangePassword(context.Background(), _userUpdReq)

	updateRes := toUpdateResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := UpdateResponse{
			Message: globalvars.ErrServerInternal.Error.Error(),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.ServerInternalError))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.ChangePassword")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := UpdateResponse{
			Message: globalvars.GetExistErrors("User").Error.Error(),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.ChangePassword")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(updateRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.ChangePassword")
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateReq UpdateRequest

	userId := r.Header.Get("userId")

	err := httpServerHelper.ExtractBody(r.Body, &updateReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.UpdateUser")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	updateReq.Id = userId

	_userUpdReq := toUserUpdateRequest(&updateReq)

	response, _ := c.UpdateUser(context.Background(), _userUpdReq)

	updateRes := toUpdateResponse(response)

	if response.Code == globalvars.ServerInternalError {
		_errRes := UpdateResponse{
			Message: globalvars.ErrServerInternal.Error.Error(),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.ServerInternalError))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.UpdateUser")
		}
	} else if response.Code == globalvars.NotFound {
		_errRes := UpdateResponse{
			Message: globalvars.GetExistErrors("User").Error.Error(),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_errRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.UpdateUser")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(updateRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.UpdateUser")
		}
	}
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var getUserReq GetUserByEmailRequest

	err := httpServerHelper.ExtractBody(r.Body, &getUserReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserByEmail")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_getUserReq := toUserGetByEmailRequest(&getUserReq)

	response, err := c.GetUserByEmail(context.Background(), _getUserReq)
	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	getUserResp := toUserGetResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &GetUserResponse{
			Message: globalvars.ErrTokenAccess.Error.Error(),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserByEmail")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getUserResp); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserByEmail")
		}
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	var getUserReq GetUserByIdRequest

	err := httpServerHelper.ExtractBody(r.Body, &getUserReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserById")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_getUserReq := toUserGetByIdRequest(&getUserReq)

	response, err := c.GetUserById(context.Background(), _getUserReq)
	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	getUserResp := toUserGetResponse(response)

	if response.Code == globalvars.Unauthorized {
		_getRes := &GetUserResponse{
			Code:    response.Code,
			Message: globalvars.ErrTokenAccess.Error.Error(),
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(_getRes); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserById")
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(int(globalvars.StatusOK))
		if err := json.NewEncoder(w).Encode(getUserResp); err != nil {
			l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserById")
		}
	}
}
