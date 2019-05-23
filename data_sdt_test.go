package astits

import (
	"testing"

	"github.com/tm4s/go-astitools/binary"
	"github.com/stretchr/testify/assert"
)

var sdt = &SDTData{
	OriginalNetworkID: 2,
	Services: []*SDTDataService{{
		Descriptors:            descriptors,
		HasEITPresentFollowing: true,
		HasEITSchedule:         true,
		HasFreeCSAMode:         true,
		RunningStatus:          5,
		ServiceID:              3,
	}},
	TransportStreamID: 1,
}

func sdtBytes() []byte {
	w := astibinary.New()
	w.Write(uint16(2))  // Original network ID
	w.Write(uint8(0))   // Reserved for future use
	w.Write(uint16(3))  // Service #1 id
	w.Write("000000")   // Service #1 reserved for future use
	w.Write("1")        // Service #1 EIT schedule flag
	w.Write("1")        // Service #1 EIT present/following flag
	w.Write("101")      // Service #1 running status
	w.Write("1")        // Service #1 free CA mode
	descriptorsBytes(w) // Service #1 descriptors
	return w.Bytes()
}

func TestParseSDTSection(t *testing.T) {
	var offset int
	var b = sdtBytes()
	d := parseSDTSection(b, &offset, len(b), uint16(1))
	assert.Equal(t, d, sdt)
}
