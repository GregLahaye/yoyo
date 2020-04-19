package yoyo

import (
	"fmt"
	"time"

	"github.com/GregLahaye/yogurt"
	"github.com/GregLahaye/yoyo/styles"
)

type Spinner struct {
	Style     styles.Style
	Condition bool
	Channel   chan bool
}

func Start(style styles.Style) *Spinner {
	s := &Spinner{
		Style:     style,
		Condition: false,
		Channel:   make(chan bool),
	}
	go s.Spin()
	return s
}

func (s *Spinner) Spin() {
	fmt.Print(yogurt.DisableCursor)

	for !s.Condition {
		for _, c := range s.Style.Characters {
			fmt.Print(c)
			time.Sleep(time.Millisecond * s.Style.Milliseconds)
			yogurt.CursorBackward(len(c))
		}
	}

	fmt.Print(yogurt.EnableCursor)

	s.Channel <- true
}

func (s *Spinner) End() {
	s.Condition = true
	<-s.Channel
}
