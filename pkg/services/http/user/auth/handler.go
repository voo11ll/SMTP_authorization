package authHandler

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/helpers/tokenHelper"
	"auth/auth_back/pkg/logger"
	"auth/auth_back/pkg/services/grpc/notification"
	"auth/auth_back/pkg/services/grpc/user"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var l = logger.Logger{}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var signInReq SignInRequest

	err := httpServerHelper.ExtractBody(r.Body, &signInReq)

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

	c := user.NewB24UserServiceClient(clientConn)

	_signInReq := toUserSignInRequest(&signInReq)

	response, err := c.SignIn(context.Background(), _signInReq)
	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	if response.User == nil {
		if response.Code == globalvars.NotFound {
			w.WriteHeader(int(globalvars.NotFound))
			err = json.NewEncoder(w).Encode(&SignInResponse{
				Token:   "",
				Message: response.Message,
				User:    UserType{},
			})
			if err != nil {
				httpServerHelper.ReturnErr(w, err, err.Error())
				return
			}
		} else if response.Code == globalvars.Unauthorized {
			w.WriteHeader(int(globalvars.Unauthorized))
			err = json.NewEncoder(w).Encode(&SignInResponse{
				Token:   "",
				Message: response.Message,
				User:    UserType{},
			})
			if err != nil {
				httpServerHelper.ReturnErr(w, err, err.Error())
				return
			}
		}
	} else {
		userId, _ := uuid.Parse(response.User.Id)

		tokens, _ := tokenHelper.CreateUserTokens(userId)

		signInResp := toSignInResponse(response, tokens.AccessToken)

		if response.Code == globalvars.NotFound {
			w.WriteHeader(int(globalvars.NotFound))
		} else {
			w.WriteHeader(int(globalvars.StatusOK))
		}

		err = json.NewEncoder(w).Encode(signInResp)

		if err != nil {
			l.LogError(err.Error(), "pkg/auth/handler.SignIn")
		}
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var signUpReq SignUpRequest

	err := httpServerHelper.ExtractBody(r.Body, &signUpReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	clientConn, err := grpc.Dial(viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		l.LogError(err.Error(), "pkg/user/supporthttp/handler.SignUp")
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	defer clientConn.Close()

	authClient := user.NewB24UserServiceClient(clientConn)

	// Sing up user - register main user
	_signUpReq := toUserSignUpRequest(&signUpReq)

	response, err := authClient.SignUp(context.Background(), _signUpReq)

	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	signUpResp := toSignUpResponse(response)

	if response.Code == globalvars.NotFound {
		w.WriteHeader(int(globalvars.NotFound))
	} else {
		w.WriteHeader(int(globalvars.StatusOK))
	}

	err = json.NewEncoder(w).Encode(signUpResp)

	if err != nil {
		l.LogError(err.Error(), "pkg/auth/handler.SignUp")
	}
}

func EmailConfirm(w http.ResponseWriter, r *http.Request) {
	linkId := r.URL.Query().Get("id")
	keyHash := r.URL.Query().Get("key")

	clientConn, err := grpc.Dial(viper.GetString("grpc.notification.host")+":"+viper.GetString("grpc.notification.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.LogError(err.Error(), "pkg/user/supporthttp/handler.EmailConfirm")
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}
	defer clientConn.Close()

	notificationServiceClient := notification.NewNotificationServiceClient(clientConn)

	response, err := notificationServiceClient.MailConfirmation(context.TODO(), &notification.MailConfirmationRequest{
		LinkId:  linkId,
		HashKey: keyHash,
	})
	if err != nil {
		httpServerHelper.ReturnErr(w, err, err.Error())
		return
	}

	emailConfirmResp := &EmailConfirmResponse{
		Status:  response.StatusENUM,
		Message: response.Message,
	}

	err = json.NewEncoder(w).Encode(emailConfirmResp)
	if err != nil {
		l.LogError(err.Error(), "pkg/auth/handler.EmailConfirm")
	}
}
