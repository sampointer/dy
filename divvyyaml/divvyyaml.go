package divvyyaml

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Represents a DivvyYaml document and its methods
type DivvyYaml struct {
	Doc string // The constructed YAML document
}

// Parse populates the structure from the given path and options
func (d *DivvyYaml) Parse(path string) error {

	// Store the current working directory
	cwd, err := os.Getwd()
	defer os.Chdir(cwd)
	if err != nil {
		return err
	}

	// Change to the directory given
	path = filepath.Clean(path)
	err = os.Chdir(path)
	if err != nil {
		return err
	}

	// Since our cwd is where we need to be, walk the current directory
	err = filepath.Walk(".", processWalk)
	if err != nil {
		return err
	}

	d.Doc = doc
	return nil
}

var doc string // Constructed YAML document

func processWalk(path string, info os.FileInfo, err error) error {
	// Pass all errors back up the call chain
	if err != nil {
		return err
	}

	// Skip the root directory
	if path == "." {
		return nil
	}

	if info.IsDir() {
		// A directory is a key
		key := filepath.Base(path)
		depth := directoryElements(path)
		doc += indentString(depth, key) + ":\n"
	} else {
		// A file has contents that exist under a key named after the file prefix
		// unless the file name starts with an underscore, in which case no key
		// is written
		var key string
		var depth int

		base := filepath.Base(path)
		if strings.HasPrefix(base, "_") {
			key = ""
			depth = directoryElements(path) - 1
		} else {
			key = strings.TrimSuffix(base, filepath.Ext(path)) + ":\n"
			depth = directoryElements(path)
		}

		doc += indentString(depth, key)

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			doc += indentString(depth+1, scanner.Text()) + "\n"
		}

		if err := scanner.Err(); err != nil {
			return err
		}

	}

	return err
}

// indentString returns the passed string prefixed with the specified number of double spaces
func indentString(indent int, s string) string {
	var rs string

	for i := 0; i < indent; i++ {
		rs += "  "
	}

	rs += s
	return rs
}

// directoryElements returns a count of the number of directories in a path string
func directoryElements(path string) int {
	return len(strings.Split(path, string(os.PathSeparator))) - 1
}
