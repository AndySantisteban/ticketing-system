package client_commands

type UpdateClientQuery struct {
	ID            int32
	Name          string
	Address       string
	District      string
	City          string
	Country       string
	Phone         string
	Ruc           string
	ContactPerson string
	Email         string
	Website       string
	AddressLine2  string
	PostalCode    string
	Fax           string
	Notes         string
}

type UpdateClientResponse = error
