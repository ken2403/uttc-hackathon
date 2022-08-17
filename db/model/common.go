package model

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

var idGenerator *ulidGenerator

func init() {
	idGenerator = newIdGenerator()
}

// ulidGenerator - ULIDに基づいた採番君
type ulidGenerator struct {
	entropy *ulid.MonotonicEntropy
}

func newIdGenerator() *ulidGenerator {
	return &ulidGenerator{
		entropy: ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0),
	}
}

func (receiver *ulidGenerator) Run() string {
	return receiver.RunWithTime(time.Now())
}

func (receiver *ulidGenerator) RunWithTime(t time.Time) string {
	return ulid.MustNew(ulid.Timestamp(t), receiver.entropy).String()
}
