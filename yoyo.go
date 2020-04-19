package yoyo

import (
	"fmt"
	"github.com/GregLahaye/yogurt"
	"github.com/GregLahaye/yoyo/styles"
	"time"
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
	for !s.Condition {
		for _, c := range s.Style.Characters {
			fmt.Print(c)
			time.Sleep(time.Millisecond * s.Style.Milliseconds)
			yogurt.CursorBackward(len(c))
		}
	}

	s.Channel <- true
}

func (s *Spinner) End() {
	s.Condition = true
	<-s.Channel
}
