package bundle

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type AutomationAnnotation struct {
}

type PackageJSON struct {
	// Main defines the entrypoint script
	Main string `json:"main"`

	// Version holds the version of the package
	Version string `json:"version"`

	// License describes the package license
	License string `json:"license"`

	// Automation holds additional information for the automation
	// engine.
	Automation AutomationAnnotation `json:"automation"`
}

type Bundle struct {
	// Path holds the path to the bundle root
	Path string

	// Main holds the path of the entrypoint script that
	// should be executed.
	Main string

	// ScriptContent holds the content of the main entrypoint file.
	ScriptContent string
}

func Discover(root string) ([]*Bundle, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var result []*Bundle

	for _, e := range entries {
		path := filepath.Join(root, e.Name())
		if e.IsDir() {
			b, err := Load(path)
			if err != nil {
				slog.Error("failed to load bundle", "error", err, "path", e.Name())
				continue
			}

			result = append(result, b)

			continue
		}

		ext := filepath.Ext(e.Name())

		var (
			target string
		)

		var gzip bool

		switch ext {
		case ".zip":
			target, err = unpackZip(path)
			if err != nil {
				slog.Error("failed to unpack zip file", "path", e.Name(), "error", err)
				continue
			}

		case ".gz":
			if filepath.Ext(strings.TrimSuffix(e.Name(), ".gz")) != ".tar" {
				continue
			}
			gzip = true
			fallthrough

		case ".tar":
			target, err = unpackTar(gzip, path)
			if err != nil {
				slog.Error("failedt o unpack tar file", "path", e.Name(), "error", err)
				continue
			}

		default:
			continue
		}

		if target != "" {
			b, err := Load(target)
			if err != nil {
				slog.Error("failed to load bundle", "error", err, "path", e.Name())
				continue
			}

			result = append(result, b)

		}
	}

	return result, nil
}

func Load(path string) (*Bundle, error) {
	pkgjson := filepath.Join(path, "package.json")

	content, err := os.ReadFile(pkgjson)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrInvalidBundle
		}

		return nil, err
	}

	var parsed PackageJSON

	if err := json.Unmarshal(content, &parsed); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidPackageJSON, err.Error())
	}

	// try to find the entrypoint if main is unset
	if parsed.Main == "" {
		if fileExists(path, "index.js") {
			parsed.Main = "index.js"
		}
	} else {
		if !fileExists(path, parsed.Main) {
			return nil, ErrInvalidBundle
		}
	}

	bundle := &Bundle{
		Path: path,
		Main: parsed.Main,
	}

	script, err := os.ReadFile(filepath.Join(path, bundle.Main))
	if err != nil {
		return nil, err
	}

	bundle.ScriptContent = string(script)

	return bundle, nil
}

func fileExists(path string, name string) bool {
	stat, err := os.Stat(filepath.Join(path, name))
	if err != nil {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true
}
