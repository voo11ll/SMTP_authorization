package user

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/passwordHelper"
	"context"
	"crypto/sha1"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func (s *GrpcServer) GetUserByEmail(ctx context.Context, in *UserGetByEmailRequest) (*UserResponse, error) {

	email := in.Email

	userFounded := s.UserRepo.FindItemByEmail(context.TODO(), email)

	if userFounded == nil {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: "User not found",
			User:    nil,
		}, nil
	}

	role := s.RoleRepo.FindItemById(ctx, userFounded.RoleID)

	return &UserResponse{
		Code:    globalvars.StatusOK,
		Message: "User get success",
		User:    toUserResponse(userFounded, role),
	}, nil
}

func (s *GrpcServer) GetUserById(ctx context.Context, in *UserGetByIdRequest) (*UserResponse, error) {

	id, _ := uuid.Parse(in.Id)

	userFounded := s.UserRepo.FindItemById(context.TODO(), id)
	role := s.RoleRepo.FindItemById(ctx, userFounded.RoleID)

	return &UserResponse{
		Code:    globalvars.StatusOK,
		Message: "User get success",
		User:    toUserResponse(userFounded, role),
	}, nil
}

func (s *GrpcServer) UpdateUser(ctx context.Context, in *UserUpdateRequest) (*UserResponse, error) {

	_user := toUserUpdate(in)
	id, _ := uuid.Parse(in.Id)

	user, message, err := s.UserRepo.UpdateItem(context.TODO(), _user, id)
	if err != nil && message == "" {
		return &UserResponse{
			Code:    globalvars.ServerInternalError,
			Message: "Internal server error on updating user",
			User:    nil,
		}, nil
	} else if err != nil && message != "" {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: message,
			User:    nil,
		}, nil
	} else {
		role := s.RoleRepo.FindItemById(ctx, user.RoleID)
		return &UserResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			User:    toUserResponse(user, role),
		}, nil
	}
}

func (s *GrpcServer) ChangePassword(ctx context.Context, in *UserChangePasswordRequest) (*UserResponse, error) {

	id, _ := uuid.Parse(in.Id)
	_password := in.Password

	pwd := sha1.New()
	pwd.Write([]byte(_password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))

	password, _ := passwordHelper.HashPassword(in.Password)

	user, message, err := s.UserRepo.ChangeUserPassword(context.TODO(), id, password)
	if err != nil && message == "" {
		return &UserResponse{
			Code:    globalvars.ServerInternalError,
			Message: "Internal server error on updating user",
			User:    nil,
		}, err
	} else if err != nil && message != "" {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: message,
			User:    nil,
		}, err
	} else {
		role := s.RoleRepo.FindItemById(ctx, user.RoleID)
		return &UserResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			User:    toUserResponse(user, role),
		}, nil
	}
}
