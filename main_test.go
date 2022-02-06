package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	want := "Hello World!\n"

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if want != out {
		t.Fatalf(`Main say = %v, want match for %v`, out, want)
	}
}

func TestSayHello(t *testing.T) {
	want := "Hello World!"
	msg := SayHello()
	if want != msg {
		t.Fatalf(`SayHello returns %v, want match for %v`, msg, want)
	}
}

func TestGreetings(t *testing.T) {
	inputs := []string{
		"Amr",
		"Ahmed",
	}

	for _, name := range inputs {
		output := Greet(name)
		want := fmt.Sprintf("Welcome %v", name)
		if output != want {
			t.Fatalf(`Greet say = %v, want match for %v`, output, want)
		}
	}
}
