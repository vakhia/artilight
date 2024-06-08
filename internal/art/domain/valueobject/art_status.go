package valueobject

type ArtStatus int

const (
	Draft     ArtStatus = 0
	Published ArtStatus = 1
	Sold      ArtStatus = 2
)

func (s ArtStatus) IsValid() bool {
	switch s {
	case Draft, Published, Sold:
		return true
	default:
		return false
	}
}

func (s ArtStatus) String() string {
	switch s {
	case Draft:
		return "Draft"
	case Published:
		return "Published"
	case Sold:
		return "Sold"
	default:
		return "Unknown"
	}
}
