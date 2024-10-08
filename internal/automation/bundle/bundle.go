package bundle

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/tierklinik-dobersberg/events-service/internal/automation"
	"github.com/tierklinik-dobersberg/events-service/internal/config"
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

type Log struct {
	Time    time.Time
	Level   slog.Level
	Message string
}

type Bundle struct {
	// Path holds the path to the bundle root
	Path string

	// Main holds the path of the entrypoint script that
	// should be executed.
	Main string

	// Version holds the bundle version.
	Version string

	// License is the SPDX license identifier for the automation bundle.
	License string

	// ScriptContent holds the content of the main entrypoint file.
	ScriptContent string

	lock    sync.Mutex
	runtime *automation.Engine
	logs    []Log
}

// Discover discovers all automation bundles at a specified root.
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
			// if package.json does not exist, check if there's a __single__ directory entry that
			// contains the package

			entries, err := os.ReadDir(path)
			if err != nil {
				return nil, fmt.Errorf("%w: failed to find package.json in %q", ErrInvalidBundle, path)
			}

			if len(entries) != 1 {
				return nil, fmt.Errorf("%w: failed to find package.json in %q but multiple files/directories exists", ErrInvalidBundle, path)
			}

			first := entries[0]
			if !first.IsDir() {
				return nil, fmt.Errorf("%w: found one entry in %q but it's not a directory", ErrInvalidBundle, path)
			}

			return Load(filepath.Join(path, first.Name()))
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
		Path:    path,
		Main:    parsed.Main,
		Version: parsed.Version,
		License: parsed.License,
	}

	script, err := os.ReadFile(filepath.Join(path, bundle.Main))
	if err != nil {
		return nil, err
	}

	bundle.ScriptContent = string(script)

	return bundle, nil
}

// Runtime returns the bundle's automation runtime. This returns nil until bundle.Prepare()
// is called once.
func (bundle *Bundle) Runtime() *automation.Engine {
	bundle.lock.Lock()
	defer bundle.lock.Unlock()

	return bundle.runtime
}

func (bundle *Bundle) ReadLogs(minLevel slog.Level) []Log {
	bundle.lock.Lock()
	defer bundle.lock.Unlock()

	result := make([]Log, 0, len(bundle.logs))

	for _, l := range bundle.logs {
		if l.Level >= minLevel {
			result = append(result, l)
		}
	}

	return result
}

func (bundle *Bundle) Prepare(cfg config.Config, broker automation.Broker, opts ...automation.EngineOption) error {
	bundle.lock.Lock()
	defer bundle.lock.Unlock()

	if bundle.runtime != nil {
		return ErrBundleRuntimePrepared
	}

	// Prepend our own engine options so users defined options may overwrite them
	opts = append([]automation.EngineOption{
		automation.WithConsole(bundle),
	}, opts...)

	runtime, err := automation.New(bundle.Path, cfg, broker, opts...)
	if err != nil {
		return err
	}

	bundle.runtime = runtime

	return nil
}

func (bundle *Bundle) internalLog(lvl slog.Level, msg string) {
	bundle.lock.Lock()
	defer bundle.lock.Unlock()

	bundle.logs = append(bundle.logs, Log{
		Time:    time.Now(),
		Level:   lvl,
		Message: msg,
	})
}

// Log logs an info level message. It implements the console.Printer interface.
func (bundle *Bundle) Log(msg string) {
	bundle.internalLog(slog.LevelInfo, msg)
}

// Warn logs an warn level message. It implements the console.Printer interface.
func (bundle *Bundle) Warn(msg string) {
	bundle.internalLog(slog.LevelWarn, msg)
}

// Error logs an error level message. It implements the console.Printer interface.
func (bundle *Bundle) Error(msg string) {
	bundle.internalLog(slog.LevelError, msg)
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
