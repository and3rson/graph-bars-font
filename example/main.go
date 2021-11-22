package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

const HistorySize = 16

func main() {
	history := [HistorySize]int{}
	for {
		loads, err := cpu.Percent(0, false)
		if err == nil {
			load := int(loads[0])
			for i := range history {
				if i < HistorySize - 1 {
					history[i] = history[i + 1]
				} else {
					history[i] = load
				}
			}

			bar := ""
			for i := 0; i < HistorySize; i++ {
				load := int(float64(history[i]) / 100 * 8) + 1
				if load > 8 {
					load = 8
				}
				bar += string(rune(0x3000 + load))
			}

			fmt.Printf("CPU: <span font_desc=\"Bars\">%s</span> %2d%%\n", bar, load)
		}

		<-time.After(time.Second / 4.0)
	}
}
