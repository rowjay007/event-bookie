package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/rowjay007/event-bookie/internal/models"
)

type BookingRepository struct {
    db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
    return &BookingRepository{db}
}

func (r *BookingRepository) CreateBooking(ctx context.Context, booking *models.Booking) error {
    query := `
        INSERT INTO bookings (user_id, event_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
    `
    _, err := r.db.ExecContext(ctx, query, booking.UserID, booking.EventID, time.Now(), time.Now())
    if err != nil {
        return err
    }
    return nil
}

func (r *BookingRepository) GetBookingByID(ctx context.Context, id int) (*models.Booking, error) {
    query := `
        SELECT id, user_id, event_id, created_at, updated_at
        FROM bookings
        WHERE id = $1
    `
    row := r.db.QueryRowContext(ctx, query, id)
    booking := &models.Booking{}
    err := row.Scan(&booking.ID, &booking.UserID, &booking.EventID, &booking.CreatedAt, &booking.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return booking, nil
}

// Implement other methods like UpdateBooking, DeleteBooking, GetBookings, etc.
