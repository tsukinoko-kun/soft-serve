package store

import (
	"context"

	"github.com/charmbracelet/soft-serve/pkg/db"
	"github.com/charmbracelet/soft-serve/pkg/db/models"
)

// IssueStore is an interface for managing issues.
type IssueStore interface {
	ListIssuesByRepo(ctx context.Context, h db.Handler, repo string) ([]models.Issue, error)
}
