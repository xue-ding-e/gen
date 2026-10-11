package generate

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
)

type XuedingeFileNameFirst struct {
	ID int
}

func (XuedingeFileNameFirst) TableName() string { return "shared_table" }

type XuedingeFileNameSecond struct {
	ID int
}

func (XuedingeFileNameSecond) TableName() string { return "shared_table" }

// Models sharing a table must retain distinct generated filenames in this fork.
func TestXuedingeFileNameUsesStructName(t *testing.T) {
	db, err := gorm.Open(tests.DummyDialector{}, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		value interface{}
		want  string
	}{
		{XuedingeFileNameFirst{}, "XuedingeFileNameFirst"},
		{XuedingeFileNameSecond{}, "XuedingeFileNameSecond"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			meta := QueryStructMeta{db: db}
			if err := meta.parseStruct(tc.value); err != nil {
				t.Fatal(err)
			}
			if meta.FileName != tc.want {
				t.Fatalf("filename = %q, want %q", meta.FileName, tc.want)
			}
			if meta.TableName != "shared_table" {
				t.Fatalf("table name changed to %q", meta.TableName)
			}
		})
	}
}
