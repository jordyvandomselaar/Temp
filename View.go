package Temp

import "io/ioutil"

type View struct {
	Path string
}

// Parse a view.
func (v View) Parse() []byte {
	html, err := ioutil.ReadFile(v.Path)

	if err != nil {
		panic(err)
	}

	return html
}
