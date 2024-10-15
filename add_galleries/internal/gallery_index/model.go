package galleryindex

import paintingdef "add_galleries/internal/painting_def"

type GalleryIndex struct {
	GeneralDescription   string
	GeneralTitle         string
	GeneralWeight        int
	GeneralFeaturedImage string
	Definitions          []*paintingdef.Definition
}
