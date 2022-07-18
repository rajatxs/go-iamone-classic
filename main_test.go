package main

import (
	"testing"

	"github.com/rajatxs/go-iamone/common"
	"github.com/rajatxs/go-iamone/db"
	"github.com/rajatxs/go-iamone/models"
	"github.com/rajatxs/go-iamone/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	common.Ensure(db.Connect(), "Failed to make database connection")
}

func TestUserDataService(t *testing.T) {
	const targetUsername = "rxx"

	var (
		udata models.UserData
		err   error
	)

	err = services.ReadUserDataByUsername(targetUsername, &udata)

	switch {
	case err == mongo.ErrNoDocuments:
		t.Errorf("User not found \"%s\"", targetUsername)

	case err != nil:
		t.Error(err)

	case udata.User.Username != targetUsername:
		t.Errorf("Getting incorrect data of user \"%s\"", targetUsername)
	}
}

func TestPageThemeService(t *testing.T) {
	const targetPageConfigId = "61d57cd38beab8098aa3af66"

	var (
		pageConfigId primitive.ObjectID
		theme        models.PageTheme
		err          error
	)

	pageConfigId, err = primitive.ObjectIDFromHex(targetPageConfigId)

	if err != nil {
		t.Errorf("Couldn't get ObjectID from %s\n%v", targetPageConfigId, err)
	}

	err = services.ReadPageThemeByPageId(pageConfigId, &theme)

	switch {
	case err == mongo.ErrNoDocuments:
		t.Errorf("Theme not found %v", pageConfigId)
	case err != nil:
		t.Error(err)
	case theme.Id != pageConfigId:
		t.Errorf("Getting incorrect data for theme %s", pageConfigId)
	}
}
