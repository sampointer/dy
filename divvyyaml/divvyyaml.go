package divvyyaml

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// DivvyYaml represents a document and its methods
type DivvyYaml struct {
	Doc string // The constructed YAML document
}

// Parse populates the structure from the given path and options
func (d *DivvyYaml) Parse(path string) error {
	// create a closure containing the d.Doc. This will be passed to filepath.Walk
	walkerfunc := func(path string, info os.FileInfo, err error) error {
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
			d.Doc += indentString(depth, key) + ":\n"
			return nil
		}

		// A file has contents that exist under a key named after the file prefix
		// unless the file name starts with an underscore, in which case no key
		// is written
		var key string
		base := filepath.Base(path)
		depth := directoryElements(path)

		if strings.HasPrefix(base, "_") {
			depth = depth - 1
		} else {
			key = strings.TrimSuffix(base, filepath.Ext(path)) + ":\n"
		}

		d.Doc += indentString(depth, key)

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			d.Doc += indentString(depth+1, scanner.Text()) + "\n"
		}
		file.Close()

		return scanner.Err()
	} // end of inner function definition

	// Store the current working directory so we can get back.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Change to the directory given
	path = filepath.Clean(path)
	err = os.Chdir(path)
	if err != nil {
		return err
	}
	defer os.Chdir(cwd)

	// Since our cwd is where we need to be, walk the current directory
	err = filepath.Walk(".", walkerfunc)
	if err != nil {
		return err
	}

	return nil
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
