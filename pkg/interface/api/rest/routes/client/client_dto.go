package client

type ListClientRouteDTO struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

type GetClientByUidRouteDTO struct {
	Id int32 `query:"id"`
}

type CreateClientRouteDTO struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	District      string `json:"district"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Phone         string `json:"phone"`
	Ruc           string `json:"ruc"`
	ContactPerson string `json:"contactPerson"`
	Email         string `json:"email"`
	Website       string `json:"website"`
	AddressLine2  string `json:"addressLine2"`
	PostalCode    string `json:"postalCode"`
	Fax           string `json:"fax"`
	Notes         string `json:"notes"`
}
type UpdateClientRouteDTO struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	District      string `json:"district"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Phone         string `json:"phone"`
	Ruc           string `json:"ruc"`
	ContactPerson string `json:"contactPerson"`
	Email         string `json:"email"`
	Website       string `json:"website"`
	AddressLine2  string `json:"addressLine2"`
	PostalCode    string `json:"postalCode"`
	Fax           string `json:"fax"`
	Notes         string `json:"notes"`
}
type DeleteClientRouteDTO struct {
	ID int32 `json:"id"`
}
