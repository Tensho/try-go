package main

import (
	"fmt"
	"math"
	"log"
	"net/http"
	"io"
	"strconv"
)

const (
	cells         = 100         // number of grid cells
	xyrange       = 30.0        // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6 // angle of x, y axes (=30°)
)

var color string
var width, height int                               // canvas size in pixels
var xyscale, zscale float64                         // pixels per x or y unit, pixels per z unit
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		width, _ = strconv.Atoi(r.URL.Query().Get("width"))
		if width == 0 {
			width = 600
		}

		height, _ = strconv.Atoi(r.URL.Query().Get("height"))
		if height == 0 {
			height = 320
		}

		color = r.URL.Query().Get("color")
		if color == "" {
			color = "black"
		}

		xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
		zscale  = float64(height) * 0.4        // pixels per z unit

		w.Header().Set("Content-Type", "image/svg+xml")

		surface(w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Fprintf(out, "<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
