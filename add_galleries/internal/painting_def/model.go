package paintingdef

type Definition struct {
	Path  string `yaml:"painting"`
	Desc  string `yaml:"painting_desc"`
	Size  string `yaml:"painting_size"`
	Title string `yaml:"painting_title"`
}
