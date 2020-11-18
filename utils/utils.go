package utils

// Contains checks whether a slice of options contains the target option
func Contains(options []string, target string) bool {
	for _, option := range options {
		if option == target {
			return true
		}
	}
	return false
}
