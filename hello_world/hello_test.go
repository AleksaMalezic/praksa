package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t*testing.T){
		got := Hello("Chriss", "")
		want := "Hello, Chriss"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say, 'Hello, World' when an empty string is supplied", func(t*testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Spanish", func(t*testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Fre*ch", func(t*testing.T){
		got := Hello("Fourier", "French")
		want := "Bonjour, Fourier"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Serbian", func(t*testing.T){
		got := Hello("Aleksa", "Serbian")
		want := "Zdravo, Aleksa"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}