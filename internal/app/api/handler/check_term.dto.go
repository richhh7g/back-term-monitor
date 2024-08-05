package handler

type CheckTermBody struct {
	Email  string   `json:"email" validate:"required,email" example:"email@example.com" errormgs:"Email inválido"`
	Termos []string `json:"termos" validate:"required" example:"Term One,Term Two" errormgs:"Termos inválidos"`
} // @name ChecarTermoBody
