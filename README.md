# drawing - a 2D drawing package

This is a simple Go package for drawing on a 2D surface. It includes primitives and PNG export, and will also support sprites and related higher-level constructs.

Currently contains:

* Create surfaces
* Clear surfaces
* Supports transparency
* Plot points
* Horizontal and vertical lines
* Non-aliased lines with fast integer routines
* Anti-aliased lines (Xiaolin Wu's algorithm)
* Rectangle outlines
* Filled rectangles
* Circle outlines
* Filled circles
* Pre-calculated color gradients
* Write as a PNG stream
* Some predefined colors
* Major routines show timings

The anti-aliased line drawing needs a little more work to look better on all mixes
of light and dark colors and backgrounds. This is in progress, but in the meantime
works fine on dark backgrounds.

The underlying cause is the use of the pixel weighting (derived from the deviation
from the line) being used to reduce the RGB values of the plot color.
This is fine on dark backgrounds but exactly the opposite of what is needed for ligher
ones. The solution in progress is to use the weighting to determine how far the plot
color needs to move along the color gradient formed by the plot color and the existing
point color on the surface. In other words asking should it be closer to the plot color
or closer to the existing background, rather than an absolute scale.

## Instructions

In progress.

## Building the example for all platforms

Substitute the script file according to your own platform:

``` sh
cd example
./scripts/macos.sh
```

## Creating and running a one-off build of the example

Substitute the file name according to your own platform:

``` sh
cd example
go build -o ./builds/macos/example && ./builds/macos/example
```
