package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserProfile struct {
	Id            primitive.ObjectID `bson:"_id"`
	Username      string             `bson:"username"`
	Email         string             `bson:"email"`
	Fullname      string             `bson:"fullname"`
	Bio           string             `bson:"bio"`
	ImageHash     string             `bson:"imageHash"`
	Location      string             `bson:"location"`
	EmailVerified bool               `bson:"emailVerified"`
	CreatedAt     primitive.DateTime `bson:"createdAt"`
}

type Social struct {
	Id          primitive.ObjectID `bson:"_id"`
	Label       string             `bson:"label"`
	Slug        string             `bson:"slug"`
	Index       uint64             `bson:"index"`
	Href        string             `json:"href"`
	UserId      primitive.ObjectID `bson:"userId"`
	CreatedAt   primitive.DateTime `bson:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt"`
	PlatformKey string             `bson:"platformKey"`
}

type Link struct {
	Id          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Favicon     string             `bson:"favicon"`
	Href        string             `bson:"href"`
	Thumb       string             `bson:"thumb"`
	Description string             `bson:"description"`
	Style       string             `bson:"style"`
	UserId      primitive.ObjectID `bson:"userId"`
	CreatedAt   primitive.DateTime `bson:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt"`
}

type PageConfig struct {
	Id           primitive.ObjectID `bson:"_id"`
	TemplateName string             `bson:"templateName"`
	UserId       primitive.ObjectID `bson:"userId"`
	ThemeMode    string             `bson:"themeMode"`
	Style        map[string]string  `bson:"style"`
	Title        string             `bson:"title"`
	Description  string             `bson:"description"`
	Theme        string             `bson:"theme"`
}

type PageTheme struct {
	Id        primitive.ObjectID `bson:"_id"`
	ThemeMode string             `bson:"themeMode"`
	Style     interface{}        `bson:"style"`
	Theme     string             `bson:"theme"`
}

type UserData struct {
	Id     primitive.ObjectID `bson:"_id"`
	User   *UserProfile       `bson:"user"`
	Social []*Social          `bson:"social"`
	Links  []*Link            `bson:"links"`
	Page   *PageConfig        `bson:"page"`
	Css    string             `json:"css"`
	Body   string             `json:"body"`
}

type SocialPlatform struct {
	Key         string
	Type        string
	Name        string
	Hint        string
	Hostname    string
	TemplateUrl string
}
