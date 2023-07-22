package types

type SensitiveWord struct {
	Word    string `json:"word"`
	Indexes []int  `json:"indexes"`
	Length  int    `json:"length"`
}
