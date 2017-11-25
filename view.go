package temp

import "io/ioutil"

type View struct {
	Path string
}

// Parse a view.
func (v View) Parse() ([]byte, error) {
	html, err := ioutil.ReadFile(v.Path)

	return html, err
}
