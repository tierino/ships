package mapper

import (
	"example.com/types"
)

func Transform(msg *types.AISMessage) (string, error) {
	return "Transformed message", nil
}
