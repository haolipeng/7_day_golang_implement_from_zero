package image

import (
	"log"
	"strings"
	"testing"
)

func TestDownloadImageIfNessary(t *testing.T) {
	err := DownloadImageIfNessary(strings.Join([]string{"alpine", "latest"}, ":"))
	if err != nil {
		log.Println("DownloadImageIfNessary error")
	}
}
