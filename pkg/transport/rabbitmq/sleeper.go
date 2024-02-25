package rabbitmq

import (
	"math"
	"time"
)

type SecondsSleeper struct {
	backoffBase        int
	backoffCoefficient int
}

func NewSecondsSleeper(backoffBase, backoffCoefficient int) *SecondsSleeper {
	return &SecondsSleeper{
		backoffBase:        backoffBase,
		backoffCoefficient: backoffCoefficient,
	}
}

func (s SecondsSleeper) calculatePause(tryCount int) int {
	return int(math.Pow(float64(s.backoffBase*s.backoffCoefficient), float64(tryCount)))
}

func (s SecondsSleeper) Sleep(retryCount int) {
	s.SleepFor(s.calculatePause(retryCount))
}

func (s SecondsSleeper) SleepFor(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
}
