package events

import (
	"MTCTrueTechHack/internal/entities"
	"MTCTrueTechHack/internal/repository"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type eventsRepo struct {
	db *sqlx.DB
}

func InitEventRepo(db *sqlx.DB) repository.EventsRepo {
	return eventsRepo{
		db: db,
	}
}

func (e eventsRepo) CreateAll(ctx context.Context, events []entities.Event) error {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var valueStrings = make([]string, len(events))
	var valueArgs = make([]interface{}, 0, len(events))
	for i, node := range events {
		valueStrings = append(
			valueStrings,
			fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d)",
				i*7+1, i*7+2, i*7+3, i*7+4, i*7+5, i*7+6, i*7+7))
		valueArgs = append(
			valueArgs,
			node.Title, node.Category, node.Location, node.MinPrice, node.ClosestStartTime, node.LastStartTime, node.Tags)
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (e eventsRepo) GetAll(ctx context.Context) ([]entities.Event, error) {
	result, err := e.db.QueryContext(ctx, `SELECT * FROM events`)
	if err != nil {
		return []entities.Event{}, err
	}

	var events []entities.Event
	for result.Next() {
		var event entities.Event
		if err = result.Scan(&event); err != nil {
			return []entities.Event{}, err
		}

		events = append(events, event)
	}

	return events, nil
}
