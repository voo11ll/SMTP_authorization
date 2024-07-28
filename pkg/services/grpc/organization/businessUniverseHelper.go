package organization

import (
	businessUniverseRepository "auth/auth_back/pkg/repositories/businessUniverse"
)

func toItemUpdateModel(in *BusinessUniverseUpdateRequest) *businessUniverseRepository.BusinessUniverse {
	return &businessUniverseRepository.BusinessUniverse{
		Name:   in.GetName(),
		Domain: in.GetDomain(),
	}
}
