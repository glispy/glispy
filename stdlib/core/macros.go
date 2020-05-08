package core

import (
	"github.com/glispy/glispy/common"
	"github.com/glispy/glispy/types"
)

// ToQuoteMacro convert a quote macro to use the quote func
func ToQuoteMacro(sc types.Scope, args types.List) (out types.Expression, err error) {
	if len(args) < 1 {
		err = common.ErrInvalidArgs
		return
	}

	out = args[0]
	return
}
