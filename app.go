package main

import (
	"context"
	"fmt"
	"strconv"
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

// OctalToBinary converts an octal string to a binary string.
func (a *App) OctalToBinary(octalString string) (string, error) {
	if octalString == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	// Handle the special case for "0" or "00..."
	isZero := true
	for _, r := range octalString {
		if r != '0' {
			isZero = false
			break
		}
	}
	if isZero {
		return "0", nil
	}

	// Parse the octal string to an integer (base 8).
	// The 0 prefix for octal is handled by ParseInt if the base is 0 or 8.
	// Here we explicitly use base 8.
	// strconv.ParseInt can handle leading zeros like "00507".
	decimalValue, err := strconv.ParseInt(octalString, 8, 64)
	if err != nil {
		// Check if the error is due to invalid characters
		for _, char := range octalString {
			if char < '0' || char > '7' {
				return "", fmt.Errorf("invalid octal string: contains non-octal character '%c'", char)
			}
		}
		// If it's another parsing error (e.g. out of range, though less likely for typical inputs)
		return "", fmt.Errorf("invalid octal string: %w", err)
	}

	// Convert the integer to a binary string.
	// strconv.FormatInt does not produce leading zeros for non-zero numbers.
	binaryString := strconv.FormatInt(decimalValue, 2)

	return binaryString, nil
}
