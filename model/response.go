package model

//Info response struct of info API
type Info struct {
	UID     string `json:"uid"`
	Name    string `json:"name"`
	Faculty string `json:"faculty"`
	Class   string `json:"class"`
}

//Dorm dormitory, response struct of dorm API
type Dorm struct {
	Name   string `json:"name"`
	Campus string `json:"campus"`
	House  string `json:"house"`
	Room   string `json:"room"`
	Bed    string `json:"bed"`
}
