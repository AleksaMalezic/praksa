package iteration

import "testing"

func TestRepeat(t*testing.T){
	t.Run("Default example for Repeat func", func(t*testing.T){
		got := Repeat("a", 0)
		expected := "aaaaa"

		assertCorrectMessage(t, got, expected)
	})
	t.Run("case when the number of repeats is defined", func(t*testing.T){
		got := Repeat("a", 7)
		expected:= "aaaaaaa"

		assertCorrectMessage(t, got, expected)
	})
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++{
		Repeat("a", 5)
	}
}

func assertCorrectMessage(t testing.TB, got, expected string){
	t.Helper()
	if got != expected{
		t.Errorf("got %q expected %q", got, expected)
	}
}