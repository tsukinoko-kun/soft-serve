package repo

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/soft-serve/pkg/db/models"
	"github.com/charmbracelet/soft-serve/pkg/proto"
	"github.com/charmbracelet/soft-serve/pkg/ui/common"
)

// Issues is the issues component page.
type Issues struct {
	spinner   spinner.Model
	common    common.Common
	repo      proto.Repository
	ref       RefMsg
	issues    []models.Issue
	isLoading bool
}

// IssuesMsg is a message sent when the issues are loaded.
type IssuesMsg struct {
	issues []models.Issue
}

// NewIssues creates a new issues model.
func NewIssues(common common.Common) *Issues {
	s := spinner.New(spinner.WithSpinner(spinner.Dot),
		spinner.WithStyle(common.Styles.Spinner))
	return &Issues{
		spinner:   s,
		common:    common,
		isLoading: true,
	}
}

// Path implements common.TabComponent.
func (i *Issues) Path() string {
	return ""
}

// TabName returns the name of the tab.
func (i *Issues) TabName() string {
	return "Issues"
}

// SetSize implements common.Component.
func (i *Issues) SetSize(width, height int) {
	i.common.SetSize(width, height)
}

// ShortHelp implements help.KeyMap.
func (i *Issues) ShortHelp() []key.Binding {
	return []key.Binding{}
}

// FullHelp implements help.KeyMap.
func (i *Issues) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

// Init implements tea.Model.
func (i *Issues) Init() tea.Cmd {
	i.isLoading = true
	return tea.Batch(i.spinner.Tick, i.updateIssuesCmd)
}

// Update implements tea.Model.
func (i *Issues) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case IssuesMsg:
		i.issues = msg.issues
		i.isLoading = false
	case RepoMsg:
		i.repo = msg
	case RefMsg:
		i.ref = msg
		cmds = append(cmds, i.Init())
	case tea.WindowSizeMsg:
		i.SetSize(msg.Width, msg.Height)
	case spinner.TickMsg:
		if i.isLoading && i.spinner.ID() == msg.ID {
			s, cmd := i.spinner.Update(msg)
			i.spinner = s
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}
	return i, tea.Batch(cmds...)
}

// View implements tea.Model.
func (i *Issues) View() string {
	if i.isLoading {
		return renderLoading(i.common, i.spinner)
	}
	// render list of issues
}

// SpinnerID implements common.TabComponent.
func (i *Issues) SpinnerID() int {
	return i.spinner.ID()
}

// StatusBarValue implements statusbar.StatusBar.
func (i *Issues) StatusBarValue() string {
	return ""
}

// StatusBarInfo implements statusbar.StatusBar.
func (i *Issues) StatusBarInfo() string {
	return ""
}

func (r *Issues) updateIssuesCmd() tea.Msg {
	m := IssuesMsg{}
	if r.repo == nil {
		return common.ErrorMsg(common.ErrMissingRepo)
	}
	be := r.common.Backend()
	if issues, err := be.Issues(r.common.Context(), r.repo.Name()); err != nil {
		return common.ErrorMsg(err)
	} else {
		m.issues = issues
	}
	return m
}
