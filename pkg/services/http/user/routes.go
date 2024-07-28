package userHttp

import (
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/middleware"
	authHandler "auth/auth_back/pkg/services/http/user/auth"
	roleHandler "auth/auth_back/pkg/services/http/user/role"
	userHandler "auth/auth_back/pkg/services/http/user/user"
)

var Routes = httpServerHelper.Routes{
	/**
	* Company auth users
	**/
	httpServerHelper.Route{
		Name:        "SignIn",
		Method:      "POST",
		Pattern:     "/user/sign-in",
		HandlerFunc: authHandler.SignIn,
	},
	httpServerHelper.Route{
		Name:        "SignUp",
		Method:      "POST",
		Pattern:     "/user/sign-up",
		HandlerFunc: authHandler.SignUp,
	},
	httpServerHelper.Route{
		Name:        "EmailConfirm",
		Method:      "GET",
		Pattern:     "/user/email-confirm",
		HandlerFunc: authHandler.EmailConfirm,
	},
	/**
	* Roles
	**/
	httpServerHelper.Route{
		Name:        "CreateRole",
		Method:      "POST",
		Pattern:     "/user/role",
		HandlerFunc: roleHandler.CreateRole,
	},
	httpServerHelper.Route{
		Name:        "GetRole",
		Method:      "GET",
		Pattern:     "/user/role",
		HandlerFunc: roleHandler.GetRole,
	},
	httpServerHelper.Route{
		Name:        "GetRoleByName",
		Method:      "GET",
		Pattern:     "/user/role/by-name",
		HandlerFunc: roleHandler.GetRoleByName,
	},
	httpServerHelper.Route{
		Name:        "GetRoles",
		Method:      "GET",
		Pattern:     "/user/roles/list",
		HandlerFunc: roleHandler.GetRoles,
	},
	httpServerHelper.Route{
		Name:        "UpdateRole",
		Method:      "PUT",
		Pattern:     "/user/role",
		HandlerFunc: roleHandler.UpdateRole,
	},
	/**
	* Company users
	**/
	httpServerHelper.Route{
		Name:        "ChangePassword",
		Method:      "PUT",
		Pattern:     "/user/password",
		HandlerFunc: userHandler.ChangePassword,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "UpdateUser",
		Method:      "PUT",
		Pattern:     "/user",
		HandlerFunc: userHandler.UpdateUser,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetUser",
		Method:      "GET",
		Pattern:     "/user",
		HandlerFunc: userHandler.GetUserById,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetUserByEmail",
		Method:      "GET",
		Pattern:     "/user/by-email",
		HandlerFunc: userHandler.GetUserByEmail,
		Middleware:  middleware.UserMiddleware,
	},
}
