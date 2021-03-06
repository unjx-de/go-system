package system

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSystem_staticRam(t *testing.T) {
	var s System
	emptyStatic := StaticInformation{}
	s.Static.Ram = staticRam()
	assert.NotEqual(t, s.Static.Ram, emptyStatic.Ram)
}

func TestSystem_liveRam(t *testing.T) {
	var s System
	emptyLive := LiveInformation{}
	err := s.liveRam()
	if err != nil {
		assert.Equal(t, s.Live.Ram, emptyLive.Ram)
	} else {
		assert.NotEqual(t, s.Live.Ram, emptyLive.Ram)
	}
}
