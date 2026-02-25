package builtins

import (
	methods "github.com/pgcamus/gonja/v2/builtins/methods"
	"github.com/pgcamus/gonja/v2/exec"
)

// ControlStructures exports all builtins controlStructures
var Methods exec.Methods = methods.All
