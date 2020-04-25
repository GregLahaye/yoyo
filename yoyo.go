package yoyo

import (
	"fmt"
	"time"

	"github.com/GregLahaye/yogurt"
	"github.com/GregLahaye/yoyo/styles"
)

type Spinner struct {
	Style styles.Style
	quit  chan bool
	done  chan bool
}

func Start(style styles.Style) *Spinner {
	s := &Spinner{
		Style: style,
		quit:  make(chan bool),
		done:  make(chan bool),
	}
	go s.Spin()
	return s
}

func (s *Spinner) End() {
	s.quit <- true
	<-s.done
}

func (s *Spinner) Spin() {
	fmt.Print(yogurt.DisableCursor)
	defer fmt.Print(yogurt.EnableCursor)

	for {
		select {
		case <-s.quit:
			s.done <- true
			return
		default:
			for _, f := range s.Style.Frames {
				fmt.Print(f)
				time.Sleep(time.Millisecond * s.Style.Milliseconds)
				for i := 0; i < len(f); i++ {
					fmt.Print(yogurt.CursorBackward(1))
					fmt.Print(" ")
					fmt.Print(yogurt.CursorBackward(1))
				}
			}
		}
	}
}
