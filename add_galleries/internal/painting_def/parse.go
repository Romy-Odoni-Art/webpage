package paintingdef

import (
	"bufio"
	"os"

	"gopkg.in/yaml.v3"
)

const separator = "---"

func ParseDefFromFile(path string) (*Definition, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fmOpen := false
	yamlBytes := []byte{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		if t == separator {
			if fmOpen {
				break
			} else {
				fmOpen = true
				continue
			}
		}
		if fmOpen {
			yamlBytes = append(yamlBytes, sc.Bytes()...)
			yamlBytes = append(yamlBytes, byte('\n'))
		}
	}
	f.Close()
	if err := sc.Err(); err != nil {
		return nil, err
	}
	var result Definition
	err = yaml.Unmarshal(yamlBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
