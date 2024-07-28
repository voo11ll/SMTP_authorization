package user

import (
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/passwordHelper"
	"auth/auth_back/pkg/logger"
	context "context"

	businessUniverseRepository "auth/auth_back/pkg/repositories/businessUniverse"
	contactTypeRepository "auth/auth_back/pkg/repositories/contactType"
	customerRepository "auth/auth_back/pkg/repositories/customer"
	customerUserRepository "auth/auth_back/pkg/repositories/customerUser"
	notificationRepository "auth/auth_back/pkg/repositories/notification"
	userRepository "auth/auth_back/pkg/repositories/user"

	roleRepository "auth/auth_back/pkg/repositories/role"
	"auth/auth_back/pkg/services/grpc/notification"

	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcServer struct {
	UserRepo             *userRepository.UserRepository
	RoleRepo             *roleRepository.RoleRepository
	CustomerUserRepo     *customerUserRepository.CustomerUserRepository
	BusinessUniverseRepo *businessUniverseRepository.BusinessUniverseRepository
	ContactTypeRepo      *contactTypeRepository.ContactTypeRepository
	CustomerRepo         *customerRepository.CustomerRepository
	NotifyRepo           *notificationRepository.NotificationRepository
	UnimplementedB24UserServiceServer
}

var l = logger.Logger{}

func (s *GrpcServer) SignUp(ctx context.Context, in *SignUpRequest) (*UserResponse, error) {

	existUserEmail := s.UserRepo.FindItemByEmail(ctx, in.Email)

	if existUserEmail != nil {
		l.LogNotify("User with "+in.Email+" exist, failed create", "pkg/repositories/UserRepository/UserRepository.CreateUser")
		role := s.RoleRepo.FindItemById(ctx, existUserEmail.RoleID)
		return &UserResponse{
			Code:    globalvars.ContentExist,
			Message: "User exist",
			User:    toUserResponse(existUserEmail, role),
		}, nil
	}

	// Get admin role for main user
	role := s.RoleRepo.FindItemByName(context.Background(), "Administrator")

	// buId, _ := uuid.Parse(in.BusinessUniverseID)
	_user := toSignUp(in)

	// Create business universe for main user
	newUniverse := GetAndCreateBusinessUniverseGrpcClient("")
	// buId, _ = uuid.Parse(newUniverse.BusinessUniverse.Id)
	_user.BusinessUniverseID = newUniverse.BusinessUniverse.Id
	_user.RoleID = role.ID.String()

	_user.Password, _ = passwordHelper.HashPassword(_user.Password)

	user, message, err := s.UserRepo.CreateItem(context.TODO(), _user)
	if err != nil && message == "" {
		return &UserResponse{
			Code:    globalvars.ServerInternalError,
			Message: "Internal server error on creating user",
			User:    nil,
		}, nil
	} else if err != nil && message != "" {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: message,
			User:    nil,
		}, nil
	}
	role = s.RoleRepo.FindItemById(ctx, user.RoleID)
	hashKey, _ := passwordHelper.HashPassword(user.Email)
	newUser := s.UserRepo.FindItemByEmail(ctx, user.Email)
	if newUser == nil {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: "User not found",
			User:    nil,
		}, nil
	}
	_, err = s.NotifyRepo.CreateMailConfirmationLink(&notificationRepository.CreateMailConfrimationLink{
		UserId:    newUser.ID,
		HashKey:   hashKey,
		SendTry:   0,
		Confirmed: false,
	})
	if err != nil {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: err.Error(),
			User:    nil,
		}, nil
	}

	confirmLink := s.NotifyRepo.FindLinkByUserId(newUser.ID)
	if confirmLink == nil {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: "Link not found",
			User:    nil,
		}, nil
	}

	clientConn, err := grpc.Dial(viper.GetString("grpc.notification.host")+":"+viper.GetString("grpc.notification.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.LogError(err.Error(), "pkg/service/grpc/notification.ClientConnect")
		return &UserResponse{
			Code:    globalvars.ServerInternalError,
			Message: err.Error(),
			User:    nil,
		}, nil
	}

	defer clientConn.Close()

	notificationServiceClient := notification.NewNotificationServiceClient(clientConn)

	response, err := notificationServiceClient.SendMailConfirmLink(ctx, &notification.SendMailConfirmLinkRequest{
		LinkId:    confirmLink.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: err.Error(),
			User:    nil,
		}, nil
	}

	return &UserResponse{
		Code:    globalvars.StatusOK,
		Message: response.Message,
		User:    toUserResponse(user, role),
	}, nil

}

func (s *GrpcServer) SignIn(ctx context.Context, in *SignInRequest) (*UserResponse, error) {

	user := s.UserRepo.FindItemByEmail(context.TODO(), in.Email)
	if user == nil {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: globalvars.ErrUserNotExist.Error.Error(),
			User:    nil,
		}, nil
	}

	mailConfirmed := s.NotifyRepo.FindLinkByUserId(user.ID)
	if mailConfirmed == nil {
		return &UserResponse{
			Code:    globalvars.NotFound,
			Message: "User has not been finished registration",
			User:    nil,
		}, nil
	}

	if !mailConfirmed.Confirmed {
		return &UserResponse{
			Code:    globalvars.Unauthorized,
			Message: "User did not confirm his email",
			User:    nil,
		}, nil
	}

	correct := passwordHelper.CheckPasswordHash(in.GetPassword(), user.Password)
	if !correct {
		return &UserResponse{
			Code:    globalvars.Unauthorized,
			Message: globalvars.ErrWrongPassword.Error.Error(),
			User:    nil,
		}, nil
	} else {
		role := s.RoleRepo.FindItemById(ctx, user.RoleID)
		return &UserResponse{
			Code:    globalvars.StatusOK,
			Message: "User logged in success",
			User:    toUserResponse(user, role),
		}, nil
	}
}
