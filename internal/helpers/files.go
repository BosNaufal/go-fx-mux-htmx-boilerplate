package helpers

import (
	"os"
	"path/filepath"

	"github.com/repeale/fp-go"
)

type SimpleDirEntry struct {
	FullPath string
	Name     string
	isDir    bool
}

func osDirEntryToSimpleDirEntry(dir string, entry os.DirEntry) SimpleDirEntry {
	return SimpleDirEntry{
		FullPath: filepath.Join(dir, entry.Name()),
		Name:     entry.Name(),
		isDir:    entry.IsDir(),
	}
}

func readDirSimple(dir string) []SimpleDirEntry {
	osEntries, _ := os.ReadDir(dir)
	simplifyOsDirEntryCurrDir := fp.Curry2(osDirEntryToSimpleDirEntry)(dir)
	return fp.Map(simplifyOsDirEntryCurrDir)(osEntries)
}

func WalkingReadDir(dir string) []SimpleDirEntry {
	simpleEntries := readDirSimple(dir)
	return fp.Reduce(func(acc []SimpleDirEntry, curr SimpleDirEntry) []SimpleDirEntry {
		// Recursive read the dir with output SimpleDirEntry{}
		if curr.isDir {
			nextDirToScan := filepath.Join(dir, curr.Name)
			return append(acc, WalkingReadDir(nextDirToScan)...)
		}

		// Return all the file list if there's no dir.
		return append(acc, curr)
	}, []SimpleDirEntry{})(simpleEntries)
}

func SimpleDirEntryFullPath(dirEntry SimpleDirEntry) string {
	return dirEntry.FullPath
}
