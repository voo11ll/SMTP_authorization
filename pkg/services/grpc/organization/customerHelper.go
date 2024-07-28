package organization

import (
	models "auth/auth_back/models/organization"
	customerRepository "auth/auth_back/pkg/repositories/customer"
	"context"
)

func toCustomerCreateModel(in *CreateRequest) *customerRepository.Customer {
	return &customerRepository.Customer{
		Name:               in.Name,
		FullName:           in.FullName,
		INN:                in.Inn,
		KPP:                in.Kpp,
		LegalAddress:       in.LegalAddress,
		Banks:              toCustomerBanksCreateModel(in.Banks),
		ContactInfos:       toCustomerContactInfosCreateModel(in.ContactInfos),
		BusinessUniverseId: in.BusinessUniverseID,
	}
}

func toCustomersGet(u []*models.Customer, s *GrpcServer) []*Customer {
	var out = make([]*Customer, len(u))

	for i, _u := range u {
		businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(context.TODO(), _u.BusinessUniverseID)
		out[i] = &Customer{
			Id:           _u.ID.String(),
			Name:         _u.Name,
			FullName:     _u.FullName,
			Inn:          _u.INN,
			Kpp:          _u.KPP,
			LegalAddress: _u.LegalAddress,
			BusinessUniverse: &BusinessUniverse{
				Id:     businessUniverse.ID.String(),
				Name:   businessUniverse.Name,
				Domain: businessUniverse.Domain,
			},
			Banks:        toCustomerBanksGet(_u.Banks),
			ContactInfos: toCustomerContactInfosGet(_u.ContactInfos, s),
		}
	}

	return out
}

func toCustomerGet(u *models.Customer, s *GrpcServer) *Customer {
	businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(context.TODO(), u.BusinessUniverseID)
	return &Customer{
		Id:           u.ID.String(),
		Name:         u.Name,
		FullName:     u.FullName,
		Inn:          u.INN,
		Kpp:          u.KPP,
		LegalAddress: u.LegalAddress,
		BusinessUniverse: &BusinessUniverse{
			Id:     businessUniverse.ID.String(),
			Name:   businessUniverse.Name,
			Domain: businessUniverse.Domain,
		},
		Banks:        toCustomerBanksGet(u.Banks),
		ContactInfos: toCustomerContactInfosGet(u.ContactInfos, s),
	}
}

func toCustomerUpdateModel(in *UpdateRequest) *customerRepository.Customer {
	return &customerRepository.Customer{
		Name:         in.GetName(),
		FullName:     in.FullName,
		INN:          in.Inn,
		KPP:          in.Kpp,
		LegalAddress: in.LegalAddress,
	}
}
