package system

import (
	"fmt"
	"github.com/jaypipes/ghw/pkg/unitutil"
)

type Storage struct {
	Readable   string  `json:"readable" validate:"required"`
	Value      float64 `json:"value" validate:"required"`
	UnitString string  `json:"unit_string" validate:"required"`
	Unit       float64 `json:"unit" validate:"required"`
}

func processStorage(s *Storage, total uint64) {
	unit, unitStr := unitutil.AmountString(int64(total))
	s.Unit = float64(unit)
	s.Value = float64(total) / s.Unit
	s.UnitString = unitStr
	s.Readable = fmt.Sprintf("%.2f %s", s.Value, s.UnitString)
}
