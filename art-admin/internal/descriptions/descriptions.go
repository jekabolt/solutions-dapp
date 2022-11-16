package descriptions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Config struct {
	Path            string `env:"DESCRIPTIONS_PATH" envDefault:"etc/descriptions.json"`
	CollectionName  string `env:"DESCRIPTIONS_COLLECTION_NAME" envDefault:"Solutions #"`
	CountPerEdition int    `env:"DESCRIPTIONS_COUNT_PER_EDITION" envDefault:"150"`
}

type Store struct {
	m map[int]desc
	c *Config
}

type adj struct {
	Adjectives []string `json:"adjectives"`
}

type mockImages struct {
	MockUrls []string `json:"mockUrls"`
}

type desc struct {
	Description []string `json:"description"`
	MintNumber  int      `json:"mintNumber"`
	Image       string   `json:"image"`
}

func (c *Config) Init() (*Store, error) {
	s := &Store{
		m: make(map[int]desc),
		c: c,
	}
	return s, s.InitialUpload()
}

func (s *Store) InitialUpload() error {
	bs, err := ioutil.ReadFile(s.c.Path)
	if err != nil {
		return err
	}
	descs := []desc{}
	err = json.Unmarshal(bs, &descs)
	if err != nil {
		return err
	}

	for i, d := range descs {
		s.m[i] = d
	}
	return nil
}

func (s *Store) GetCollectionName(number int) string {
	return fmt.Sprintf("%s%d", s.c.CollectionName, number)
}

func (s *Store) GetDescription(sequenceNumber int) string {
	return strings.Join(s.m[sequenceNumber].Description, ",")
}

func (s *Store) GetImage(sequenceNumber int) string {
	return s.m[sequenceNumber].Image
}

// TODO: GetEditionBySequence
func (s *Store) GetEditionBySequence(sequenceNumber int) int32 {
	// if dbz := sequenceNumber < 0; dbz {
	// 	return -1
	// }
	return 1
}
