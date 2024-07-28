package organization

import (
	models "auth/auth_back/models/organization"
	companyRepository "auth/auth_back/pkg/repositories/company"
	"context"
)

func toCompanyCreateModel(in *CreateRequest) *companyRepository.Company {
	return &companyRepository.Company{
		Name:               in.Name,
		FullName:           in.FullName,
		INN:                in.Inn,
		KPP:                in.Kpp,
		LegalAddress:       in.LegalAddress,
		Banks:              toBanksCreateModel(in.Banks),
		ContactInfos:       toCompanyContactInfosCreateModel(in.ContactInfos),
		BusinessUniverseId: in.BusinessUniverseID,
	}
}

func toCompaniesGet(u []*models.Company, s *GrpcServer) []*Company {
	var out = make([]*Company, len(u))

	for i, _u := range u {
		businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(context.TODO(), _u.BusinessUniverseID)
		out[i] = &Company{
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
			Banks:        toBanksGet(_u.Banks),
			ContactInfos: toCompanyContactInfosGet(_u.ContactInfos, s),
		}
	}

	return out
}

func toCompanyGet(u *models.Company, s *GrpcServer) *Company {
	businessUniverse := s.BusinessUniverseRepo.FindBusinessUniverseById(context.TODO(), u.BusinessUniverseID)
	return &Company{
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
		Banks:        toBanksGet(u.Banks),
		ContactInfos: toCompanyContactInfosGet(u.ContactInfos, s),
	}
}

func toCompanyUpdateModel(in *UpdateRequest) *companyRepository.Company {
	return &companyRepository.Company{
		Name:         in.GetName(),
		FullName:     in.FullName,
		INN:          in.Inn,
		KPP:          in.Kpp,
		LegalAddress: in.LegalAddress,
	}
}
