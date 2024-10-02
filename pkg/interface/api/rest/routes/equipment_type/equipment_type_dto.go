package equipment_type

type GetEquipmentTypeQueryDTO struct {
	ID int32 `query:"id"`
}

type ListAllEquipmentTypeQueryDTO struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

type CreateEquipmentTypeQueryDTO struct {
	ID   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DeleteEquipmentTypeRouteDTO struct {
	ID int32 `query:"id"`
}
