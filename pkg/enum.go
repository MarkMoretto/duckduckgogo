package pkg

// Enums

// URL return format.
type ReturnFormat uint8

// URL return format enum.
const (
	Xml  ReturnFormat = iota
	Json
)

// URL return format stringer.
func (rf ReturnFormat) String() string {
	return []string{"xml", "json"}[rf]
}


// func (rf ReturnFormat) String() string {
// 	m := map[string]bool{"xml":true, "json":true}
// 	if _, ok := m[rf]; ok {
// 		return ""
// 	}
// 	return map[string]bool{"xml":true, "json":true}[rf]
// }