package valueobject

type AuctionStatus int

const (
	Unknown AuctionStatus = 0
	Created AuctionStatus = 1
	Active  AuctionStatus = 2
	Closed  AuctionStatus = 3
)

func (s AuctionStatus) IsValid() bool {
	switch s {
	case Unknown, Created, Active, Closed:
		return true
	default:
		return false
	}
}

func (s AuctionStatus) String() string {
	switch s {
	case Unknown:
		return "Unknown"
	case Created:
		return "Created"
	case Active:
		return "Active"
	case Closed:
		return "Closed"
	default:
		return "Unknown"
	}
}
