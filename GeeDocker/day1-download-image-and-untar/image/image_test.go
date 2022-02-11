package image

import (
	"log"
	"strings"
	"testing"
)

func TestProcessLayers(t *testing.T) {
	err := DownloadImageIfNessary(strings.Join([]string{"alpine", "latest"}, ":"))
	if err != nil {
		log.Println("DownloadImageIfNessary error")
	}
}
