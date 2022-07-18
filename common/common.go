package common

import (
	"encoding/json"
	"os"

	"github.com/rajatxs/go-iamone/logger"
)

func Ensure(err error, errMsg string) {
	if err != nil {
		logger.Err(errMsg, err)
	}
}

/* Reads content from given file as JSON */
func ReadJSON(fpath string, res interface{}) (err error) {
	var fsource []byte

	fsource, err = os.ReadFile(fpath)

	if err != nil {
		return err
	}

	return json.Unmarshal(fsource, res)
}
