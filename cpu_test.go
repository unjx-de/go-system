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
