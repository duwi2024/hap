// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/LUJUNQUAN/hap/characteristic"
)

type WindowCoveringVertical struct {
	*S

	CurrentPosition          *characteristic.CurrentPosition
	TargetPosition           *characteristic.TargetPosition
	PositionState            *characteristic.PositionState
	CurrentVerticalTiltAngle *characteristic.CurrentVerticalTiltAngle
	TargetVerticalTiltAngle  *characteristic.TargetVerticalTiltAngle
}

func NewWindowCoveringVertical() *WindowCoveringVertical {
	s := WindowCoveringVertical{}
	s.S = New(TypeWindowCovering)

	s.CurrentPosition = characteristic.NewCurrentPosition()
	s.AddC(s.CurrentPosition.C)

	s.TargetPosition = characteristic.NewTargetPosition()
	s.AddC(s.TargetPosition.C)

	s.PositionState = characteristic.NewPositionState()
	s.AddC(s.PositionState.C)

	s.CurrentVerticalTiltAngle = characteristic.NewCurrentVerticalTiltAngle()
	s.AddC(s.CurrentVerticalTiltAngle.C)

	s.TargetVerticalTiltAngle = characteristic.NewTargetVerticalTiltAngle()
	s.AddC(s.TargetVerticalTiltAngle.C)

	return &s
}
