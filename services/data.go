package services

import (
	"fmt"

	"github.com/aymerick/raymond"
	"github.com/rajatxs/go-iamone/config"
	"github.com/rajatxs/go-iamone/db"
	"github.com/rajatxs/go-iamone/logger"
	"github.com/rajatxs/go-iamone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Resolves href from social template url */
func ResolveSocialHref(udata *models.UserData) (err error) {
	var (
		platf models.SocialPlatform
		href  string
	)

	for _, v := range udata.Social {
		platf = config.SocialPlatforms[v.PlatformKey]

		if len(platf.Key) == 0 {
			continue
		}

		href, err = raymond.Render(platf.TemplateUrl, v)

		if err != nil {
			logger.Err(fmt.Sprintf("Couldn't parse template url %s", platf.TemplateUrl))
			continue
		}

		v.Href = href
	}

	return err
}

/* Reads UserData by given Username */
func ReadUserDataByUsername(uname string, udata *models.UserData) error {
	return db.FindOneDoc(
		db.USER_DATA_COLLECTION,
		bson.D{{Key: "user.username", Value: uname}},
		udata)
}

/* Reads User Theme by given Page config Id */
func ReadPageThemeByPageId(pid primitive.ObjectID, theme *models.PageTheme) error {
	return db.FindOneDoc(
		db.PAGE_THEME_COLLECTION,
		bson.D{{Key: "_id", Value: pid}},
		theme)
}
