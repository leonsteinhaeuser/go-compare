package compare

type Reports []Report

// swap swaps the values of a and b.
func (r *Reports) swap() {
	for i := range *r {
		(*r)[i].swap()
	}
}

type Report struct {
	// Type defines the actual type of A and B
	// This field has no functional purpose and only exists for discovery.
	Type string
	// Index is the index of the byte that differs
	Index int
	// Original represents the value of the byte at index in A
	Original *byte
	// New defines the value of b passed to the Equal function.
	New byte
}

// swap swaps the values of a and b.
func (r *Report) swap() {
	new := r.New
	orig := r.Original
	r.Original = &new
	r.New = *orig
}
