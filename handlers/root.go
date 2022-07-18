package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/rajatxs/go-iamone/config"
	"github.com/rajatxs/go-iamone/logger"
	"github.com/rajatxs/go-iamone/models"
	"github.com/rajatxs/go-iamone/services"
)

/* Root middleware which responsible to target action based on url path */
func RootNavigator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			upath string = r.URL.Path
			uname string
		)

		/*
			each request url will be consider as static resource
			target if it contains "." or equal to "/"
		*/
		if upath == "/" || strings.Contains(upath, ".") {
			h.ServeHTTP(w, r)
		} else {
			// extract uname from url path
			uname = strings.TrimSpace(upath)[1:]
			SendProfilePage(w, r, uname)
		}
	})
}

/* Sends profile page by given username */
func SendProfilePage(w http.ResponseWriter, r *http.Request, uname string) {
	var (
		udata  models.UserData
		source string
		err    error
	)

	if err = services.ReadUserDataByUsername(uname, &udata); err != nil {
		logger.Err("Couldn't get user data", err)
		Send404Response(w, r)
		return
	}

	if err = services.ResolveSocialHref(&udata); err != nil {
		logger.Err("Couldn't resolve social href", err)
		Send500Response(w, r)
		return
	}

	if source, err = services.ReadAndCompileMarkup(config.TEMPLATE_DEFAULT_CONTENT, &udata); err != nil {
		logger.Err("Couldn't compile specified template", err)
		Send500Response(w, r)
		return
	}

	io.WriteString(w, source)
}

/* Sends 404 response page */
func Send404Response(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	http.ServeFile(w, r, config.PAGE_404)
}

/* Sends 500 response page */
func Send500Response(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	http.ServeFile(w, r, config.PAGE_500)
}
