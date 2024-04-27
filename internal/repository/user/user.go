package user_repo

import (
	"MTCTrueTechHack/internal/entities"
	"MTCTrueTechHack/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func InitUserRepo(db *sqlx.DB) repository.UserRepo {
	return userRepo{
		db: db,
	}
}

func (u userRepo) Create(ctx context.Context, user entities.User) (int, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var userID int
	result := tx.QueryRowContext(ctx, `INSERT INTO users (id, name, surname, email) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.ID, user.Name, user.Surname, user.Email, userID)

	err = result.Scan(&userID)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return userID, err
}
