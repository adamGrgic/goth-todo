package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func ComputeFileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil))[:8], nil
}
