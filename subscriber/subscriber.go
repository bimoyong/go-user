package subscriber

import (
	"github.com/bimoyong/go-util/config/database"
)

// Close function
func Close() error {
	return database.Close()
}
