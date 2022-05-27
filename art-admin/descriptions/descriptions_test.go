package descriptions

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestGenDescriptions(t *testing.T) {
	is := is.New(t)

	// Load adjectives
	a, err := os.Open("../etc/adjectives.json")
	is.NoErr(err)

	bs, err := ioutil.ReadAll(a)
	is.NoErr(err)

	adj := adj{}

	err = json.Unmarshal(bs, &adj)
	is.NoErr(err)

	err = json.Unmarshal(bs, &adj)
	is.NoErr(err)

	// Load images
	i, err := os.Open("../etc/mockimages.json")
	is.NoErr(err)

	bs, err = ioutil.ReadAll(i)
	is.NoErr(err)

	mi := mockImages{}

	err = json.Unmarshal(bs, &mi)
	is.NoErr(err)

	totalDescriptions := 10000
	wordsInDescription := 3
	descs := make([]desc, totalDescriptions)

	for i := 0; i < totalDescriptions; i++ {
		for j := 0; j < wordsInDescription; j++ {
			descs[i].Description = append(descs[i].Description, adj.Adjectives[rand.Intn(len(adj.Adjectives))])
			descs[i].MintNumber = i + 1
			descs[i].Image = mi.MockUrls[rand.Intn(len(mi.MockUrls))]
		}
	}

	bs, err = json.Marshal(descs)
	is.NoErr(err)

	f, err := os.Create("../etc/descriptions.json")
	is.NoErr(err)

	_, err = f.WriteString(string(bs))
	is.NoErr(err)
}
