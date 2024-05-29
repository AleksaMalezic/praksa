package main

import "testing"
import "slices"

func TestSum(t*testing.T){
	t.Run("Example with slices", func(t*testing.T){
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, got, expected)
	})
}

func TestSumAll(t*testing.T){
	t.Run("Varying number of slices", func(t*testing.T){

		got := SumAll([]int{1, 2}, []int{0, 9, 8})
		expected := []int{3, 17}

		if !slices.Equal(got, expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}

func TestSumAllTails(t *testing.T){

	CheckSums := func(t testing.TB, got, expected []int){
		t.Helper()
		if !slices.Equal(got, expected){
			t.Errorf("got %v expected %v", got, expected)
		}
	}
	t.Run("make the su of some tails", func(t*testing.T){
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		CheckSums(t, got, expected)
	})
	t.Run("make the su of some tails", func(t*testing.T){
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		CheckSums(t, got, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, expected int){
	if got != expected{
		t.Errorf("got %d expected %d", got, expected)
	}
}