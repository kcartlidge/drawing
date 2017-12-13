# drawing - a 2D drawing package

This is a simple Go package for drawing on a 2D surface.
It will include primitives, sprites, higher-level constructs, and PNG export.

Currently contains:

* Create surfaces
* Clear surfaces
* Supports transparency
* Plot points
* Draw horizontal and vertical lines
* Draw non-aliased lines with fast integer routines
* Draw anti-aliased lines (Xiaolin Wu's algorithm)
* Fill rectangles
* Draw rectangle outlines
* Write as a PNG stream
* Some predefined colors

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
