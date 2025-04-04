package custom_identifiers

type Identifier interface {
	Identifier() // Method, which should be implemented by struct to implement the interface
	Equals(other any) bool
	String() string
}
