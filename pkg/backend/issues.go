package backend

import (
	"context"

	"github.com/charmbracelet/soft-serve/pkg/db"
	"github.com/charmbracelet/soft-serve/pkg/db/models"
	"github.com/charmbracelet/soft-serve/pkg/utils"
)

// Issues returns the repository's issues.
func (d *Backend) Issues(ctx context.Context, repo string) ([]models.Issue, error) {
	repo = utils.SanitizeRepo(repo)
	var issues []models.Issue
	if err := d.db.TransactionContext(ctx, func(tx *db.Tx) error {
		var err error
		issues, err = d.store.ListIssuesByRepo(ctx, tx, repo)
		return err
	}); err != nil {
		return nil, db.WrapError(err)
	}

	return issues, nil
}
