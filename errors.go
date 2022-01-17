package thumbnailer

// #include "ffmpeg.h"
import "C"
import (
	"fmt"
	"io"
)

// Cast FFmpeg error to Go error
func castError(err C.int) error {
	switch err {
	case C.AVERROR_EOF:
		return io.EOF
	case C.AVERROR_STREAM_NOT_FOUND:
		return ErrStreamNotFound
	default:
		return AVError(err)
	}
}

// AVError converts an FFmpeg error code to a Go error with a human-readable
// error message
type AVError C.int

// Error formats the FFmpeg error in human-readable format
func (f AVError) Error() string {
	buf := C.malloc(1024)
	defer C.free(buf)
	C.av_strerror(C.int(f), (*C.char)(buf), 1024)
	return fmt.Sprintf("ffmpeg: %s", C.GoString((*C.char)(buf)))
}

// Code returns the underlying AVERROR error code
func (f AVError) Code() C.int {
	return C.int(f)
}
