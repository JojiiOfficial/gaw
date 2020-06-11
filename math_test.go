package gaw

import "testing"

func TestGetFigureCount(t *testing.T) {
	nr := 920341097234810285
	fCount := GetFigureCountInt(nr)

	if fCount != 18 {
		t.Errorf("Expected 18 but got %d", fCount)
	}
}

func BenchmarkGetFigureCount(b *testing.B) {
	b.ReportAllocs()

	nr := 920341097234810285

	for i := 0; i < b.N; i++ {
		GetFigureCountInt(nr)
	}
}
