[![GoDoc](https://godoc.org/github.com/MEnigma/thumbnailer?status.svg)](https://godoc.org/github.com/MEnigma/thumbnailer)
[![Build Status](https://travis-ci.com/MEnigma/thumbnailer.svg?branch=master)](https://travis-ci.com/MEnigma/thumbnailer)
# thumbnailer
Package thumbnailer provides a more efficient media thumbnailer than available
with native Go processing libraries through ffmpeg bindings.

Use 
```
go get -u github.com/MEnigma/thumbnailer/v2
```
to install the library in your project. 

For a comprehensive list of file formats supported by default see
main.go:Process().

## Dependencies
* Go >= 1.10
* C11 compiler
* make
* pkg-config
* pthread
* ffmpeg >= 4.1 libraries (libavcodec, libavutil, libavformat, libswscale)

NB:
* ffmpeg should be compiled with all the dependency libraries for formats you
want to process. On most Linux distributions you should be fine with
the packages in the stock repositories.
* Ubuntu patches to ffmpeg on some Ubuntu versions <19.10 break this library.
In this case, please compile from unmodified ffmpeg sources using:

```
sudo apt build-dep ffmpeg
git clone https://git.ffmpeg.org/ffmpeg.git ffmpeg
cd ffmpeg
git checkout n4.1
./configure
make -j`nproc`
sudo make install
```
