package qs

import (
	"net/url"
	"slices"
	"testing"
)

const urlq = `/indexes/default?searchableAttributes=title&attributesForFaceting=tags&attributesForFaceting=authors&attributesForFaceting=series&attributesForFaceting=narrators`

type params struct {
	SrchAttr  []string `query:"searchableAttributes"`
	FacetAttr []string `query:"attributesForFaceting"`
	Index     string   `path:"index"`
}

var testParams = &params{
	SrchAttr:  []string{"title"},
	FacetAttr: []string{"tags", "authors", "series", "narrators"},
	Index:     "default",
}

func TestDecoder(t *testing.T) {
	pv := map[string]string{
		"index": "default",
	}
	dest := params{}
	dec := NewDecoder(pv)
	err := dec.Decode(urlq, &dest)
	//q := parsed()
	//err := Decode(q, &p)
	if err != nil {
		t.Error(err)
	}
	sw := []string{"title"}
	if !slices.Equal(dest.SrchAttr, sw) {
		t.Errorf("got %v, expected %v\n", dest.SrchAttr, sw)
	}
	facets := []string{"tags", "authors", "series", "narrators"}
	if !slices.Equal(dest.FacetAttr, facets) {
		t.Errorf("got %v, expected %v\n", dest.FacetAttr, facets)
	}
	i := "default"
	if dest.Index != i {
		t.Errorf("got %v, expected %v\n", dest.Index, i)
	}
}

//func TestMarshal(t *testing.T) {
//  v, err := Encode(testParams)
//  if err != nil {
//    t.Error(err)
//  }
//  //fmt.Printf("%#v\n", v)
//  sw := []string{"title"}
//  if !slices.Equal(v["searchableAttributes"], sw) {
//    t.Errorf("got %v, expected %v\n", v["searchableAttributes"], sw)
//  }
//  facets := []string{"tags", "authors", "series", "narrators"}
//  if !slices.Equal(v["attributesForFaceting"], facets) {
//    t.Errorf("got %v, expected %v\n", v["attributesForFaceting"], facets)
//  }
//  i := []string{"default"}
//  if !slices.Equal(v["index"], i) {
//    t.Errorf("got %v, expected %v\n", v["index"], i)
//  }
//}

func parsed() url.Values {
	v, _ := url.ParseQuery(urlq)
	return v
}
