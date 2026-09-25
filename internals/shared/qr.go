package shared

import (
	"os"

	"github.com/mdp/qrterminal/v4"
)

func GenerateQr(url string) {
	config := qrterminal.Config{
		Level:      qrterminal.M,
		Writer:     os.Stdout,
		HalfBlocks: true,
		QuietZone:  1,
	}

	qrterminal.GenerateWithConfig(url, config)
}
