package main


// Transform const string into *string
func ConstToRef(s string) *string {
	var outs *string
	inRunes := []rune(s)
	*outs = string(inRunes)
	return outs
}