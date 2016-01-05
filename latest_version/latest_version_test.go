package latest_version

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsLatest(t *testing.T) {
	b := New(nil)
	var i (*LatestVersion) = nil
	err := AssertThat(b, Implements(i).Message("check implements type Latest"))
	if err != nil {
		t.Fatal(err)
	}
}
