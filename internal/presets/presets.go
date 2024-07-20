package presets

import (
	"bytes"
	"crypto/rand"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Shutdown
func Shutdown() error {
	return exec.Command("shutdown", "/s", "/f", "/t", "0").Run()
}

// Kill app
func Kill(name string) error {
	return exec.Command("taskkill", "/f", "/im", strings.TrimSpace(name)).Run()
}

// Secure delete
func SecureDelete(path string) error {
	stats, err := os.Stat(path)
	if err != nil {
		return err
	}

	if stats.IsDir() {
		filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				secureDelete(path, int(info.Size()))
			}
			return nil
		})

		var actionOn []string
		filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				actionOn = append(actionOn, path)
			}
			return nil
		})
		for i := 0; i < len(actionOn); i++ {
			renameAndDel(actionOn[len(actionOn)-i-1])
		}
	} else {
		secureDelete(path, int(stats.Size()))
	}

	return nil
}

func secureDelete(file string, size int) {
	for i := 0; i < 4; i++ {
		handle, err := os.OpenFile(file, os.O_WRONLY, 0666)
		if err != nil {
			return
		}

		var written int
		for {
			chunk := bytes.Repeat([]byte(string(3-i)), 10240)
			_, err := handle.Write(chunk)
			written += len(chunk)
			if err != nil {
				handle.Close()
				break
			} else if written >= size {
				handle.Close()
				break
			}
		}
	}
	renameAndDel(file)
}

func renameAndDel(path string) {
	var newFilename string
	sections := strings.Split(path, "/")

	for {
		randBuffer := make([]byte, len(sections[len(sections)-1]))
		rand.Read(randBuffer)
		randBufferStr := strings.ToValidUTF8(string(randBuffer), "a")
		for _, char := range []string{".", "/"} {
			randBufferStr = strings.ReplaceAll(randBufferStr, char, "j")
		}
		err := os.Rename(path, randBufferStr)
		if err == nil {
			newFilename = randBufferStr
			break
		}
	}

	os.RemoveAll(newFilename)
}
