package repos

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dehwyy/acheron/libraries/go/logg"
)

const (
	playlistFilename = "playlist.m3u8"
)

var (
	ErrFileNotFound = os.ErrNotExist
)

type FileRepository struct {
	Log logg.Logger

	M3u8StreamsDirectory string
}

// Process `FileRepository`: make OS-agnostic path(s).
func (f *FileRepository) prepare() *FileRepository {
	f.M3u8StreamsDirectory = filepath.FromSlash(f.M3u8StreamsDirectory)
	return f
}

// Read .m3u8 playlist as sequence of bytes.
//
// Errors:
//   - `FileNotFoundError`: playlist file does not exist.
//   - `FileError`: fallback error.
func (f *FileRepository) ReadM3u8Playlist(streamName string) ([]byte, error) {
	path := filepath.Join(f.M3u8StreamsDirectory, streamName, playlistFilename)

	playlistFiledata, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("playlist file not found: %w", err)
		}
		return nil, fmt.Errorf("failed to read playlist file: %w", err)
	}

	return playlistFiledata, nil
}

// Read .ts segment as sequence of bytes.
//
// Errors:
//   - `FileNotFoundError`: segment file does not exist.
//   - `FileError`: fallback error.
func (f *FileRepository) ReadSegment(streamName string, segmentName string) ([]byte, error) {
	path := filepath.Join(f.M3u8StreamsDirectory, streamName, segmentName)

	segmentFiledata, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("segment file not found: %w", err)
		}
		return nil, fmt.Errorf("failed to read segment file: %w", err)
	}

	return segmentFiledata, nil
}
