package backup_test

import (
	"log"
	"os"
	"testing"

	"github.com/cheekybits/is"

	backup "."
)

func TestDirHash(t *testing.T) {
	res1, err1 := backup.DirHash("./dirash_test/")

	is := is.New(t)

	is.Equal(err1, nil)

	f, err := os.Create("./dirash_test/file")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("content\n")

	res2, err2 := backup.DirHash("./dirash_test")

	is.NotEqual(res1, res2)
	is.Equal(err2, nil)

	err = os.Remove("./dirash_test/file")
	if err != nil {
		log.Fatal(err)
	}
}
