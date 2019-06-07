package models

//
type BusinessObject struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
		Job    string    `bson:"job" json:"job,omitempty"`
	City    string    `bson:"city" json:"city,omitempty"`
	Name    string    `bson:"name" json:"name,omitempty"`
	Company    string    `bson:"company" json:"company,omitempty"`

}

//
type BusinessObjects []BusinessObject


