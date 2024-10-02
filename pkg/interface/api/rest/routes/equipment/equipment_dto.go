package equipment

type GetEquipmentQueryDTO struct {
	ID int32 `query:"id"`
}

type ListAllEquipmentQueryDTO struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

type CreateEquipmentQueryDTO struct {
	ID           int32   `json:"id,omitempty"`
	TypeID       *int32  `json:"type_id,omitempty"`
	Name         string  `json:"name,omitempty"`
	SerialNumber string  `json:"serial_number,omitempty"`
	Notes        *string `json:"notes,omitempty"`
}

type DeleteEquipmentTypeQueryDTO struct {
	Id int32 `query:"id"`
}
