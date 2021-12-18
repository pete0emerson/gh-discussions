package cmd

import (
	"fmt"
	"time"
)

func pluralize(num int, thing string) string {
	if num == 1 {
		return fmt.Sprintf("%d %s", num, thing)
	}
	return fmt.Sprintf("%d %ss", num, thing)
}

func timeAgo(ago time.Duration) string {
	if ago < time.Minute {
		return "just now"
	}
	if ago < time.Hour {
		return pluralize(int(ago.Minutes()), "minute") + " ago"
	}
	if ago < 24*time.Hour {
		return pluralize(int(ago.Hours()), "hour") + " ago"
	}
	if ago < 30*24*time.Hour {
		return pluralize(int(ago.Hours())/24, "day") + " ago"
	}
	if ago < 365*24*time.Hour {
		return pluralize(int(ago.Hours())/24/30, "month") + " ago"
	}
	return pluralize(int(ago.Hours()/24/365), "year") + " ago"
}
