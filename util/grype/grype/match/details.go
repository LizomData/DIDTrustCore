package match

import (
	"fmt"
	"strings"

	"github.com/gohugoio/hashstructure"
)

type Details []Detail

type Detail struct {
	Type       Type        `json:"type"`
	SearchedBy interface{} `json:"searchedBy"`
	Found      interface{} `json:"found"`
	Matcher    MatcherType `json:"matcher"`
	Confidence float64     `json:"confidence"`
}

// String is the string representation of select match fields.
func (m Detail) String() string {
	return fmt.Sprintf("Detail(searchedBy=%q found=%q matcher=%q)", m.SearchedBy, m.Found, m.Matcher)
}

func (m Details) Matchers() (tys []MatcherType) {
	if len(m) == 0 {
		return nil
	}
	for _, d := range m {
		tys = append(tys, d.Matcher)
	}
	return tys
}

func (m Details) Types() (tys []Type) {
	if len(m) == 0 {
		return nil
	}
	for _, d := range m {
		tys = append(tys, d.Type)
	}
	return tys
}

func (m Detail) ID() string {
	f, err := hashstructure.Hash(&m, &hashstructure.HashOptions{
		ZeroNil:      true,
		SlicesAsSets: true,
	})
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", f)
}

func (m Details) Len() int {
	return len(m)
}

func (m Details) Less(i, j int) bool {
	a := m[i]
	b := m[j]

	if a.Type != b.Type {
		// exact-direct-match < exact-indirect-match < cpe-match

		at := typeOrder[a.Type]
		bt := typeOrder[b.Type]
		if at == 0 {
			return false
		} else if bt == 0 {
			return true
		}
		return at < bt
	}

	// sort by confidence
	if a.Confidence != b.Confidence {
		// flipped comparison since we want higher confidence to be first
		return a.Confidence > b.Confidence
	}

	// if the types are the same, then sort by the ID (costly, but deterministic)
	return strings.Compare(a.ID(), b.ID()) < 0
}

func (m Details) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
