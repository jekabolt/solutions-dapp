package descriptions

import (
	"fmt"

	"golang.org/x/exp/rand"
)

type MintDescription interface {
	GetDescriptionOn(sequenceNumber int) (*Description, error)
}

type Config struct {
	CollectionName  string `env:"DESCRIPTIONS_COLLECTION_NAME" envDefault:"Solutions #"`
	CountPerEdition int    `env:"DESCRIPTIONS_COUNT_PER_EDITION" envDefault:"150"`
	RandSeed        int64  `env:"DESCRIPTIONS_RAND_SEED" envDefault:"1337"`
	TotalCount      int
}

func (c *Config) New(total int) MintDescription {
	rand.Seed(uint64(c.RandSeed))
	c.TotalCount = total
	nums := make([]int, c.CountPerEdition*c.TotalCount)
	for i := 0; i < c.CountPerEdition*c.TotalCount; i++ {
		nums[i] = rand.Intn(len(adjectives))
	}

	return &Descriptions{
		adjectives: adjectives,
		c:          c,
		nums:       nums,
	}
}

type Descriptions struct {
	nums       []int
	adjectives []string
	c          *Config
}

type Description struct {
	Description []string `json:"description"`
	MintNumber  int      `json:"mintNumber"`
}

func (d *Description) String() string {
	return fmt.Sprintf("%v", d.Description)
}

func (d *Descriptions) GetDescriptionOn(sequenceNumber int) (*Description, error) {
	if sequenceNumber < 0 || sequenceNumber > d.c.TotalCount {
		return nil, fmt.Errorf("invalid sequence number: %d should be greater than zero and less or equal to total count", sequenceNumber)
	}
	rand.Seed(uint64(d.c.RandSeed))
	desc := &Description{
		Description: []string{},
		MintNumber:  sequenceNumber,
	}

	for i := 1; i <= d.c.CountPerEdition; i++ {
		desc.Description = append(desc.Description, d.adjectives[d.nums[(i*sequenceNumber)-1]])
	}
	return desc, nil
}
