package models

//Progress model
type Progress struct {
	Progress float64  `json:"progress" bson:"progress"`
	Active bool       `json:"active" bson:"active"`
}
