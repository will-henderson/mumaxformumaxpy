package xc

import (
	"fmt"
	"github.com/barnex/cuda4/safe"
)

type Conv struct {
	size          [3]int
	input, output [3][]float32
	realBuf       [3]safe.Float32s
	fftBuf        [3]safe.Float32s
}

func (c *Conv) Init(input, output [3][]float32, size [3]int) {
	N := prod(size)
	for c := 0; c < 3; c++ {
		if len(output[c]) != N || len(input[c]) != N {
			panic(fmt.Errorf("xc.Conv.Init: inconsistent sizes"))
		}
	}
	c.size = size
	c.input = input
	c.output = output

	// don't leak on 2nd init
	c.realBuf[0].Free()
	c.fftBuf[0].Free()

	r := safe.MakeFloat32s(3 * N)
	c.realBuf = [3]safe.Float32s{r.Slice(0*N, 1*N), r.Slice(1*N, 2*N), r.Slice(2*N, 3*N)}
}

func prod(size [3]int) int {
	return size[0] * size[1] * size[2]
}
