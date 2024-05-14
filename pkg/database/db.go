// pkg/database/supabase.go

package database

import (
	"fmt"
	"os"

	"github.com/rowjay007/event-bookie/config"
)

// NewSupabaseDB creates a new Supabase database connection using configuration values
func NewSupabaseDB(cfg *config.Config) (*SupabaseDB, error) {
    // Initialize Supabase configuration
    supabaseURL := os.Getenv("SUPABASE_URL")
    supabaseKey := os.Getenv("SUPABASE_KEY")

    // Check if Supabase configuration is missing
    if supabaseURL == "" || supabaseKey == "" {
        return nil, fmt.Errorf("Supabase configuration missing")
    }

    // You can add more initialization logic here if needed

    return &SupabaseDB{
        cfg:          cfg,
        supabaseURL:  supabaseURL,
        supabaseKey:  supabaseKey,
    }, nil
}

type SupabaseDB struct {
    cfg          *config.Config
    supabaseURL  string
    supabaseKey  string
    // Add any other fields or methods as needed
}
