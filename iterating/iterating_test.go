package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Example with slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, got, expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Varying number of slices", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{0, 9, 8})
		expected := []int{3, 17}

		if !slices.Equal(got, expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	CheckSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	}
	t.Run("make the su of some tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		CheckSums(t, got, expected)
	})
	t.Run("make the su of some tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		CheckSums(t, got, expected)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatinate of strings", func(t *testing.T) {
		multiply := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "bb", "c"}, multiply, ""), "abbc")
	})
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}
	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}

func assertCorrectMessage(t testing.TB, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func AssertEqual[T any](t testing.TB, got, expected T) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}
