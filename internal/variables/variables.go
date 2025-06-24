package variables

import (
	"rest_io_bound/internal/models"
	"time"
)

const (
	HTTP_PORT_KEY     = "HTTP_PORT"
	MSG               = "Not successful"
	CONFIG_DIRECTORY  = ".config/riob"
	CONFIG_FILE_NAME  = "config"
	CONFIG_FILE_TYPE  = "yaml"
	DEFAULT_HTTP_PORT = "8080"
)

var Storage = map[string]*models.Task{
	"1": {
		Title:   "1",
		Status:  "Always In Progress",
		Created: time.Now(),
	},
}
