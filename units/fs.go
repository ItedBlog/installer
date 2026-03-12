package units

import (
	"io"
	"os"
)

func createDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func createTempDir() (string, error) {
	return os.MkdirTemp("", "itedctl-")
}

func removeDir(path string) error {
	return os.RemoveAll(path)
}

func removeFile(path string) error {
	return os.Remove(path)
}
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func lsDir(path string) ([]os.FileInfo, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	entries, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func writeFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func unzipFile(zipPath, destPath string) error {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		filePath := filepath.Join(destPath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, file.Mode())
		} else {
			os.MkdirAll(filepath.Dir(filePath), file.Mode())
			extractFile := func() error {
				rc, err := file.Open()
				if err != nil {
					return err
				}
				defer rc.Close()

				w, err := os.Create(filePath)
				if err != nil {
					return err
				}
				defer w.Close()

				_, err = io.Copy(w, rc)
				return err
			}
			if err := extractFile(); err != nil {
				return err
			}
		}
	}	
	return nil
}
