package inputs

import (
	"github.com/google/uuid"
)

type BrandRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt string     `json:"published_at,omitempty"`
	BrandID     *uuid.UUID `json:"brand_id",omitempty`
}
