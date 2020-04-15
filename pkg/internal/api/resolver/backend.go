package resolver

import (
	"github.com/parrotmac/littleblue/pkg/internal/client/prisma"
)

type backend struct {
	prisma          *prisma.Client
	queryBackend    queryBackend
	mutationBackend mutationBackend
}
