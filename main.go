package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rajatxs/go-iamone/common"
	"github.com/rajatxs/go-iamone/db"
	"github.com/rajatxs/go-iamone/handlers"
	"github.com/rajatxs/go-iamone/logger"
)

/* Starts HTTP Server on given port */
func bootServer() (err error) {
	var (
		port string = os.Getenv("PORT")
		fs   http.Handler
	)

	fs = http.FileServer(http.Dir("public"))
	http.Handle("/", handlers.RootNavigator(fs))
	logger.Info(fmt.Sprintf("Server starting on port %s", port))

	return http.ListenAndServe(port, nil)
}

func main() {
	common.Ensure(db.Connect(), "Failed to make database connection")
	common.Ensure(bootServer(), "Failed to boot server")
}
