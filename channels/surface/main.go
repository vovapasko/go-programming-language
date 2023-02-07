// Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3 or the 3-D surface
// computation of Section 3.2, and execute its main loop in parallel using channels for communication. How much faster
// does it run on a multiprocessor machine? What is the optimal number of goroutines to use?

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type Coordinate struct {
	X, Y float64
}

func SVG2(w io.Writer) {
	zmin, zmax := minmax()
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			aCord := corner(i+1, j)
			bCord := corner(i, j)
			cCord := corner(i, j+1)
			dCord := corner(i+1, j+1)
			if math.IsNaN(aCord.X) || math.IsNaN(aCord.Y) ||
				math.IsNaN(bCord.X) || math.IsNaN(bCord.Y) ||
				math.IsNaN(cCord.X) || math.IsNaN(cCord.Y) ||
				math.IsNaN(dCord.X) || math.IsNaN(dCord.Y) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, zmin, zmax), aCord.X, aCord.Y, bCord.X, bCord.Y, cCord.X, cCord.Y, dCord.X, dCord.Y)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func SVG(w io.Writer) {
	zmin, zmax := minmax()
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			aCord := corner(i+1, j)
			bCord := corner(i, j)
			cCord := corner(i, j+1)
			dCord := corner(i+1, j+1)
			if math.IsNaN(aCord.X) || math.IsNaN(aCord.Y) ||
				math.IsNaN(bCord.X) || math.IsNaN(bCord.Y) ||
				math.IsNaN(cCord.X) || math.IsNaN(cCord.Y) ||
				math.IsNaN(dCord.X) || math.IsNaN(dCord.Y) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, zmin, zmax), aCord.X, aCord.Y, bCord.X, bCord.Y, cCord.X, cCord.Y, dCord.X, dCord.Y)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func main() {
	//SVG(os.Stdout)
	file, err := os.Create("channels/surface/figure.svg")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	if err != nil {
		panic(err)
	}
	SVG2(file)
}

// minmax returns the min and max values for z given the min/max values of x
// and y and assuming a square domain.
func minmax() (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return
}

func color(i, j int, zmin, zmax float64) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}

func corner(i, j int) Coordinate {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return Coordinate{X: sx, Y: sy}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
