package uuidgen

import (
	"strings"

	"github.com/google/uuid"
)

func NewUndashed() string {
	newUuid := uuid.New().String()
	return strings.Replace(newUuid, "-", "", -1)
}
