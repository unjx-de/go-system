package system

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestSystem_uptime(t *testing.T) {
	var s System
	s.uptime()
	tempUptime := s.Live.ServerUptime
	time.Sleep(1 * time.Second)
	s.uptime()
	difference := s.Live.ServerUptime - tempUptime
	assert.Equal(t, uint64(1000), difference)
}
