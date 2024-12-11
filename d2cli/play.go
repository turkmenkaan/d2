package d2cli

import (
	"context"
	"fmt"
	"os"

	"oss.terrastruct.com/d2/lib/urlenc"
	"oss.terrastruct.com/util-go/xbrowser"
	"oss.terrastruct.com/util-go/xmain"
)

func playSubcommand(ctx context.Context, ms *xmain.State) error {
	if len(ms.Opts.Flags.Args()) != 2 {
		return xmain.UsageErrorf("play must be passed one file to open")
	}
	filepath := ms.Opts.Flags.Args()[1]

	theme, err := ms.Opts.Flags.GetInt64("theme")
	if err != nil {
		return err
	}

	sketch, err := ms.Opts.Flags.GetBool("sketch")
	if err != nil {
		return err
	}

	var sketchNumber int
	if sketch {
		sketchNumber = 1
	} else {
		sketchNumber = 0
	}

	fileRaw, err := readFile(filepath)
	if err != nil {
		return err
	}

	encoded, err := urlenc.Encode(fileRaw)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://play.d2lang.com/?l=&script=%s&sketch=%d&theme=%d&", encoded, sketchNumber, theme)
	openBrowser(ctx, ms, url)
	return nil
}

func readFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", xmain.UsageErrorf(err.Error())
	}

	return string(data), nil
}

func openBrowser(ctx context.Context, ms *xmain.State, url string) {
	ms.Log.Info.Printf("opening playground: %s", url)

	err := xbrowser.Open(ctx, ms.Env, url)
	if err != nil {
		ms.Log.Warn.Printf("failed to open browser to %v: %v", url, err)
	}
}
