# drawing - a 2D drawing package

This is a simple Go package for drawing on a 2D surface.
It will include primitives, sprites, higher-level constructs, and PNG export.

Currently contains:

* Create surfaces
* Clear surfaces
* Supports transparency
* Plot points
* Draw horizontal and vertical lines
* Fill rectangles
* Draw rectangle outlines
* Write as a PNG stream
* Some predefined colors

## instructions

In progress.

## building the example for all platforms

Substitute the script file according to your own platform:

``` sh
cd example
./scripts/macos.sh
```

## creating and running a one-off build of the example

Substitute the file name according to your own platform:

``` sh
cd example
go build -o ./builds/macos/example && ./builds/macos/example
```
