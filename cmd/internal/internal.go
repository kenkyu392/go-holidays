package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kenkyu392/go-holidays"
)

// Data ...
type Data struct {
	UpdateTime time.Time         `json:"updateTime"`
	Holidays   holidays.Holidays `json:"holidays"`
}

// FilePath creates a normalized file path.
func FilePath(dir, file string) (string, error) {
	outDir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	stat, err := os.Stat(outDir)
	if err != nil {
		return "", err
	}

	if !stat.IsDir() {
		return "", fmt.Errorf("%s is not a directory", outDir)
	}
	outDir = outDir + "/"
	return filepath.Join(outDir, file), nil
}
