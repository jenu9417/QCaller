package uuid

import (
	"QCaller/logger"
	"os"
	"strings"

	"github.com/google/uuid"
)

// New : returns new uuid
func New() string {
	id, err := uuid.NewUUID()
	if err != nil {
		logger.Get().Fatalf("Error while creating uuid")
		os.Exit(1)
	}

	return strings.Replace(id.String(), "-", "", -1)
}
