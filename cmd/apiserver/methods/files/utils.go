package files

import (
	"fmt"
	"github.com/oklog/ulid"
	"math/rand"
	"strings"
	"time"
)

const APIUrl = "http://localhost:9090/api/upload/"
const DirF = "./cmd/apiserver/upload/"

//случайный uid
func randomFileName() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return strings.ToLower(fmt.Sprintf("%v", ulid.MustNew(ulid.Timestamp(t), entropy)))
}
