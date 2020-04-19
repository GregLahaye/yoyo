package styles

import "time"

type Style struct {
	Milliseconds time.Duration
	Characters   []string
}

var Simple = Style{200, []string{".  ", ".. ", "...", " ..", "  .", "   "}}
var Star = Style{80, []string{"+", "x", "*"}}
var Point = Style{125, []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}}
var Balloon = Style{140, []string{" ", ".", "o", "O", "@", "*", " "}}
var Bounce = Style{80, []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"}}
var Bar = Style{120, []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"}}
var Box = Style{100, []string{"▌", "▀", "▐", "▄"}}
var Noise = Style{100, []string{"▓", "▒", "░"}}
