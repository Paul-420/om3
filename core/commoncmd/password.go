package commoncmd

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func ReadPasswordFromStdinOrPrompt(prompt string) ([]byte, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		if b, err := os.ReadFile("/dev/stdin"); err != nil {
			return nil, err
		} else {
			return b, nil
		}
	}

	fmt.Fprintf(os.Stderr, prompt)
	if b, err := term.ReadPassword(int(os.Stdin.Fd())); err != nil {
		fmt.Fprintln(os.Stderr, "")
		return nil, err
	} else {
		fmt.Fprintln(os.Stderr, "")
		return b, nil
	}
}
