package qs

import (
	"net/url"
)

// Decoder is the struct for decoding a URL string.
type Decoder struct {
	pathVals map[string]string
}

// NewDecoder initializes a Decoder with optional url path values.
func NewDecoder(pathVals ...map[string]string) *Decoder {
	dec := &Decoder{}
	if len(pathVals) > 0 {
		dec.pathVals = pathVals[0]
	}
	return dec
}

// Decode decodes a url to a destination struct.
func (d *Decoder) Decode(uri string, dest any) error {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	bind := &DefaultBinder{}
	if d.pathVals != nil {
		err = bind.BindPathParams(d.pathVals, dest)
		if err != nil {
			return err
		}
	}

	err = bind.BindQueryParams(u.Query(), dest)
	if err != nil {
		return err
	}

	return nil
}
