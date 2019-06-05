package models

//
type BusinessObject struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
		Name    string    `bson:"name" json:"name,omitempty"`
	Grados    string    `bson:"grados" json:"grados,omitempty"`
	Ingradientes    string    `bson:"ingradientes" json:"ingradientes,omitempty"`
	Sabores    string    `bson:"sabores" json:"sabores,omitempty"`

}

//
type BusinessObjects []BusinessObject


