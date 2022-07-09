package system

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSystem_staticCPU(t *testing.T) {
	var s System
	emptyStatic := StaticInformation{}
	s.Static.CPU = staticCpu()
	assert.NotEqual(t, s.Static.CPU, emptyStatic.CPU)
}

func TestSystem_liveCPU(t *testing.T) {
	var s System
	emptyLive := LiveInformation{}
	emptyLive.CPU.Percentage = make([]float64, 60)
	s.Live.CPU.Percentage = make([]float64, 60)
	err := s.liveCpu()
	if err != nil {
		assert.Equal(t, s.Live.CPU, emptyLive.CPU)
	} else {
		assert.NotEqual(t, s.Live.CPU, emptyLive.CPU)
	}
}
