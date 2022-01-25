package image

import "testing"

func TestProcessLayers(t *testing.T) {
	imageHexHash := "e7d88de73db3"
	ProcessLayers(imageHexHash, imageHexHash)
}
