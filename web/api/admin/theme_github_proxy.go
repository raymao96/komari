package admin

import (
	"strings"

	"github.com/raymao96/komari/pkg/config"
	"github.com/raymao96/komari/pkg/themehttp"
)

func loadGitHubProxy() string {
	value, err := config.GetAs[string](config.GitHubProxyKey)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(value)
}

func applyThemeGitHubProxy(rawURL string) string {
	return themehttp.RewriteGitHubURL(rawURL, loadGitHubProxy())
}
