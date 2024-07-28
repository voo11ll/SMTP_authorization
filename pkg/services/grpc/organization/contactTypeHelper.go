package organization

import (
	models "auth/auth_back/models/organization"
	contactTypeRepository "auth/auth_back/pkg/repositories/contactType"
)

func toContactTypeCreateModel(in *ContactTypeCreateRequest) *contactTypeRepository.ContactType {
	return &contactTypeRepository.ContactType{
		Name:   in.GetName(),
		Notion: in.GetNotion(),
	}
}

func toContactTypesGet(u []*models.ContactType) []*ContactType {
	var out = make([]*ContactType, len(u))

	for i, _u := range u {
		out[i] = &ContactType{
			Id:     _u.ID.String(),
			Name:   _u.Name,
			Notion: _u.Notion,
		}
	}

	return out
}

func toContactTypeUpdateModel(in *ContactTypeUpdateRequest) *contactTypeRepository.ContactType {
	return &contactTypeRepository.ContactType{
		Name:   in.GetName(),
		Notion: in.GetNotion(),
	}
}
