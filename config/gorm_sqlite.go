package config

import (
	"path/filepath"
)

type Sqlite struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
}

func (s *Sqlite) Dsn() string {
	return filepath.Join(s.Path, s.Dbname+".db")
}

func (s *Sqlite) GetLogMode() string {
	return s.LogMode
}
