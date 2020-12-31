package golang_coding_problems

// IsBpermutationOfA Checks if B is permutation of A
func IsBpermutationOfA(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	//cheating a little bit since freqMatrix is already implemented
	aFreqMatrix := makeFreqMatrix(a)
	bFreqMatrix := makeFreqMatrix(b)

	return equalFreqMatrix(aFreqMatrix, bFreqMatrix)
}
