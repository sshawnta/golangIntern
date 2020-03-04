package models

type Response struct {
	Error bool `json:"error"`
	ErrorText string `json:"errorText"`
	Data *Data `json:"data"`
	CustomError map[string]string`json:"customError"`
}

type Data struct {
	Res bool `json:"res"`
}
