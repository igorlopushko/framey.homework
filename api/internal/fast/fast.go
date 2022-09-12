package fast

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// Fast represents measurement structure
type Fast struct {
	Up       float64
	Down     float64
	UpUnit   string
	DownUnit string
}

// Measure does the main job.
// It returns *Fast and error
func Measure() (*Fast, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 180*time.Second)
	defer cancel()

	var upStr, downStr string

	fast := new(Fast)
	cmds := []chromedp.Action{
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &downStr, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#speed-units.succeeded`, &fast.DownUnit, chromedp.NodeVisible, chromedp.ByQuery),
	}

	cmds = append(cmds, chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#upload-value.succeeded`, &upStr, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#upload-units.succeeded`, &fast.UpUnit, chromedp.NodeVisible, chromedp.ByQuery),
	)

	err := chromedp.Run(ctx, cmds...)

	fast.Up, err = strconv.ParseFloat(upStr, 32)
	if err != nil {
		return fast, err
	}
	fast.Down, err = strconv.ParseFloat(downStr, 32)
	if err != nil {
		return fast, err
	}

	return fast, err
}

// Run is the ready to use API.
// For customization call Measure().
func Run() (Fast, error) {
	fast, err := Measure()
	if err != nil {
		return *fast, err
	}

	return *fast, nil
}
