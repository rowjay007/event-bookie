package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/rowjay007/event-bookie/internal/models"
)

type VenueRepository struct {
    db *sql.DB
}

func NewVenueRepository(db *sql.DB) *VenueRepository {
    return &VenueRepository{db}
}

func (r *VenueRepository) CreateVenue(ctx context.Context, venue *models.Venue) error {
    query := `
        INSERT INTO venues (name, description, capacity, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `
    _, err := r.db.ExecContext(ctx, query, venue.Name, venue.Description, venue.Capacity, time.Now(), time.Now())
    if err != nil {
        return err
    }
    return nil
}

func (r *VenueRepository) GetVenueByID(ctx context.Context, id int) (*models.Venue, error) {
    query := `
        SELECT id, name, description, capacity, created_at, updated_at
        FROM venues
        WHERE id = $1
    `
    row := r.db.QueryRowContext(ctx, query, id)
    venue := &models.Venue{}
    err := row.Scan(&venue.ID, &venue.Name, &venue.Description, &venue.Capacity, &venue.CreatedAt, &venue.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return venue, nil
}

// Implement other methods like UpdateVenue, DeleteVenue, GetVenues, etc.
