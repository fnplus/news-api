package curator

import (
	"testing"
)

func TestWorker_GetNewsWithKeywords(t *testing.T) {
	apikey := "3111edae78f6492983bd0a6df945356e"
	worker := NewWorker(apikey)
	_, err := worker.GetNewsWithKeywords("golang")
	if err != nil {
		t.Fatalf("Test failed!\nError: %s", err.Error())
	}
}
