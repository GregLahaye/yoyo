package styles

import "time"

type Style struct {
	Milliseconds time.Duration
	Frames       []string
}

var YoYo = Style{80, []string{"@    ", "~@   ", "~~@  ", "~~~@ ", "~~~~@", "~~~~@", "~~~~@", "~~~~@", "~~~~@", "~~@  ", "@    ", "@    ", "@    ", "@    ", "@    ", "@    ", "@    ", "@    ", "@    ", "@    "}}
var Simple = Style{200, []string{".  ", ".. ", "...", " ..", "  .", "   "}}
var Star = Style{80, []string{"+", "x", "*"}}
var Point = Style{125, []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}}
var Balloon = Style{140, []string{" ", ".", "o", "O", "@", "*", " "}}
var Bounce = Style{80, []string{"[ I    ]", "[  I   ]", "[   I  ]", "[    I ]", "[     I]", "[    I ]", "[   I  ]", "[  I   ]", "[ I    ]", "[I     ]"}}
var Bar = Style{120, []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"}}
var Box = Style{100, []string{"▌", "▀", "▐", "▄"}}
var Noise = Style{100, []string{"▓", "▒", "░"}}
