package public

type GetTrackingInfoDTO struct {
	Id int32 `query:"id"`
}

type TecSupportDTO struct {
	Correo  string `json:"correo"`
	Asunto  string `json:"asunto"`
	Mensaje string `json:"mensaje"`
}
