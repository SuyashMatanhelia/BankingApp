package models

type Details struct {
	CustId int `json:"id"`
	Amt    int `json:"amt"`
	Term   int `json:"term"`
}
