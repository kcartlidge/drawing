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
* Write as a PNG stream
* Some predefined colors
* Major routines show timings

The anti-aliased line drawing needs a little more work to look better on all mixes
of light and dark colors and backgrounds. This is in progress, but in the meantime
works fine on dark backgrounds.

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
