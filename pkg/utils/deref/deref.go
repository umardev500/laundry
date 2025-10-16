package deref

// String returns the value of *s or an empty string if nil.
func String(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// Bool returns the value of *b or false if nil.
func Bool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// Float64 returns the value of *f or 0 if nil.
func Float64(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}

// Int returns the value of *i or 0 if nil.
func Int(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

// Enum returns the value of *e or the provided default if nil.
func Enum[T any](e *T, defaultValue T) T {
	if e != nil {
		return *e
	}
	return defaultValue
}
