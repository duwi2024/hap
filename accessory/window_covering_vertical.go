package accessory

import (
	"github.com/LUJUNQUAN/hap/service"
)

type WindowCoveringVertical struct {
	*A
	WindowCoveringVertical *service.WindowCoveringVertical
}

// NewWindowCoveringVertical returns a window accessory.
func NewWindowCoveringVertical(info Info) *WindowCoveringVertical {
	a := WindowCoveringVertical{}
	a.A = New(info, TypeWindowCovering)

	a.WindowCoveringVertical = service.NewWindowCoveringVertical()
	a.AddS(a.WindowCoveringVertical.S)

	return &a
}
