package integers

import "testing"
import "fmt"

func TestAdder(t*testing.T){
	t.Run("Problem u funkciji Add(x, y)", func(t*testing.T){
		sum := Add(2, 2)
		expected := 4
	
		assertCorrectMessage(t, sum, expected)
	})
	
}

func ExampleAdder() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, got, want int){
	t.Helper()
	if got != want{
		t.Errorf("expected %d got %d", want, got)
	}
}