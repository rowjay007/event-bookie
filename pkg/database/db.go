// pkg/database/supabase.go

package database

import (
	"fmt"
	supabase "github.com/nedpals/supabase-go"
	"github.com/rowjay007/event-bookie/config"
)

// NewSupabaseDB creates a new Supabase database connection using configuration values
func NewSupabaseDB(cfg *config.Config) (*SupabaseDB, error) {
	// Check if Supabase configuration is missing
	if cfg.SupabaseURL == "" || cfg.SupabaseKey == "" {
		return nil, fmt.Errorf("Supabase configuration missing")
	}

	// Create a new Supabase client
	client := supabase.CreateClient(cfg.SupabaseURL, cfg.SupabaseKey)

	// Log connection success
	fmt.Println("Connected to Supabase successfully")

	// You can add more initialization logic here if needed

	return &SupabaseDB{
		cfg:    cfg,
		client: client,
	}, nil
}

type SupabaseDB struct {
	cfg    *config.Config
	client *supabase.Client
	// Add any other fields or methods as needed
}
