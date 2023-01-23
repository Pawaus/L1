package L1_25

import "time"

func TwentyFive(timeSec int) {
	ticker := time.NewTicker(time.Second)
	i := 0
	for {
		select {
		case <-ticker.C:
			i++
			//wait for timeSec seconds
			if i == timeSec {
				return
			}
		}
	}

}
