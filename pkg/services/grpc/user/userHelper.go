package user

import (
	userRepository "auth/auth_back/pkg/repositories/user"
)

func toSignUp(u *SignUpRequest) *userRepository.User {
	return &userRepository.User{
		Email:     u.GetEmail(),
		Password:  u.GetPassword(),
		FirstName: u.GetFirstName(),
		LastName:  u.GetLastName(),
		Phone:     u.GetPhone(),
	}
}

func toUserUpdate(u *UserUpdateRequest) *userRepository.User {
	return &userRepository.User{
		FirstName:  u.GetFirstName(),
		LastName:   u.GetLastName(),
		SecondName: u.GetSecondName(),
		Phone:      u.GetPhone(),
	}
}
