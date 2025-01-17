package utils

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"sync"
)

// Copyright https://github.com/ntk148v/ggnf/commit/e076671fdb26271def018a7636c4b58a8aa57554

type LineWriter struct {
	*MultiProgressBar
	id int
}

func (lw *LineWriter) Write(p []byte) (n int, err error) {
	lw.guard.Lock()
	defer lw.guard.Unlock()
	lw.move(lw.id, lw.output)
	return lw.output.Write(p)
}

type MultiProgressBar struct {
	output     io.Writer
	curLine    int
	Bars       []*progressbar.ProgressBar
	guard      sync.Mutex
	writeGuard sync.RWMutex
}

func NewMultiProgressBar(output io.Writer) *MultiProgressBar {
	mpb := &MultiProgressBar{
		curLine:    0,
		Bars:       []*progressbar.ProgressBar{},
		guard:      sync.Mutex{},
		writeGuard: sync.RWMutex{},
		output:     output,
	}

	return mpb
}

func (mpb *MultiProgressBar) Add(b *progressbar.ProgressBar) int {
	mpb.writeGuard.Lock()
	defer mpb.writeGuard.Unlock()
	mpb.Bars = append(mpb.Bars, b)
	id := len(mpb.Bars) - 1
	progressbar.OptionSetWriter(&LineWriter{
		MultiProgressBar: mpb,
		id:               id,
	})(b)

	return id
}

func (mpb *MultiProgressBar) Get(id int) *progressbar.ProgressBar {
	return mpb.Bars[id]
}

// Move cursor to the beginning of the current progressbar.
func (mpb *MultiProgressBar) move(id int, writer io.Writer) (int, error) {
	bias := mpb.curLine - id
	mpb.curLine = id
	if bias > 0 {
		// move up
		return fmt.Fprintf(writer, "\r\033[%dA", bias)
	} else if bias < 0 {
		// move down
		return fmt.Fprintf(writer, "\r\033[%dB", -bias)
	}
	return 0, nil
}

// End Move cursor to the end of the Progressbars.
func (mpb *MultiProgressBar) End() {
	mpb.move(len(mpb.Bars), mpb.output)
}
