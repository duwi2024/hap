package accessory

import (
	"github.com/LUJUNQUAN/hap/service"
)

type WindowCoveringHorizontal struct {
	*A
	WindowCoveringHorizontal *service.WindowCoveringHorizontal
}

// NewWindowCoveringHorizontal returns a window accessory.
func NewWindowCoveringHorizontal(info Info) *WindowCoveringHorizontal {
	a := WindowCoveringHorizontal{}
	a.A = New(info, TypeWindowCovering)

	a.WindowCoveringHorizontal = service.NewWindowCoveringHorizontal()
	a.AddS(a.WindowCoveringHorizontal.S)

	return &a
}
