package main

import (
	galleryindex "add_galleries/internal/gallery_index"
	paintingdef "add_galleries/internal/painting_def"
	_ "embed"
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed template/index.md.tmpl
var tmpFile string

type inputParams struct {
	searchDir   string
	title       string
	description string
	weight      int
	outFile     string
}

func getParams() (*inputParams, error) {
	title := flag.String("title", "", "Title of the gallery")
	desc := flag.String("description", "", "Description text for the gallery")
	searchDir := flag.String("dir", "", "Content directory to search for paintings")
	weight := flag.Int("weight", 0, "Weight of the gallery to determine order")
	outFile := flag.String("out", "", "Name of the file to output")
	flag.Parse()
	out := &inputParams{
		searchDir:   *searchDir,
		title:       *title,
		description: *desc,
		weight:      *weight,
		outFile:     *outFile,
	}
	if out.searchDir == "" || out.title == "" || out.weight == 0 || out.outFile == "" {
		flag.Usage()
		return nil, errors.New("required parameter was not set. Required are title, dir, weight, featured and out")
	}
	return out, nil
}

func main() {
	inputParams, err := getParams()
	if err != nil {
		log.Fatal(err)
	}
	files, err := os.ReadDir(inputParams.searchDir)
	if err != nil {
		log.Fatal(err)
	}
	defs := []*paintingdef.Definition{}
	for _, f := range files {

		if strings.HasPrefix(f.Name(), "painting") && filepath.Ext(f.Name()) == ".md" {
			def, err := paintingdef.ParseDefFromFile(filepath.Join(inputParams.searchDir, f.Name()))
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Adding painting " + def.Title)
			defs = append(defs, def)
		}
	}
	if len(defs) > 0 {
		index := &galleryindex.GalleryIndex{
			GeneralDescription: inputParams.description,
			GeneralTitle:       inputParams.title,
			GeneralWeight:      inputParams.weight,
			Definitions:        defs,
		}
		fOut, err := os.Create(inputParams.outFile)
		if err != nil {
			log.Fatal(err)
		}
		err = index.Write(fOut, tmpFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("No panting files found, skipping")
	}
}
