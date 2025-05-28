// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/LUJUNQUAN/hap/characteristic"
)

type WindowCoveringHorizontal struct {
	*S

	CurrentPosition            *characteristic.CurrentPosition
	TargetPosition             *characteristic.TargetPosition
	PositionState              *characteristic.PositionState
	CurrentHorizontalTiltAngle *characteristic.CurrentHorizontalTiltAngle
	TargetHorizontalTiltAngle  *characteristic.TargetHorizontalTiltAngle
}

func NewWindowCoveringHorizontal() *WindowCoveringHorizontal {
	s := WindowCoveringHorizontal{}
	s.S = New(TypeWindowCovering)

	s.CurrentPosition = characteristic.NewCurrentPosition()
	s.AddC(s.CurrentPosition.C)

	s.TargetPosition = characteristic.NewTargetPosition()
	s.AddC(s.TargetPosition.C)

	s.PositionState = characteristic.NewPositionState()
	s.AddC(s.PositionState.C)

	s.CurrentHorizontalTiltAngle = characteristic.NewCurrentHorizontalTiltAngle()
	s.AddC(s.CurrentHorizontalTiltAngle.C)

	s.TargetHorizontalTiltAngle = characteristic.NewTargetHorizontalTiltAngle()
	s.AddC(s.TargetHorizontalTiltAngle.C)

	return &s
}
