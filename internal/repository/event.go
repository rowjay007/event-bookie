package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/rowjay007/event-bookie/internal/models"
)

type EventRepository struct {
    db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
    return &EventRepository{db}
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *models.Event) error {
    query := `
        INSERT INTO events (name, description, start_date, end_date, venue_id, organizer_id, category_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `
    _, err := r.db.ExecContext(ctx, query, event.Name, event.Description, event.StartDate, event.EndDate, event.VenueID, event.OrganizerID, event.CategoryID, time.Now(), time.Now())
    if err != nil {
        return err
    }
    return nil
}

// Implement other methods like GetEventByID, UpdateEvent, DeleteEvent, etc.
