package utils

import "github.com/fatih/color"

// Color and style variables
var (
	Green         = color.New(color.FgGreen)
	Blue          = color.New(color.FgBlue)
	Red           = color.New(color.FgRed)
	White         = color.New(color.FgWhite)
	Yellow        = color.New(color.FgYellow)
	Faint         = color.New(color.Faint)
	Bold          = color.New(color.Bold)
	Underline     = color.New(color.Underline)
	BoldUnderline = color.New(color.Bold, color.Underline)
	Success       = Green.Sprintf("✓")
	Warning       = Yellow.Sprintf("!")
	Error         = Red.Sprintf("X")
)
