package obs

import (
	"testing"
)

func TestGenPresign(t *testing.T) {
	obs := OBS{}
	v := obs.GenPresign()
	if v != 1 {
		t.Error()
	}
}
