package customerUserHandler

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

func AddCustomerUser(w http.ResponseWriter, r *http.Request) {
	var addReq AddUserRequest

	err := httpServerHelper.ExtractBody(r.Body, &addReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	clientConn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		l.LogError(err.Error(), "pkg/user/supporthttp/handler.SignIn")
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	defer clientConn.Close()

	authClient := user.NewB24UserServiceClient(clientConn)

	// Get admin role for main customer user
	responseRole, err := authClient.GetRoleByName(context.Background(), &user.RoleGetByNameRequest{
		Name: "admin",
	})

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	// Add user - add main user
	_addReq := toAddUserRequest(&addReq, responseRole.Role.Id)

	response, err := authClient.AddCustomerUser(context.Background(), _addReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	addResp := toAddUserResponse(response)

	if response.Code == globalvars.NotFound {
		w.WriteHeader(int(globalvars.NotFound))
	} else {
		w.WriteHeader(int(globalvars.StatusOK))
	}

	err = json.NewEncoder(w).Encode(addResp)

	if err != nil {
		l.LogError(err.Error(), "pkg/auth/handler.AddUser")
	}
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var changePassReq ChangePasswordRequest

	userId := r.Header.Get("userId")

	err := httpServerHelper.ExtractBody(r.Body, &changePassReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.ChangePassword")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	changePassReq.Id = userId

	_userUpdReq := toChangePasswordRequest(&changePassReq)

	response, _ := c.ChangeCustomerUserPassword(context.Background(), _userUpdReq)

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

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.UpdateUser")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	updateReq.Id = userId

	_userUpdReq := toUserUpdateRequest(&updateReq)

	response, _ := c.UpdateCustomerUser(context.Background(), _userUpdReq)

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

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserByEmail")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_getUserReq := toUserGetByEmailRequest(&getUserReq)

	response, err := c.GetCustomerUserByEmail(context.Background(), _getUserReq)

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

	conn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithInsecure())
	if err != nil {
		l.LogError(err.Error(), "pkg/user/userhttp/handler.GetUserById")
	}
	defer conn.Close()

	c := user.NewB24UserServiceClient(conn)

	_getUserReq := toUserGetByIdRequest(&getUserReq)

	response, err := c.GetCustomerUserById(context.Background(), _getUserReq)

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
