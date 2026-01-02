package model

type Project struct {
	ID       int    `json:"id"`
	UID      int    `json:"uid"`
	CreateAT int    `json:"create_at"`
	Title    string `json:"title"`
}
