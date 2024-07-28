package organizationHttp

import (
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/middleware"
	businessUniverseHandler "auth/auth_back/pkg/services/http/organization/businessUniverse"
	bankHandler "auth/auth_back/pkg/services/http/organization/company/bank"
	companyHandler "auth/auth_back/pkg/services/http/organization/company/company"
	contactInfoHandler "auth/auth_back/pkg/services/http/organization/company/contactInfo"
	contactTypeHandler "auth/auth_back/pkg/services/http/organization/contactType"
	customerBankHandler "auth/auth_back/pkg/services/http/organization/customer/bank"
	customerContactInfoHandler "auth/auth_back/pkg/services/http/organization/customer/contactInfo"
	customerHandler "auth/auth_back/pkg/services/http/organization/customer/customer"
)

var Routes = httpServerHelper.Routes{
	/**
	* Business Universe routes
	**/
	httpServerHelper.Route{
		Name:        "CreateBusinessUniverse",
		Method:      "POST",
		Pattern:     "/organization/universe",
		HandlerFunc: businessUniverseHandler.CreateBusinessUniverse,
	},
	httpServerHelper.Route{
		Name:        "GetBusinessUniverse",
		Method:      "GET",
		Pattern:     "/organization/universe",
		HandlerFunc: businessUniverseHandler.GetBusinessUniverse,
	},
	httpServerHelper.Route{
		Name:        "UpdateBusinessUniverse",
		Method:      "PUT",
		Pattern:     "/organization/universe",
		HandlerFunc: businessUniverseHandler.UpdateBusinessUniverse,
	},
	/**
	* Contact type routes
	**/
	httpServerHelper.Route{
		Name:        "CreateContactType",
		Method:      "POST",
		Pattern:     "/organization/contact-type",
		HandlerFunc: contactTypeHandler.CreateContactType,
	},
	httpServerHelper.Route{
		Name:        "GetContactType",
		Method:      "GET",
		Pattern:     "/organization/contact-type",
		HandlerFunc: contactTypeHandler.GetContactType,
	},
	httpServerHelper.Route{
		Name:        "GetContactTypes",
		Method:      "GET",
		Pattern:     "/organization/contact-type/list",
		HandlerFunc: contactTypeHandler.GetContactTypes,
	},
	httpServerHelper.Route{
		Name:        "UpdateContactType",
		Method:      "PUT",
		Pattern:     "/organization/contact-type",
		HandlerFunc: contactTypeHandler.UpdateContactType,
	},
	/**
	* Company routes
	**/
	httpServerHelper.Route{
		Name:        "CreateCompany",
		Method:      "POST",
		Pattern:     "/organization/company",
		HandlerFunc: companyHandler.CreateCompany,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetCompany",
		Method:      "GET",
		Pattern:     "/organization/company",
		HandlerFunc: companyHandler.GetCompany,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetCompanies",
		Method:      "GET",
		Pattern:     "/organization/company/list",
		HandlerFunc: companyHandler.GetCompanies,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "UpdateCompany",
		Method:      "PUT",
		Pattern:     "/organization/company",
		HandlerFunc: companyHandler.UpdateCompany,
		Middleware:  middleware.UserMiddleware,
	},
	/**
	* Company bank routes
	**/
	httpServerHelper.Route{
		Name:        "CreateBank",
		Method:      "POST",
		Pattern:     "/organization/company/bank",
		HandlerFunc: bankHandler.CreateBank,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetBank",
		Method:      "GET",
		Pattern:     "/organization/company/bank",
		HandlerFunc: bankHandler.GetBank,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetBanks",
		Method:      "GET",
		Pattern:     "/organization/company/bank/list",
		HandlerFunc: bankHandler.GetBanks,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "UpdateBank",
		Method:      "PUT",
		Pattern:     "/organization/company/bank",
		HandlerFunc: bankHandler.UpdateBank,
		Middleware:  middleware.UserMiddleware,
	},
	/**
	* Company contact info routes
	**/
	httpServerHelper.Route{
		Name:        "CreateContactInfo",
		Method:      "POST",
		Pattern:     "/organization/company/contact-info",
		HandlerFunc: contactInfoHandler.CreateContactInfo,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetContactInfo",
		Method:      "GET",
		Pattern:     "/organization/company/contact-info",
		HandlerFunc: contactInfoHandler.GetContactInfo,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetContactInfos",
		Method:      "GET",
		Pattern:     "/organization/company/contact-info/list",
		HandlerFunc: contactInfoHandler.GetContactInfos,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "UpdateContactInfo",
		Method:      "PUT",
		Pattern:     "/organization/company/contact-info",
		HandlerFunc: contactInfoHandler.UpdateContactInfo,
		Middleware:  middleware.UserMiddleware,
	},
	/**
	* Customer routes
	**/
	httpServerHelper.Route{
		Name:        "CreateCustomer",
		Method:      "POST",
		Pattern:     "/organization/customer",
		HandlerFunc: customerHandler.CreateCustomer,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetCustomer",
		Method:      "GET",
		Pattern:     "/organization/customer",
		HandlerFunc: customerHandler.GetCustomer,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetCustomers",
		Method:      "GET",
		Pattern:     "/organization/customer/list",
		HandlerFunc: customerHandler.GetCustomers,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "UpdateCustomer",
		Method:      "PUT",
		Pattern:     "/organization/customer",
		HandlerFunc: customerHandler.UpdateCustomer,
		Middleware:  middleware.UserMiddleware,
	},
	/**
	* Customer bank routes
	**/
	httpServerHelper.Route{
		Name:        "CreateBank",
		Method:      "POST",
		Pattern:     "/organization/customer/bank",
		HandlerFunc: customerBankHandler.CreateBank,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetBank",
		Method:      "GET",
		Pattern:     "/organization/customer/bank",
		HandlerFunc: customerBankHandler.GetBank,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetBanks",
		Method:      "GET",
		Pattern:     "/organization/customer/bank/list",
		HandlerFunc: customerBankHandler.GetBanks,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "UpdateBank",
		Method:      "PUT",
		Pattern:     "/organization/customer/bank",
		HandlerFunc: customerBankHandler.UpdateBank,
		Middleware:  middleware.UserMiddleware,
	},
	/**
	* Customer contact info routes
	**/
	httpServerHelper.Route{
		Name:        "CreateContactInfo",
		Method:      "POST",
		Pattern:     "/organization/customer/contact-info",
		HandlerFunc: customerContactInfoHandler.CreateContactInfo,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetContactInfo",
		Method:      "GET",
		Pattern:     "/organization/customer/contact-info",
		HandlerFunc: customerContactInfoHandler.GetContactInfo,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "GetContactInfos",
		Method:      "GET",
		Pattern:     "/organization/customer/contact-info/list",
		HandlerFunc: customerContactInfoHandler.GetContactInfos,
		Middleware:  middleware.UserMiddleware,
	},
	httpServerHelper.Route{
		Name:        "UpdateContactInfo",
		Method:      "PUT",
		Pattern:     "/organization/customer/contact-info",
		HandlerFunc: customerContactInfoHandler.UpdateContactInfo,
		Middleware:  middleware.UserMiddleware,
	},
}
