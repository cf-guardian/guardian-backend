package identity

import (
	api "github.com/cf-guardian/guardian-backend/temp_libcontainer_api"
	"github.com/cf-guardian/guardian-backend/utils"
	"strings"
	"testing"
)

func init() {
	utils.OptimiseScheduling()
}

// Test the SimpleIdentifier implementation of Identifier

func TestSICreate(t *testing.T) {
	ider := CreateSimpleIdentifier()
	id := ider.Generate()
	if sid, snm := string(id), string(ider.Name(id)); sid != snm {
		t.Fatalf("id not converted verbatim; id=%q, Name(id)=%q", sid, snm)
	}
}

const (
	LARGEISH_NUMBER int = 1000
)

func TestSIForm(t *testing.T) {
	ider := CreateSimpleIdentifier()

	id := ider.Generate()

	sid := string(id)
	if len(sid) != idLength {
		t.Fatalf("Identifiers should be exactly %d characters long, was %q.", idLength, sid)
	}
	if !strings.HasPrefix(sid, idPrefix) {
		t.Fatalf("Identifiers should begin with %q.", idPrefix)
	}

	digsId := strings.TrimPrefix(sid, idPrefix)
	if str := strings.Trim(digsId, "0123456789"); "" != str {
		t.Fatalf("Identifier contains non-digits after prefix: %q.", str)
	}
}

func TestSIMulti(t *testing.T) {
	var genArr [LARGEISH_NUMBER]api.Id

	ider := CreateSimpleIdentifier()
	for i := 0; i < LARGEISH_NUMBER; i++ {
		genArr[i] = ider.Generate()
	}

	// Now check they are all different
	set := make(map[api.Id]struct{})
	for i := 0; i < LARGEISH_NUMBER; i++ {
		set[genArr[i]] = struct{}{}
	}

	if l := len(set); l < LARGEISH_NUMBER {
		t.Fatalf("Non-unique identifiers: %d distinct ids returned in first %d.", l, LARGEISH_NUMBER)
	}
}

func TestSIClash(t *testing.T) {
	var genArr [LARGEISH_NUMBER]api.Id

	ider1 := CreateSimpleIdentifier()
	ider2 := CreateSimpleIdentifier()

	for i := 0; i < LARGEISH_NUMBER; i++ {
		genArr[i] = ider1.Generate()
	}

	set := make(map[api.Id]struct{})
	for i := 0; i < LARGEISH_NUMBER; i++ {
		set[genArr[i]] = struct{}{}
	}

	// Now check the ones from ider2 don't clash
	for i := 0; i < LARGEISH_NUMBER; i++ {
		gid := ider2.Generate()
		if _, has := set[gid]; has {
			t.Fatalf("Identifiers clashed: %q generated before!", gid)
		}
	}

	if l := len(set); l < 2*LARGEISH_NUMBER {
	}
}

func TestSIThreadSafe(t *testing.T) {
	// Should be run with GOMAXPROCS>1 to detect thread-unsafeness.
	ider := CreateSimpleIdentifier()
	done1 := make(chan []api.Id)
	done2 := make(chan []api.Id)

	go callSIGenerate(ider, done1)
	go callSIGenerate(ider, done2)

	genArr := <-done1

	set := make(map[api.Id]struct{})
	for i := 0; i < LARGEISH_NUMBER; i++ {
		set[genArr[i]] = struct{}{}
	}

	// Now check the ones from done2 don't clash
	for _, gid := range <-done2 {
		if _, has := set[gid]; has {
			t.Fatalf("Identifiers clashed: %q generated before!", gid)
		}
	}
}

func callSIGenerate(ider Identifier, result chan []api.Id) {
	var genArr []api.Id = make([]api.Id, LARGEISH_NUMBER)

	for i := 0; i < LARGEISH_NUMBER; i++ {
		genArr[i] = ider.Generate()
	}
	result <- genArr
}
