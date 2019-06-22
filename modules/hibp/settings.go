package hibp

import (
	"time"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const (
	defaultTitle   = "HIBP"
	minRefreshSecs = 21600 // TODO: Finish implementing this
)

type colors struct {
	ok    string
	pwned string
}

// Settings defines the configuration properties for this module
type Settings struct {
	colors
	common *cfg.Common

	accounts []string
	since    string
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		accounts: wtf.ToStrs(ymlConfig.UList("accounts")),
		since:    ymlConfig.UString("since", ""),
	}

	settings.colors.ok = ymlConfig.UString("colors.ok", "green")
	settings.colors.pwned = ymlConfig.UString("colors.pwned", "red")

	return &settings
}

// HasSince returns TRUE if there's a valid "since" value setting, FALSE if there is not
func (sett *Settings) HasSince() bool {
	if sett.since == "" {
		return false
	}

	_, err := sett.SinceDate()
	if err != nil {
		return false
	}

	return true
}

// SinceDate returns the "since" settings as a proper Time instance
func (sett *Settings) SinceDate() (time.Time, error) {
	dt, err := time.Parse("2006-01-02", sett.since)
	if err != nil {
		return time.Now(), err
	}

	return dt, nil
}