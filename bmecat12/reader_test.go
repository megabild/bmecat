package bmecat12_test

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/megabild/bmecat/bmecat12"
)

type testHandler struct {
	firstPassOnly bool
	header        *bmecat12.Header
	articles      []*bmecat12.Article
	catalogGroups []*bmecat12.CatalogGroup
	featureGroups []*bmecat12.FeatureGroup
}

func (h *testHandler) HandleHeader(header *bmecat12.Header) error {
	h.header = header
	if h.firstPassOnly {
		return io.EOF
	}
	return nil
}

func (h *testHandler) HandleArticle(article *bmecat12.Article) error {
	h.articles = append(h.articles, article)
	return nil
}
func (h *testHandler) HandleCatalogGroup(group *bmecat12.CatalogGroup) error {
	h.catalogGroups = append(h.catalogGroups, group)
	return nil
}
func (h *testHandler) HandleFeatureGroup(feature *bmecat12.FeatureGroup) error {
	h.featureGroups = append(h.featureGroups, feature)
	return nil
}

func TestReadCatalog(t *testing.T) {
	f, err := os.Open(filepath.Join("C:\\", "temp", "bmecat_klein.xml"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	h := &testHandler{}
	r := bmecat12.NewReader(f)
	err = r.Do(context.Background(), h)
	if err != nil {
		t.Fatal(err)
	}
	if h.header == nil {
		t.Fatal("want Header, have nil")
	}
	if want, have := 1, len(h.featureGroups); want != have {
		t.Fatalf("want len(featureGroups) = %d, have %d", want, have)
	}
	if want, have := 338, len(h.catalogGroups); want != have {
		t.Fatalf("want len(catalogGroups) = %d, have %d", want, have)
	}
	if want, have := 255, len(h.articles); want != have {
		t.Fatalf("want len(articles) = %d, have %d", want, have)
	}
}

func TestReadUpdateProducts(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "update_products.golden.xml"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	h := &testHandler{}
	r := bmecat12.NewReader(f)
	err = r.Do(context.Background(), h)
	if err != nil {
		t.Fatal(err)
	}
	if h.header == nil {
		t.Fatal("want Header, have nil")
	}
	if want, have := 2, len(h.articles); want != have {
		t.Fatalf("want len(articles) = %d, have %d", want, have)
	}
}

func TestReadUpdatePrices(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "update_prices.golden.xml"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	h := &testHandler{}
	r := bmecat12.NewReader(f)
	err = r.Do(context.Background(), h)
	if err != nil {
		t.Fatal(err)
	}
	if h.header == nil {
		t.Fatal("want Header, have nil")
	}
	if want, have := 1, len(h.articles); want != have {
		t.Fatalf("want len(articles) = %d, have %d", want, have)
	}
}

func BenchmarkReader(b *testing.B) {
	b.ReportAllocs()

	buf, err := ioutil.ReadFile(filepath.Join("testdata", "update_prices.golden.xml"))
	if err != nil {
		b.Fatal(err)
	}
	buffer := strings.NewReader(string(buf))

	for i := 0; i < b.N; i++ {
		if _, err := buffer.Seek(0, io.SeekStart); err != nil {
			b.Fatal(err)
		}

		h := &testHandler{}
		r := bmecat12.NewReader(buffer)
		err = r.Do(context.Background(), h)
		if err != nil {
			b.Fatal(err)
		}
		if h.header == nil {
			b.Fatal("want Header, have nil")
		}
		if want, have := 1, len(h.articles); want != have {
			b.Fatalf("want len(articles) = %d, have %d", want, have)
		}
	}
}
