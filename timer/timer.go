package timer

import (
	"embed"
	"fmt"
	"os"
	"time"

	"github.com/andybrewer/mack"
	"github.com/drgrib/ttimer/agent"
	"github.com/drgrib/ttimer/parse"
)

//go:embed assets/clock_red.icns
var assets embed.FS

func Start(duration time.Duration, title string) error {
	if len(title) < 1 {
		title = fmt.Sprintf("%v timer", duration)
	}
	timer := agent.Timer{
		Title:    title,
		AutoQuit: true,
	}
	timer.Start(duration)
	timer.CountDown()
	if err := notify(title); err != nil {
		return fmt.Errorf("failed to notify: %w", err)
	}
	return nil
}

func notify(title string) error {
	if len(title) < 1 {
		title = "timer"
	}
	// TODO: probably do this in cmd/timer.main() and pass the file name as a param
	// TODO: save clock_red.icns to tmp directory
	iconFile := "clock_red.icns"
	iconFilePath := fmt.Sprintf("assets/%s", iconFile)
	iconContents, err := assets.ReadFile(iconFilePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", iconFilePath, err)
	}
	if err := os.WriteFile(iconFile, iconContents, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", iconFile, err)
	}
	if _, err := mack.DialogBox(mack.DialogOptions{
		Title: title,
		Text:  "time's up!",
		Icon:  fmt.Sprintf("./%s", iconFile),
	}); err != nil {
		return fmt.Errorf("failed to display dialog: %w", err)
	}
	return nil
}

func ParseDuration(duration string) (time.Duration, error) {
	d, _, err := parse.Args(duration)
	if err != nil {
		return 0, fmt.Errorf("")
	}
	return d, nil
}
