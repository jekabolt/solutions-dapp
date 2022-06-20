package bucket

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func pngFromB64(b64Image []byte) (image.Image, error) {
	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(b64Image))
	i, err := png.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("PNGFromB64:image.Decode: [%v]", err.Error())
	}
	return i, nil
}

func jpgFromB64(b64Image []byte) (image.Image, error) {
	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(b64Image))
	i, err := jpeg.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("JPGFromB64:image.Decode")
	}
	return i, nil
}

func encodeJPG(w io.Writer, img image.Image, quality int) error {
	var err error
	var rgba *image.RGBA
	if nrgba, ok := img.(*image.NRGBA); ok {
		if nrgba.Opaque() {
			rgba = &image.RGBA{
				Pix:    nrgba.Pix,
				Stride: nrgba.Stride,
				Rect:   nrgba.Rect,
			}
		}
	}

	if rgba != nil {
		err = jpeg.Encode(w, rgba, &jpeg.Options{Quality: quality})
	} else {
		err = jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}

	return err
}
