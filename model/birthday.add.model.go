package model

type BirthayADD struct {
	UserID    string `bson:"_id"`
	Name      string `bson:"name"`
	Birthdate int    `bson:"birthdate"`
	Birthday  string `bson:"birthday"`
}
