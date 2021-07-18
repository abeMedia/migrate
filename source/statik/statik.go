package statik

import (
	nurl "net/url"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/rakyll/statik/fs"
)

func init() {
	source.Register("file", &Statik{})
}

type Statik struct {
	httpfs.PartialDriver
	url  string
	path string
}

func (f *Statik) Open(url string) (source.Driver, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil, err
	}

	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	nf := &Statik{
		url:  url,
		path: u.Path,
	}
	if err := nf.Init(statikFS, "p"); err != nil {
		return nil, err
	}
	return nf, nil
}
