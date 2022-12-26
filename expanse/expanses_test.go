package expanse_test

import "testing"

type Expenses struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

func TestExpanse(t *testing.T) {
	e := Expenses{
		Id:     1,
		Title:  "test",
		Amount: 100.00,
		Note:   "test note",
		Tags:   []string{"gadget"},
	}

	if e.Id != 1 {
		t.Errorf("Expected id to be 1, got %d", e.Id)
	}
	if e.Title != "test" {
		t.Errorf("Expected title to be 'test', got '%s'", e.Title)
	}
	if e.Amount != 100.00 {
		t.Errorf("Expected amount to be 100.00, got %f", e.Amount)
	}
	if e.Note != "test note" {
		t.Errorf("Expected note to be 'test note', got '%s'", e.Note)
	}
	if len(e.Tags) != 1 || e.Tags[0] != "gadget" {
		t.Errorf("Expected tags to be ['gadget'], got %v", e.Tags)
	}
}

func TestEr(t *testing.T) {
	err := Err{
		Message: "test error",
	}
	if err.Message != "test error" {
		t.Errorf("Expected message to be 'test error', got '%s'", err.Message)
	}
}
