package sleep

import "time"
import "testing"

func TestSleep(t *testing.T) {
	start := time.Now()
	sleep(3)
	end := time.Now()

	if end.Sub(start) < 3 {
		t.Error("Awakes before 3s")
	}
}
