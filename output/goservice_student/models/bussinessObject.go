package models

//
type BusinessObject struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
		Name    string    `bson:"name" json:"name,omitempty"`
	Attendance    bool    `bson:"attendance" json:"attendance,omitempty"`

}

//
type BusinessObjects []BusinessObject


