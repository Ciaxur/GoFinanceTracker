package Utils

import "github.com/fatih/color"

// Output Strucutre for Color Scheme Prints
type Output struct {
	Error     *color.Color
	Warning   *color.Color
	Info      *color.Color
	Prompt    *color.Color
	Important *color.Color
}

/**
 * Creates an Output Object with Pre-Set
 *  color scheme for use
 */
func InitOutput() *Output {
	return &Output{
		Error:     color.New(color.FgRed).Add(color.Bold),
		Warning:   color.New(color.FgHiMagenta).Add(color.Bold),
		Info:      color.New(color.FgHiGreen).Add(color.Bold),
		Prompt:    color.New(color.FgHiCyan).Add(color.Bold),
		Important: color.New(color.FgHiMagenta),
	}
}

// GLOBAL VARIABLE FOR PRINTS
var Out *Output = InitOutput()
