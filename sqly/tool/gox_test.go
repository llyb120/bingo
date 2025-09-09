package tool

import (
	"testing"

	"github.com/llyb120/gox"
)

func TestGox(t *testing.T) {
	compiler := gox.Compiler{
		SrcPath:         "../../test/sql/src",
		DestPath:        "../../test/sql/dest",
		DebugMode:       true,
		RemoveGenerated: true,
	}
	compiler.Compile()
}
