package entities

import (
	"bytes"
	"encoding/json"
)

type Founding int64

const (
	Startup Founding = iota
	Sme
)

func (f Founding) String() string {
	switch f {
	case Startup:
		return "Startup"
	case Sme:
		return "SME"
	}
	return "unknown"
}

func (f Founding) toID(val string) int64 {
	switch {
	case val == Startup.String():
		return int64(Startup)
	case val == Sme.String():
		return int64(Sme)
	}
	return -1
}

func (f Founding) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(f.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *Founding) UnmarshalJSON(b []byte) error {

	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*s = Founding(s.toID(j))

	return nil
}
