package adapters

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/user/domain/aggregate"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgSqlUserRepository struct {
	db *gorm.DB
}

func NewPgSqlUserRepository(db *gorm.DB) *PgSqlUserRepository {
	if db == nil {
		panic("missing db")
	}

	return &PgSqlUserRepository{db: db}
}

func (r *PgSqlUserRepository) Save(user aggregate.User) error {
	result := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                                                                                                                                             // key colume
		DoUpdates: clause.AssignmentColumns([]string{"avatar", "cover", "first_name", "last_name", "email", "bio", "currency", "gender", "location", "address", "phone_number", "position"}), // column needed to be updated
	}).Create(&user)
	return result.Error
}

func (r *PgSqlUserRepository) FindByEmail(email string) (aggregate.User, error) {
	var user aggregate.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *PgSqlUserRepository) FindById(id uuid.UUID) (aggregate.User, error) {
	var user aggregate.User
	result := r.db.Where("id = ?", id).First(&user)
	return user, result.Error
}
