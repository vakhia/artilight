package adapters

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/aggregate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgSqlAuctionRepository struct {
	db *gorm.DB
}

func NewPgSqlAuctionRepository(db *gorm.DB) *PgSqlAuctionRepository {
	if db == nil {
		panic("missing db")
	}

	return &PgSqlAuctionRepository{db: db}
}

func (r *PgSqlAuctionRepository) Save(auction aggregate.Auction) error {
	if err := r.db.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                       // key colume
		DoUpdates: clause.AssignmentColumns([]string{"current_price"}), // column needed to be updated
	}).Create(&auction).Error; err != nil {
		return err
	}
	return nil
}

func (r *PgSqlAuctionRepository) GetAllAuctions() ([]aggregate.Auction, error) {
	var auctions []aggregate.Auction
	if err := r.db.Find(&auctions).Error; err != nil {
		return nil, err
	}
	return auctions, nil
}

func (r *PgSqlAuctionRepository) FindAuctionById(id uuid.UUID) (aggregate.Auction, error) {
	var auction aggregate.Auction
	if err := r.db.Where("id = ?", id).Preload("Item").Preload("Bids").Preload("Bids.Item").Preload("Bids.Bidder").First(&auction).Error; err != nil {
		return aggregate.Auction{}, err
	}
	return auction, nil
}
