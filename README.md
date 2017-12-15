# drawing - a 2D drawing package

This is a simple Go package for drawing on a 2D surface. It includes primitives and PNG export, and will also support sprites and related higher-level constructs.

Currently contains:

* Create surface
* Clear surface, with transparency
* Plot point
* Horizontal and vertical line
* Non-aliased line with fast integer routines
* Anti-aliased line and plot
* Rectangle outline
* Filled rectangle
* Circle outline
* Filled circle
* Pre-calculated color gradient
* Weighted color on a scale
* Write as a PNG stream
* Some predefined colors
* Some functionality shows timings

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
