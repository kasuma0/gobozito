package controller

import (
	"context"

	"github.com/kasuma0/gobozito/model"
	"github.com/kasuma0/gobozito/persistence"
)

func BirthdayADD(ctx context.Context, birthObj model.BirthayADD) error {
	mong := persistence.MongoDb{
		Collection: "birthday",
		DB:         "gobozito",
	}
	_, err := mong.InsertDocument(ctx, birthObj)
	if err != nil {
		return err
	}
	return nil
}
