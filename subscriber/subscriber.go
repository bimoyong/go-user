package subscriber

import (
	"gitlab.com/bimoyong/go-util/config/database"
)

// Close function
func Close() error {
	return database.Close()
}
