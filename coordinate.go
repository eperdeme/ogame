package ogame

import (
	"fmt"
	"strings"
)

// Coordinate represent an ogame coordinate
type Coordinate struct {
	Galaxy   int
	System   int
	Position int
	Type     CelestialType
}

func (c Coordinate) String() string {
	return fmt.Sprintf("[%c:%d:%d:%d]", strings.ToUpper(c.Type.String())[0], c.Galaxy, c.System, c.Position)
}

// Equal returns either two coordinates are equal or not
func (c Coordinate) Equal(v Coordinate) bool {
	return c.Galaxy == v.Galaxy &&
		c.System == v.System &&
		c.Position == v.Position &&
		c.Type == v.Type
}
