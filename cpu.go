package system

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"math"
	"runtime"
)

type CPU struct {
	Name         string `json:"name" validate:"required"`
	Threads      int    `json:"threads" validate:"required"`
	Architecture string `json:"architecture" validate:"required"`
}

type CpuSystemInformation struct {
	Value      string    `json:"value" validate:"required"`
	Percentage []float64 `json:"percentage" validate:"required"`
}

func staticCpu() CPU {
	var p CPU
	p.Threads = runtime.NumCPU()
	p.Architecture = runtime.GOARCH
	c, err := cpu.Info()
	if err == nil {
		p.Name = c[0].ModelName
	} else {
		p.Name = "none detected"
	}
	return p
}

func (s *System) liveCpu() error {
	p, err := cpu.Percent(0, false)
	if err != nil {
		return fmt.Errorf("cannot update live cpu")
	}
	s.Live.CPU.Value = s.Static.CPU.Name
	s.Live.CPU.Percentage = append(s.Live.CPU.Percentage[1:], math.RoundToEven(p[0]))
	return nil
}
