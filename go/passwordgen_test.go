package main

import "testing"

func TestGeneratePassword(t *testing.T) {
	const (
		LevelLow      = "low"
		LevelMedium   = "medium"
		LevelStrong   = "strong"
		LevelParanoid = "paranoid"
	)
	levels := []string{LevelLow, LevelMedium, LevelStrong, LevelParanoid}
	for _, lvl := range levels {
		seen := make(map[string]bool)
		for i := 0; i < 100; i++ {
			pw, err := GeneratePassword(24, lvl)
			if err != nil {
				t.Fatalf("level %s error: %v", lvl, err)
			}
			if len([]rune(pw)) != 24 {
				t.Fatalf("level %s length mismatch", lvl)
			}
			if seen[pw] {
				t.Fatalf("level %s duplicate generated", lvl)
			}
			seen[pw] = true
		}
	}
}
