package system

import (
	"time"
)

var LiveInformationCh chan LiveInformation

type BasicSystemInformation struct {
	Value      string  `json:"value" validate:"required"`
	Percentage float64 `json:"percentage" validate:"required"`
}

type LiveInformation struct {
	CPU          CpuSystemInformation   `json:"cpu" validate:"required"`
	Ram          BasicSystemInformation `json:"ram" validate:"required"`
	Disk         BasicSystemInformation `json:"disk" validate:"required"`
	ServerUptime uint64                 `json:"server_uptime" validate:"required"`
}

type StaticInformation struct {
	CPU  CPU     `json:"cpu" validate:"required"`
	Ram  Storage `json:"ram" validate:"required"`
	Disk Storage `json:"disk" validate:"required"`
}

type System struct {
	Live   LiveInformation   `json:"live" validate:"required"`
	Static StaticInformation `json:"static" validate:"required"`
}

func (s *System) UpdateLiveInformation() {
	for {
		_ = s.liveCpu()
		_ = s.liveRam()
		_ = s.liveDisk()
		s.uptime()
		LiveInformationCh <- s.Live
		time.Sleep(1 * time.Second)
	}
}

func (s *System) Initialize() {
	s.Static.CPU = staticCpu()
	s.Static.Ram = staticRam()
	s.Static.Disk = staticDisk()
	s.Live.CPU.Percentage = make([]float64, 60)
	LiveInformationCh = make(chan LiveInformation)
	go s.UpdateLiveInformation()
}
