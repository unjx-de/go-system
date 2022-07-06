package system

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

var s System

func TestSystem_uptime(t *testing.T) {
	s.uptime()
	tempUptime := s.Live.ServerUptime
	time.Sleep(1 * time.Second)
	s.uptime()
	difference := s.Live.ServerUptime - tempUptime
	assert.Equal(t, uint64(1000), difference)
}
