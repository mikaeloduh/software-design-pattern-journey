package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPatient_AddCase(t *testing.T) {
	p := NewPatient("a0000001", "p1", Male, 87, 159, 100)
	p.AddCase(Case{CaseTime: time.Now()})

	assert.Len(t, p.Cases, 1)
}
