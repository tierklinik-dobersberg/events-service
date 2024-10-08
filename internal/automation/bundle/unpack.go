package bundle

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// TODO(ppacher): for now we ignore any file and directory attributes like ownership and
// permissions
func unpackTar(isGzip bool, path string) (string, error) {
	// try to open the tar[.gz] file
	reader, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open archive: %w", err)
	}
	defer reader.Close()

	// wrap the file stream in a gzip decompressor if isGzip is strue
	var stream io.Reader = reader
	if isGzip {
		uncompressedStream, err := gzip.NewReader(stream)
		if err != nil {
			return "", fmt.Errorf("failed to open gzip stream: %w", err)
		}

		defer uncompressedStream.Close()

		stream = uncompressedStream
	}

	// open the tar file
	tarReader := tar.NewReader(reader)

	// create the base directory
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	dest, err := os.MkdirTemp("", fmt.Sprintf("%s-XXX", base))
	if err != nil {
		return "", err
	}

	// unpack files and directories
L:
	for {
		var header *tar.Header

		header, err = tarReader.Next()
		if err != nil {
			break L
		}

		hd := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err = os.Mkdir(hd, 0755); err != nil {
				err = fmt.Errorf("failed to create directory %q: %w", hd, err)
				break L
			}

		case tar.TypeReg:
			err = writeFile(tarReader, hd)
			if err != nil {
				break L
			}

		default:
			err = fmt.Errorf("unknown tar header type for %q: %v", header.Name, header.Typeflag)
			break L
		}
	}

	if err != nil && !errors.Is(err, io.EOF) {
		defer os.RemoveAll(dest)
		return "", err
	}

	return dest, nil
}

func unpackZip(path string) (string, error) {
	// Open the zip file
	zipFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	// get file stats as we need the file size for the reader
	stat, err := zipFile.Stat()
	if err != nil {
		return "", err
	}

	// create the temporary directory
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	dest, err := os.MkdirTemp("", fmt.Sprintf("%s-XXX", base))
	if err != nil {
		return "", err
	}

	// create a zip reader
	r, err := zip.NewReader(zipFile, stat.Size())
	if err != nil {
		defer os.RemoveAll(dest)
		return "", err
	}

	// finally, actually unpack the files to the temporary directory.
	for _, f := range r.File {
		var err error

		if strings.HasSuffix(f.Name, "/") {
			err = os.MkdirAll(filepath.Join(dest, f.Name), 0o755)
		} else {
			// ensure we create the parent directories as well as not all zip files
			// contain entries for directories.
			err = os.MkdirAll(filepath.Dir(filepath.Join(dest, f.Name)), 0o755)

			// finnally, write the zip file to disk
			if err == nil {
				var fo io.ReadCloser
				fo, err = f.Open()
				if err == nil {
					err = writeFile(fo, filepath.Join(dest, f.Name))
				}
			}
		}

		// in case of an error, remove everything we already unpacked
		// and abort
		if err != nil {
			defer os.RemoveAll(dest)

			return "", fmt.Errorf("failed to open file %q in zip archive: %w", f.Name, err)
		}
	}

	return dest, nil
}

func writeFile(f io.Reader, dst string) error {
	d, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}

	defer d.Close()

	if _, err := io.Copy(d, f); err != nil {
		return fmt.Errorf("failed to write zip file to dist: %w", err)
	}

	return nil
}
