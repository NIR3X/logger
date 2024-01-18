package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

func fprintln(w io.Writer, a ...any) (int, error) {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Fprintln(
		w,
		append([]any{
			fmt.Sprintf(
				"%s %s:%d:",
				time.Now().Format("2006/01/02 15:04:05"),
				file,
				line,
			),
		}, a...)...,
	)
}

func Fprintln(w io.Writer, a ...any) (int, error) {
	return fprintln(w, a...)
}

func Eprintln(a ...any) (int, error) {
	return fprintln(os.Stderr, append([]any{"error:"}, a...)...)
}

func Println(a ...any) (int, error) {
	return fprintln(os.Stdout, a...)
}
