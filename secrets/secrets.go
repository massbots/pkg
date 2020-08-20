package secrets

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Load walks through the Swarm secrets directory and stores
// them as environment variables if such presented.
func Load() error {
	return filepath.Walk("/run/secrets",
		func(path string, fi os.FileInfo, _ error) error {
			if fi == nil || fi.IsDir() {
				return nil
			}

			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			return os.Setenv(fi.Name(), string(data))
		})
}