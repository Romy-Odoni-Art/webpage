package galleryindex

import (
	"io"
	"text/template"
)

func (g *GalleryIndex) Write(out io.Writer, tmpFileContent string) error {
	tmpl, err := template.New("t1").Parse(tmpFileContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(out, g)
}
