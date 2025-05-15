package main

import (
	"context"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ValidateString checks if the input string conforms to the pattern a^n b^n c^n.
func (a *App) ValidateString(input string) string {
	if input == "" {
		return "Y" // n = 0 case
	}

	n := len(input)
	if n%3 != 0 {
		return "N"
	}

	count := n / 3

	// Check for 'a's
	for i := 0; i < count; i++ {
		if input[i] != 'a' {
			return "N"
		}
	}

	// Check for 'b's
	for i := count; i < 2*count; i++ {
		if input[i] != 'b' {
			return "N"
		}
	}

	// Check for 'c's
	for i := 2 * count; i < 3*count; i++ {
		if input[i] != 'c' {
			return "N"
		}
	}

	// Check if there are any other characters after the c's
	// This is implicitly handled by the n%3 check and the loop bounds,
	// but an explicit check for other characters in the string can be added for robustness
	// if needed, e.g. by checking allowed characters.
	// For now, the provided logic covers the a^n b^n c^n structure.

	// A more robust check for invalid characters:
	allowedChars := "abc"
	for _, char := range input {
		if !strings.ContainsRune(allowedChars, char) {
			return "N" // Found a character not in {a, b, c}
		}
	}

	return "Y"
}
