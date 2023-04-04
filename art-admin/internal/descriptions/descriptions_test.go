package descriptions

import (
	"testing"

	"github.com/matryer/is"
)

func TestGenDescriptions(t *testing.T) {
	is := is.New(t)

	c := Config{
		CollectionName:  "test",
		CountPerEdition: 3,
		TotalCount:      100,
		RandSeed:        1337,
	}
	d := c.New(100)

	d1 := []Description{}
	d2 := []Description{}
	for i := 1; i <= 100; i++ {
		desc, err := d.GetDescriptionOn(i)
		is.NoErr(err)
		d1 = append(d1, *desc)
	}

	for i := 1; i <= 100; i++ {
		desc, err := d.GetDescriptionOn(i)
		is.NoErr(err)
		d2 = append(d2, *desc)
	}

	is.Equal(d1, d2)

	_, err := d.GetDescriptionOn(101)
	is.True(err != nil)

	_, err = d.GetDescriptionOn(-1)
	is.True(err != nil)

}
