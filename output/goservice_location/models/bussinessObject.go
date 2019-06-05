package models

//
type BusinessObject struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
		Player    string    `bson:"player" json:"player,omitempty"`
	X    float64    `bson:"x" json:"x,omitempty"`
	Y    float64    `bson:"y" json:"y,omitempty"`
	Z    float64    `bson:"z" json:"z,omitempty"`

}

//
type BusinessObjects []BusinessObject


