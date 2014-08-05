package identity

import (
	"fmt"
	api "github.com/cf-guardian/guardian-backend/temp_libcontainer_api"
	rand "math/rand"
	"sync/atomic"
	"time"
)

type simpleIdentifier struct {
	seed int64
}

// implementation parameters
const (
	idPrefix string = "SI"
	idLength int    = 22
	idFormat string = "SI%020d" // consistent with idPrefix and idLength
)

func CreateSimpleIdentifier() Identifier {
	rand.Seed(time.Now().UTC().UnixNano())

	return &simpleIdentifier{seed: rand.Int63()}
}

func (this *simpleIdentifier) Generate() api.Id {
	// Next line makes this thread-safe
	next := atomic.AddInt64(&this.seed, 1) // increment this.seed, atomically
	return api.Id(fmt.Sprintf(idFormat, next))
}

func (*simpleIdentifier) Name(id api.Id) api.Name {
	return api.Name(id)
}
