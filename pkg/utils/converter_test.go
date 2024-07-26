package utils

import (
    "testing"
)

func TestToHex(t *testing.T) {
    result, err := ToHex("255")
    if err != nil {
        t.Fatalf("ToHex(255) returned error: %v", err)
    }
    if result != "0xFF" {
        t.Errorf("ToHex(255) = %s; want 0xFF", result)
    }
}

func TestFromHex(t *testing.T) {
  result, err := FromHex("0x10")
  if err != nil {
      t.Fatalf("FromHex(0x10) returned error: %v", err)
  }
  if result != 16 {
      t.Errorf("FromHex(0x10) = %d; want 16", result)
  }
}