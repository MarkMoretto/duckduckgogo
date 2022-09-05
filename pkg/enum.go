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
	m := map[ReturnFormat]string{0:"xml", 1:"json"}
	if _, ok := m[rf]; !ok {
		return "unknown"
	} else {
		return m[rf]
	}
}

