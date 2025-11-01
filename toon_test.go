package toon

import (
	"testing"
)

func TestEncode(t *testing.T) {
	t.Skip("Implementation pending - namespace reservation")

	data := map[string]interface{}{
		"users": []map[string]interface{}{
			{"id": 1, "name": "Alice", "role": "admin"},
			{"id": 2, "name": "Bob", "role": "user"},
		},
	}

	result, err := Encode(data, nil)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	expected := `users[2]{id,name,role}:
  1,Alice,admin
  2,Bob,user`

	if result != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}

func TestDecode(t *testing.T) {
	t.Skip("Implementation pending - namespace reservation")

	input := `users[2]{id,name,role}:
  1,Alice,admin
  2,Bob,user`

	result, err := Decode(input, nil)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Validate structure
	data, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map[string]interface{}")
	}

	users, ok := data["users"].([]interface{})
	if !ok {
		t.Fatal("Expected users to be a slice")
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestEncodeWithOptions(t *testing.T) {
	t.Skip("Implementation pending - namespace reservation")

	data := map[string]interface{}{
		"tags": []string{"admin", "ops", "dev"},
	}

	opts := &EncodeOptions{
		Indent:       4,
		Delimiter:    '\t',
		LengthMarker: true,
	}

	result, err := Encode(data, opts)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	expected := "tags[#3\t]: admin\tops\tdev"

	if result != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}
