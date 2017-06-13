package latest_tar_gz_url

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsLatest(t *testing.T) {
	b := New(nil)
	var i (*LatestTarGzUrl) = nil
	err := AssertThat(b, Implements(i).Message("check implements type Latest"))
	if err != nil {
		t.Fatal(err)
	}
}
