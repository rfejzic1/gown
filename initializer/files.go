package initializer

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"strings"
	"text/template"
)

//go:embed all:files
var files embed.FS
var temp *template.Template = template.Must(template.ParseFS(files, "files/*"))

type walkFilesFunc func(fileName string, content []byte) error

// Walks the entire 'files' directory, rendering all the files with the given 'data'.
// To embed all files from the 'files' directory, we need to prefix the `go.mod` file
// to avoid go:embed from detecting it as a separate module. The `fileName` parameter
// is the unprefiexed file name (without the '__').
func renderFiles(data interface{}, f walkFilesFunc) error {
	return fs.WalkDir(files, "files", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		fileName := d.Name()

		content := &bytes.Buffer{}

		if err := renderFile(content, fileName, data); err != nil {
			return err
		}

		cleanFileName := strings.TrimLeft(fileName, "__")

		return f(cleanFileName, content.Bytes())
	})
}

// Renders a file as template from the 'initializer/files' directory
func renderFile(w io.Writer, fileName string, data interface{}) error {
	return temp.ExecuteTemplate(w, fileName, data)
}
