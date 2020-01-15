// COPY FROM: https://go101.org/article/defer-more.html#kind-of-resource-leaking
package memory_leak

import (
	"os"
)

type File struct {
	path    string
	content string
}

func BadManyFiles(files []File) error {
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(file.content)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

func GoodManyFiles(files []File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.path)
			if err != nil {
				return err
			}
			// The close method will be called at
			// the end of the current loop step.
			defer f.Close()

			_, err = f.WriteString(file.content)
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}

	return nil
}
