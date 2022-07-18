package services

import (
	"bytes"
	"os"
	"path"

	"github.com/aymerick/raymond"
	"github.com/rajatxs/go-iamone/common"
	"github.com/rajatxs/go-iamone/config"
	"github.com/rajatxs/go-iamone/logger"
	"github.com/rajatxs/go-iamone/models"
)

/* Returns theme properties of given theme name  */
func GetTemplateThemeProps(tname string) (tprops map[string]string, err error) {
	err = common.ReadJSON(path.Join(config.THEMES_DIR, tname), &tprops)
	return tprops, err
}

/* Returns template source of given template name */
func GetTemplateSource(name string) (templ string, err error) {
	var (
		cont []byte
		buff bytes.Buffer
	)
	cont, err = os.ReadFile(path.Join(config.TEMPLATE_DIR, name))

	if err != nil {
		return "", err
	}

	_, err = buff.Write(cont)
	templ = buff.String()
	return templ, err
}

/* Read template theme and returns compiled version of stylesheet */
func ReadAndCompileStylesheet(templateName string, theme string, customStyle map[string]string) (string, error) {
	var (
		source     string
		themeProps map[string]string
		props      map[string]string = map[string]string{}
		err        error
	)

	// 1. Retrieve style template source code
	source, err = GetTemplateSource(templateName)

	if err != nil {
		logger.Err("Couldn't get style template", err)
		return "", err
	}

	// 2. Get theme properties
	themeProps, err = GetTemplateThemeProps(theme)

	if err != nil {
		logger.Err("Couldn't get theme props", err)
		return "", err
	}

	/* 3. assign values into `props` based on priority
		1. Custom style
	 	2. Theme style
	 	3. Default style */
	for prop, value := range config.DefaultStyleProps {
		if len(customStyle[prop]) > 0 {
			props[prop] = customStyle[prop]
		} else if len(themeProps[prop]) > 0 {
			props[prop] = themeProps[prop]
		} else {
			props[prop] = value
		}
	}

	return raymond.Render(source, props)
}

/* Read template and returns compiles version of markup */
func ReadAndCompileMarkup(layout string, data *models.UserData) (string, error) {
	var (
		temp string
		err  error
	)

	// read template source for body
	if temp, err = GetTemplateSource(config.TEMPLATE_DEFAULT_CONTENT); err != nil {
		logger.Err("Couldn't load content markup", err)
		return "", err
	}

	// compile body into plain markup
	if data.Body, err = raymond.Render(temp, data); err != nil {
		logger.Err("Couldn't render template body", err)
		return "", err
	}

	// get stylesheet code
	if data.Css, err = ReadAndCompileStylesheet(
		config.TEMPLATE_DEFAULT_STYLE,
		data.Page.Theme+"."+config.THEMES_EXT,
		data.Page.Style); err != nil {
		logger.Err("Couldn't compile stylesheet", err)
		return "", err
	}

	// compile layout by passing body and css
	if temp, err = GetTemplateSource(config.TEMPLATE_DEFAULT_LAYOUT); err != nil {
		logger.Err("Couldn't load layout markup", err)
		return "", err
	}

	return raymond.Render(temp, data)
}
