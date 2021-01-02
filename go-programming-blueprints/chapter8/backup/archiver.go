package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Archiver to archiver
type Archiver interface {
	DestFmt() string
	Archive(src, dest string) error
}

type zipper struct{}

func (z *zipper) DestFmt() string {
	return "%d.zip"
}

func (z *zipper) Archive(src, dest string) error {
	// create folders
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}

	out, err := os.Create(dest) // create new file
	if err != nil {
		return err
	}
	defer out.Close()

	w := zip.NewWriter(out) // to write into the file created
	defer w.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil // skip
		}
		if err != nil {
			return err
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		f, err := w.Create(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}
		return nil
	})
}

// ZIP is an Archiver that zips and unzips files.
var ZIP Archiver = (*zipper)(nil)
