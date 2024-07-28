package organization

import (
	// "auth/auth_back/pkg/logger"
	businessUniverseRepository "auth/auth_back/pkg/repositories/businessUniverse"
	companyRepository "auth/auth_back/pkg/repositories/company"
	contactTypeRepository "auth/auth_back/pkg/repositories/contactType"
	customerRepository "auth/auth_back/pkg/repositories/customer"
)

type GrpcServer struct {
	BusinessUniverseRepo *businessUniverseRepository.BusinessUniverseRepository
	ContactTypeRepo      *contactTypeRepository.ContactTypeRepository
	CompanyRepo          *companyRepository.CompanyRepository
	CustomerRepo         *customerRepository.CustomerRepository
	UnimplementedB24OrganizationServiceServer
}

// var l = logger.Logger{}
