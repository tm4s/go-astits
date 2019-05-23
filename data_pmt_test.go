package astits

import (
	"testing"

	"github.com/tm4s/go-astitools/binary"
	"github.com/stretchr/testify/assert"
)

var pmt = &PMTData{
	ElementaryStreams: []*PMTElementaryStream{{
		ElementaryPID:               2730,
		ElementaryStreamDescriptors: descriptors,
		StreamType:                  StreamTypeMPEG1Audio,
	}},
	PCRPID:             5461,
	ProgramDescriptors: descriptors,
	ProgramNumber:      1,
}

func pmtBytes() []byte {
	w := astibinary.New()
	w.Write("111")                       // Reserved bits
	w.Write("1010101010101")             // PCR PID
	w.Write("1111")                      // Reserved
	descriptorsBytes(w)                  // Program descriptors
	w.Write(uint8(StreamTypeMPEG1Audio)) // Stream #1 stream type
	w.Write("111")                       // Stream #1 reserved
	w.Write("0101010101010")             // Stream #1 PID
	w.Write("1111")                      // Stream #1 reserved
	descriptorsBytes(w)                  // Stream #1 descriptors
	return w.Bytes()
}

func TestParsePMTSection(t *testing.T) {
	var offset int
	var b = pmtBytes()
	d := parsePMTSection(b, &offset, len(b), uint16(1))
	assert.Equal(t, d, pmt)
}
