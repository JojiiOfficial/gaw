package gaw

import "testing"

func TestResolveDir(t *testing.T) {
	dir := "."
	result := GetCurrentDir()

	if ResolveFullPath(dir) != result {
		t.Errorf("Expected '%s' got '%s'", dir, result)
	}
}
