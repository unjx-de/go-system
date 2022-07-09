package system

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSystem_staticDisk(t *testing.T) {
	var s System
	emptyStatic := StaticInformation{}
	s.Static.Disk = staticDisk()
	assert.NotEqual(t, s.Static.Disk, emptyStatic.Disk)
}

func TestSystem_liveDisk(t *testing.T) {
	var s System
	emptyLive := LiveInformation{}
	err := s.liveDisk()
	if err != nil {
		assert.Equal(t, s.Live.Disk, emptyLive.Disk)
	} else {
		assert.NotEqual(t, s.Live.Disk, emptyLive.Disk)
	}
}
