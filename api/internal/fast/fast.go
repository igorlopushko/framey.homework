// Package fast copied from the https://github.com/adhocore/fast and simplified to fulfill the task requirements.
package fast

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// A Fast represents measurement structure.
type Fast struct {
	Up       float64
	Down     float64
	UpUnit   string
	DownUnit string
}

// Runs speed test check.
func Run() (*Fast, error) {
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
