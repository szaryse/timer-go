package ui

import "fmt"

func calcRowY(idx int) int {
	return py + (idx*fontSize + (idx * my * 2))
}

func formatTime(time int) string {
	if time < 0 {
		time *= -1
	}
	minutes := time / 60
	seconds := time - minutes*60
	minutesStr := fmt.Sprintf("%2d", minutes)
	secondsStr := fmt.Sprintf("%02d", seconds)
	return fmt.Sprintf("%s:%s", minutesStr, secondsStr)
}

func formatFullTime(time int) string {
	if time < 0 {
		time *= -1
	}
	hours := time / 3600
	minutes := (time - hours*3600) / 60
	seconds := time - hours*3600 - minutes*60
	hoursStr := fmt.Sprintf("%d", hours)
	minutesStr := fmt.Sprintf("%02d", minutes)
	secondsStr := fmt.Sprintf("%02d", seconds)
	return fmt.Sprintf("%s:%s:%s", hoursStr, minutesStr, secondsStr)
}
