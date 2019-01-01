package curator

import (
	"testing"
)

func Test_runWorker(t *testing.T) {
	_, err := getNewsWithKeywords("python")
	if err != nil {
		t.Fatalf("Test failed!\nError: %s", err.Error())
	}
}
