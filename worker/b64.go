package worker

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"github.com/vincent-petithory/dataurl"
)

func ReadImageB64DataUrl(url string, w io.Writer) error {
	dataURL, err := dataurl.DecodeString(url)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(dataURL.Data))
	if err != nil {
		return err
	}

	switch dataURL.MediaType.ContentType() {
	case "image/png":
		err = png.Encode(w, img)
		// Add cases for other image formats if necessary
	default:
		return fmt.Errorf("unsupported image format: %s", dataURL.MediaType.ContentType())
	}

	return err
}

func SaveImageB64DataUrl(url, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return ReadImageB64DataUrl(url, file)
}
