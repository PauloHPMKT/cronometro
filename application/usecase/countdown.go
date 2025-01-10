// Essa função poderá sser utilizada no caso de uso do controller
package countdown

import (
	"fmt"
	"time"
)

type CountdownTimer struct {
	seconds int
}

func NewCountdownTimer(minutes int) *CountdownTimer {
	return &CountdownTimer{
		seconds: minutes * 60,
	}
}

func (t *CountdownTimer) Start() {
	for t.seconds > 0 {
		min := t.seconds / 60
		sec := t.seconds % 60
		fmt.Printf("\rTempo restante: %02d:%02d", min, sec)

		time.Sleep(1 * time.Second)

		t.seconds--
		fmt.Printf("\rTempo restante: %02d:%02d", min, sec)
	}

	fmt.Println("\rTempo esgotado!")
}
