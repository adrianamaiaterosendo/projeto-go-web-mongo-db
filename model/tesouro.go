package model

import "time"

type TesouroInput struct {
	ID              string    `json:"id"`
	Confrade        string    `json:"confrade"`
	DataCadastro    time.Time `json:"data_cadastro"`
	Comunhao        int       `json:"comunhao"`
	Vigilia         int       `json:"vigilia"`
	Adoracao        int       `json:"adoracao"`
	Meditacao       int       `json:"meditacao"`
	Rosario         int       `json:"rosario"`
	Terco           int       `json:"terco"`
	OficioLaudes    int       `json:"oficio_laudes"`
	OficioVesperas  int       `json:"oficio_vesperas"`
	OficioSexta     int       `json:"oficio_sexta"`
	OficioCompletas int       `json:"oficio_completas"`
	Jejum           int       `json:"jejum"`
	Sacrificio      int       `json:"sacrificio"`
	Esmola          int       `json:"esmola"`
	Dia             string    `json:"dia"`
	Mes             string    `json:"mes"`
	Ano             string    `json:"ano"`
}

type Subscription struct {
	Name  string
	Email string
}
