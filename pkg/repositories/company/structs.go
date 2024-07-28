package companyRepository

type Company struct {
	Name               string         `json:"name"`
	FullName           string         `json:"fullName"`
	INN                int32          `json:"inn"`
	KPP                int32          `json:"kpp"`
	LegalAddress       string         `json:"legalAddress"`
	Banks              []*Bank        `json:"banks"`
	ContactInfos       []*ContactInfo `json:"contactInfos"`
	BusinessUniverseId string         `json:"businessUniverseId"`
}

type ContactInfo struct {
	ContactTypeID string `json:"contactTypeId"`
	Value         string `json:"value"`
	CompanyId     string `json:"companyId"`
}

type Bank struct {
	Name              string `json:"name"`
	AccountNumber     string `json:"accountNumber"`
	Bik               string `json:"bik"`
	CorrAccountNumber string `json:"corrAccountNumber"`
	IsOpen            bool   `json:"isOpen"`
	CompanyId         string `json:"companyId"`
}
