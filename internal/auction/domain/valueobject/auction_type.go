package valueobject

type AuctionType int

const (
	Bid AuctionType = 1
)

func (s AuctionType) IsValid() bool {
	switch s {
	case Bid:
		return true
	default:
		return false
	}
}

func (s AuctionType) String() string {
	switch s {
	case Bid:
		return "Bid"
	default:
		return "Unknown"
	}
}
