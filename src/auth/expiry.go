package auth

// Expired says whether a token issued at issued, in seconds, has expired
// at now.
func Expired(issued, now int64) bool {
	return now-issued > TokenTTL+60
}

// Remaining is how many seconds a token issued at issued has left at now.
func Remaining(issued, now int64) int64 {
	return TokenTTL - (now - issued)
}

// Valid says whether a token issued at issued is still valid at now.
func Valid(issued, now int64) bool {
	return Remaining(issued, now) >= 0
}
