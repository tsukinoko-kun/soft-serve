package web

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/soft-serve/pkg/config"
)

type browserIndex struct{}

func (browserIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cfg := config.FromContext(ctx)
	if cfg == nil {
		http.Error(w, "Config not found", http.StatusInternalServerError)
		return
	}

	reposDir := filepath.Join(cfg.DataPath, "repos")
	reposDirs, err := os.ReadDir(reposDir)
	if err != nil {
		http.Error(w, "Failed to read repos directory", http.StatusInternalServerError)
		return
	}

	sb := strings.Builder{}
	sb.WriteString(`<ul>`)
	for _, route := range reposDirs {
		if !route.IsDir() {
			continue
		}
		name := route.Name()
		if strings.HasSuffix(name, ".git") {
			name = name[:len(name)-4]
		} else {
			continue
		}
		sb.WriteString(fmt.Sprintf(`<li><a href="%s">%s</a></li>`, name, name))
	}
	sb.WriteString(`</ul>`)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(sb.String()))
}
