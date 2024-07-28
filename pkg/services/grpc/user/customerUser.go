package user

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/passwordHelper"
	"context"
	"crypto/sha1"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func (s *GrpcServer) AddCustomerUser(ctx context.Context, in *CustomerUserAddRequest) (*CustomerUserResponse, error) {

	existUserEmail := s.CustomerUserRepo.FindItemByEmail(ctx, in.Email)

	if existUserEmail != nil {
		l.LogNotify("User with "+in.Email+" exist, failed create", "pkg/repositories/UserRepository/UserRepository.CreateUser")

		return &CustomerUserResponse{
			Code:         globalvars.NotFoundContent,
			Message:      "User exist",
			CustomerUser: nil,
		}, nil
	}

	_user := toCustomerCreateUser(in)

	_user.Password, _ = passwordHelper.HashPassword(_user.Password)

	user, err, message := s.CustomerUserRepo.CreateItem(context.TODO(), _user)

	if err != nil && message == "" {
		return &CustomerUserResponse{
			Code:         globalvars.ServerInternalError,
			Message:      "Internal server error on creating user",
			CustomerUser: nil,
		}, nil
	} else if err != nil && message != "" {
		return &CustomerUserResponse{
			Code:         globalvars.NotFound,
			Message:      message,
			CustomerUser: nil,
		}, nil
	} else {
		role := s.RoleRepo.FindItemById(ctx, user.RoleID)
		businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(ctx, user.BusinessUniverseID)
		customer := s.CustomerRepo.FindItemById(ctx, user.CustomerID)
		return &CustomerUserResponse{
			Code:         globalvars.StatusOK,
			Message:      message,
			CustomerUser: toCustomerUserResponse(user, role, businessUniverse, customer, s),
		}, nil
	}
}

func (s *GrpcServer) GetCustomerUserByEmail(ctx context.Context, in *CustomerUserGetByEmailRequest) (*CustomerUserResponse, error) {

	email := in.Email

	userFounded := s.CustomerUserRepo.FindItemByEmail(context.TODO(), email)
	role := s.RoleRepo.FindItemById(ctx, userFounded.RoleID)
	businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(ctx, userFounded.BusinessUniverseID)
	customer := s.CustomerRepo.FindItemById(ctx, userFounded.CustomerID)

	return &CustomerUserResponse{
		Code:         globalvars.StatusOK,
		Message:      "User get success",
		CustomerUser: toCustomerUserResponse(userFounded, role, businessUniverse, customer, s),
	}, nil
}

func (s *GrpcServer) GetCustomerUserById(ctx context.Context, in *CustomerUserGetByIdRequest) (*CustomerUserResponse, error) {

	id, _ := uuid.Parse(in.Id)

	userFounded := s.CustomerUserRepo.FindItemById(context.TODO(), id)
	role := s.RoleRepo.FindItemById(ctx, userFounded.RoleID)
	businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(ctx, userFounded.BusinessUniverseID)
	customer := s.CustomerRepo.FindItemById(ctx, userFounded.CustomerID)

	return &CustomerUserResponse{
		Code:         globalvars.StatusOK,
		Message:      "User get success",
		CustomerUser: toCustomerUserResponse(userFounded, role, businessUniverse, customer, s),
	}, nil
}

func (s *GrpcServer) UpdateCustomerUser(ctx context.Context, in *CustomerUserUpdateRequest) (*CustomerUserResponse, error) {

	_user := toCustomerUserUpdate(in)
	id, _ := uuid.Parse(in.Id)

	user, err, message := s.CustomerUserRepo.UpdateItem(context.TODO(), _user, id)
	if err != nil && message == "" {
		return &CustomerUserResponse{
			Code:         globalvars.ServerInternalError,
			Message:      "Internal server error on updating user",
			CustomerUser: nil,
		}, nil
	} else if err != nil && message != "" {
		return &CustomerUserResponse{
			Code:         globalvars.NotFound,
			Message:      message,
			CustomerUser: nil,
		}, nil
	} else {
		role := s.RoleRepo.FindItemById(ctx, user.RoleID)
		businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(ctx, user.BusinessUniverseID)
		customer := s.CustomerRepo.FindItemById(ctx, user.CustomerID)
		return &CustomerUserResponse{
			Code:         globalvars.StatusOK,
			Message:      message,
			CustomerUser: toCustomerUserResponse(user, role, businessUniverse, customer, s),
		}, nil
	}
}

func (s *GrpcServer) ChangeCustomerUserPassword(ctx context.Context, in *CustomerUserChangePasswordRequest) (*CustomerUserResponse, error) {

	id, _ := uuid.Parse(in.Id)
	_password := in.Password

	pwd := sha1.New()
	pwd.Write([]byte(_password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))

	password, _ := passwordHelper.HashPassword(in.Password)

	user, err, message := s.CustomerUserRepo.ChangeUserPassword(context.TODO(), id, password)
	if err != nil && message == "" {
		return &CustomerUserResponse{
			Code:         globalvars.ServerInternalError,
			Message:      "Internal server error on updating user",
			CustomerUser: nil,
		}, err
	} else if err != nil && message != "" {
		return &CustomerUserResponse{
			Code:         globalvars.NotFound,
			Message:      message,
			CustomerUser: nil,
		}, err
	} else {
		role := s.RoleRepo.FindItemById(ctx, user.RoleID)
		businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(ctx, user.BusinessUniverseID)
		customer := s.CustomerRepo.FindItemById(ctx, user.CustomerID)
		return &CustomerUserResponse{
			Code:         globalvars.StatusOK,
			Message:      message,
			CustomerUser: toCustomerUserResponse(user, role, businessUniverse, customer, s),
		}, nil
	}
}
