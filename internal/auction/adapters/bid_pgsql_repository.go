package adapters

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/aggregate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgSqlBidRepository struct {
	db *gorm.DB
}

func NewPgSqlBidRepository(db *gorm.DB) *PgSqlBidRepository {
	if db == nil {
		panic("missing db")
	}

	return &PgSqlBidRepository{db: db}
}

func (r *PgSqlBidRepository) Save(bid aggregate.Bid) error {
	if err := r.db.Omit(clause.Associations).Create(&bid).Error; err != nil {
		return err
	}
	return nil
}

func (r *PgSqlBidRepository) GetAllBids() ([]aggregate.Bid, error) {
	var bids []aggregate.Bid
	if err := r.db.Find(&bids).Error; err != nil {
		return nil, err
	}
	return bids, nil
}

func (r *PgSqlBidRepository) GetBidsByAuctionId(auctionId uuid.UUID) ([]aggregate.Bid, error) {
	var bids []aggregate.Bid
	if err := r.db.Where("auction_id = ?", auctionId).Find(&bids).Error; err != nil {
		return nil, err
	}
	return bids, nil
}
