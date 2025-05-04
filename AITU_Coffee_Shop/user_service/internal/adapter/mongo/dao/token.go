package dao

type Token struct {
	refresh string `bson:"refresh"`
	access  string `bson:"access"`
}
