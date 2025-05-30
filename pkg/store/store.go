package store

// Store is an interface for managing repositories, users, and settings.
type Store interface {
	RepositoryStore
	UserStore
	CollaboratorStore
	IssueStore
	SettingStore
	LFSStore
	AccessTokenStore
	WebhookStore
}
