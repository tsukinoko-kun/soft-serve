package database

import (
	"context"

	"github.com/charmbracelet/soft-serve/pkg/db"
	"github.com/charmbracelet/soft-serve/pkg/db/models"
	"github.com/charmbracelet/soft-serve/pkg/store"
	"github.com/charmbracelet/soft-serve/pkg/utils"
)

type issueStore struct{}

var _ store.IssueStore = (*issueStore)(nil)

// ListIssuesByRepo implements store.CollaboratorStore.
func (*issueStore) ListIssuesByRepo(ctx context.Context, tx db.Handler, repo string) ([]models.Issue, error) {
	var m []models.Issue

	repo = utils.SanitizeRepo(repo)
	query := tx.Rebind(`
		SELECT
			issues.*
		FROM
			issues
		INNER JOIN repos ON repos.id = issues.repo_id
		WHERE
			repos.name = ?
	`)

	err := tx.SelectContext(ctx, &m, query, repo)
	return m, err
}
