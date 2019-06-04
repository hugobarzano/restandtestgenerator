package models

//
type BusinessObject struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
		Alias    string    `bson:"alias" json:"alias,omitempty"`
	Reina    bool    `bson:"reina" json:"reina,omitempty"`
	Age    float64    `bson:"age" json:"age,omitempty"`
	Color    string    `bson:"color" json:"color,omitempty"`
	Name    string    `bson:"name" json:"name,omitempty"`

}

//
type BusinessObjects []BusinessObject


