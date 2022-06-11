package bunt

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestMetadata(t *testing.T) {
	is := is.New(t)

	c := Config{
		DBPath: ":memory:",
	}

	bdb, err := c.InitDB()
	is.NoErr(err)

	ms := bdb.MetadataStore()

	for i := 0; i < 100; i++ {
		err := ms.AddOffchainMetadata(fmt.Sprint(i))
		is.NoErr(err)
	}

	all, err := ms.GetAllOffchainMetadata()
	is.NoErr(err)
	is.Equal(len(all), 100)

}
