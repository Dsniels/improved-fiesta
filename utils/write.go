package utils

import "os"

func WriteFile(filename string, content string) error {

	os.WriteFile(filename, []byte(content), os.ModeAppend.Perm());

	return nil
}
