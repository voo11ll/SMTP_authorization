package globalvars

import "errors"

type CustomError struct {
	Error error
	Enum  string
}

var (
	// Auth
	ErrUserExist = &CustomError{
		Error: errors.New("user exist"),
		Enum:  "ErrUserExist",
	}
	ErrUserNotExist = &CustomError{
		Error: errors.New("user not exist"),
		Enum:  "ErrUserNotExist",
	}
	ErrPriceNotExist = &CustomError{
		Error: errors.New("price not exist"),
		Enum:  "ErrPriceNotExist",
	}
	ErrIncorrectFields = &CustomError{
		Error: errors.New("invalid fields passed"),
		Enum:  "ErrIncorrectFields",
	}
	ErrUpdateUser = &CustomError{
		Error: errors.New("user update error"),
		Enum:  "ErrUpdateUser",
	}
	ErrUpdatePrice = &CustomError{
		Error: errors.New("price update error"),
		Enum:  "ErrUpdatePrice",
	}
	ErrTime = &CustomError{
		Error: errors.New("too frequent use"),
		Enum:  "ErrTime",
	}
	ErrGenTokens = &CustomError{
		Error: errors.New("token creation error"),
		Enum:  "ErrGenTokens",
	}
	ErrGetIp = &CustomError{
		Error: errors.New("IP detection error"),
		Enum:  "ErrGetIp",
	}
	ErrUserStatus = &CustomError{
		Error: errors.New("invalid user status for this operation"),
		Enum:  "ErrStatusUser",
	}
	ErrTokenAccess = &CustomError{
		Error: errors.New("invalid Access token"),
		Enum:  "ErrTokenAccess",
	}
	ErrTokenRefresh = &CustomError{
		Error: errors.New("invalid Refresh token"),
		Enum:  "ErrTokenRefresh",
	}
	ErrTokenExtractData = &CustomError{
		Error: errors.New("failed to get token metadata"),
		Enum:  "ErrTokenExtractData",
	}
	ErrWrongPassword = &CustomError{
		Error: errors.New("wrong password"),
		Enum:  "ErrWrongPassword",
	}
	ErrParseUUID = &CustomError{
		Error: errors.New("parsing error uuid"),
		Enum:  "ErrParseUUID",
	}
	ErrGetFile = &CustomError{
		Error: errors.New("error retrieving the file"),
		Enum:  "ErrGetFile",
	}
	ErrAccess = &CustomError{
		Error: errors.New("error access to method"),
		Enum:  "ErrAccess",
	}
	ErrResponseGRPC = &CustomError{
		Error: errors.New("error get response from grpc server"),
		Enum:  "ErrResponseGRPC",
	}
	ErrDetectCountry = &CustomError{
		Error: errors.New("error detect country by IP"),
		Enum:  "ErrDetectCountry",
	}
	// SendMail
	ErrMailNotSend = &CustomError{
		Error: errors.New("e-mail didn't send"),
		Enum:  "ErrMailNotSend",
	}
	ErrPlatformExist = &CustomError{
		Error: errors.New("platform exist"),
		Enum:  "ErrPlatformExist",
	}
	ErrPlatformNotExist = &CustomError{
		Error: errors.New("platform not exist"),
		Enum:  "ErrPlatformNotExist",
	}
	ErrJustError = &CustomError{
		Error: errors.New("just Error"),
		Enum:  "Error",
	}
	ErrSSHKeyNotExist = &CustomError{
		Error: errors.New("SSHKey not exist"),
		Enum:  "ErrSSHKeyNotExist",
	}
	ErrUpdateSSHKey = &CustomError{
		Error: errors.New("SSHKey update error"),
		Enum:  "ErrUpdateSSHKey",
	}
	ErrAdminRulesNotExist = &CustomError{
		Error: errors.New("admin rules not exist"),
		Enum:  "ErrAdminRulesNotExist",
	}
	ErrAdminRulesExist = &CustomError{
		Error: errors.New("admin rules exist"),
		Enum:  "ErrAdminRulesExist",
	}
	ErrAdminAccess = &CustomError{
		Error: errors.New("not access to admin"),
		Enum:  "ErrAdminAccess",
	}
	ErrPriceAccess = &CustomError{
		Error: errors.New("not access to price"),
		Enum:  "ErrPriceAccess",
	}
	ErrServerInternal = &CustomError{
		Error: errors.New("Internal server error"),
		Enum:  "ErrServerInternal",
	}
)

type CastomResError struct {
	StatusENUM string `json:"statusEnum"`
	Message    string `json:"message"`
}

func GetNotFoundErrors(name string) *CustomError {
	return &CustomError{
		Error: errors.New(name + " not found"),
		Enum:  "Err" + name + "NotFound",
	}
}

func GetExistErrors(name string) *CustomError {
	return &CustomError{
		Error: errors.New(name + " exist"),
		Enum:  "Err" + name + "Exist",
	}
}

func GetUpdateErrors(name string) *CustomError {
	return &CustomError{
		Error: errors.New(name + " update error"),
		Enum:  "Err" + name + "Update",
	}
}
