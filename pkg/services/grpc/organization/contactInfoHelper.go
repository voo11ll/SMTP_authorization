package organization

import (
	models "auth/auth_back/models/organization"
	companyRepository "auth/auth_back/pkg/repositories/company"
	customerRepository "auth/auth_back/pkg/repositories/customer"
	"context"

	"github.com/google/uuid"
)

func getContactType(s *GrpcServer, contactTypeId string) *models.ContactType {
	id, _ := uuid.Parse(contactTypeId)
	return s.ContactTypeRepo.FindItemById(context.TODO(), id)
}

/**
* Company contact info
**/

func toCompanyContactInfoCreateModel(in *AddCompanyContactInfoRequest) *companyRepository.ContactInfo {
	return &companyRepository.ContactInfo{
		ContactTypeID: in.ContactInfo.ContactTypeID,
		Value:         in.ContactInfo.Value,
		CompanyId:     in.CompanyId,
	}
}

func toCompanyContactInfosCreateModel(in []*ContactInfoCreateType) []*companyRepository.ContactInfo {

	var out = make([]*companyRepository.ContactInfo, len(in))

	for i, _u := range in {
		out[i] = &companyRepository.ContactInfo{
			ContactTypeID: _u.ContactTypeID,
			Value:         _u.Value,
		}
	}
	return out
}

func toCompanyContactInfosGet(u []*models.CompanyContactInfo, s *GrpcServer) []*ContactInfo {
	var out = make([]*ContactInfo, len(u))

	for i, _u := range u {
		contactType := getContactType(s, _u.ContactTypeID.String())
		out[i] = &ContactInfo{
			Id: _u.ID.String(),
			ContactType: &ContactType{
				Id:     contactType.ID.String(),
				Name:   contactType.Name,
				Notion: contactType.Notion,
			},
			Value: _u.Value,
		}
	}

	return out
}

func toCompanyContactInfoGet(u *models.CompanyContactInfo, s *GrpcServer) *ContactInfo {
	contactType := getContactType(s, u.ContactTypeID.String())
	return &ContactInfo{
		Id: u.ID.String(),
		ContactType: &ContactType{
			Id:     contactType.ID.String(),
			Name:   contactType.Name,
			Notion: contactType.Notion,
		},
		Value: u.Value,
	}
}

func toCompanyContactInfoUpdateModel(in *UpdateContactInfoRequest) *companyRepository.ContactInfo {
	return &companyRepository.ContactInfo{
		ContactTypeID: in.ContactTypeID,
		Value:         in.Value,
	}
}

/**
* Customer contact info
**/

func toCustomerContactInfoCreateModel(in *AddCustomerContactInfoRequest) *customerRepository.ContactInfo {
	return &customerRepository.ContactInfo{
		ContactTypeID: in.ContactInfo.ContactTypeID,
		Value:         in.ContactInfo.Value,
		CustomerId:    in.CustomerId,
	}
}

func toCustomerContactInfosCreateModel(in []*ContactInfoCreateType) []*customerRepository.ContactInfo {

	var out = make([]*customerRepository.ContactInfo, len(in))

	for i, _u := range in {
		out[i] = &customerRepository.ContactInfo{
			ContactTypeID: _u.ContactTypeID,
			Value:         _u.Value,
		}
	}
	return out
}

func toCustomerContactInfosGet(u []*models.CustomerContactInfo, s *GrpcServer) []*ContactInfo {
	var out = make([]*ContactInfo, len(u))

	for i, _u := range u {
		contactType := getContactType(s, _u.ContactTypeID.String())
		out[i] = &ContactInfo{
			Id: _u.ID.String(),
			ContactType: &ContactType{
				Id:     contactType.ID.String(),
				Name:   contactType.Name,
				Notion: contactType.Notion,
			},
			Value: _u.Value,
		}
	}

	return out
}

func toCustomerContactInfoGet(u *models.CustomerContactInfo, s *GrpcServer) *ContactInfo {
	contactType := getContactType(s, u.ContactTypeID.String())
	return &ContactInfo{
		Id: u.ID.String(),
		ContactType: &ContactType{
			Id:     contactType.ID.String(),
			Name:   contactType.Name,
			Notion: contactType.Notion,
		},
		Value: u.Value,
	}
}

func toCustomerContactInfoUpdateModel(in *UpdateContactInfoRequest) *customerRepository.ContactInfo {
	return &customerRepository.ContactInfo{
		ContactTypeID: in.ContactTypeID,
		Value:         in.Value,
	}
}
