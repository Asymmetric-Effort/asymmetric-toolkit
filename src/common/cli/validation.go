package cli

type ValidationFunc func(i *string) bool  // Return true if valid, false if invalid.

