package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"

	pigo "github.com/esimov/pigo/core"
	"github.com/fogleman/gg"
)

var qThresh float32 = 5.0

type FaceMarker struct {
}

func NewFaceMarker() FaceMarker {
	fm := FaceMarker{}
	return fm
}

// drawMarker mark the detected face region with the provided
// marker (rectangle or circle) and write it to io.Writer.
func (fm FaceMarker) drawMarkerJPG(w io.Writer, image *image.NRGBA, detections []pigo.Detection, quality int) error {
	cols, rows := GetImageColsAndRows(image)

	dc := gg.NewContext(cols, rows)
	dc.DrawImage(image, 0, 0)

	for i := 0; i < len(detections); i++ {
		if detections[i].Q > qThresh {
			dc.DrawRectangle(
				float64(detections[i].Col-detections[i].Scale/2),
				float64(detections[i].Row-detections[i].Scale/2),
				float64(detections[i].Scale),
				float64(detections[i].Scale),
			)

			dc.SetLineWidth(3.0)
			dc.SetStrokeStyle(gg.NewSolidPattern(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
			dc.Stroke()
		}
	}

	var opt jpeg.Options
	opt.Quality = quality

	return jpeg.Encode(w, dc.Image(), &opt)
}
