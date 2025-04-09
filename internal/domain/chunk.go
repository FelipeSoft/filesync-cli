package chunk

import (
	"log"
	"os"
)

type Chunk struct {
	StartBytes int64
	EndBytes   int64
}

func ProcessFileInChunks(filepath string, chunkSizeInMB int64) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	totalChunks := int64(fileSize / chunkSizeInMB)

	for c := range totalChunks {
		offset := chunkSizeInMB * c
		currentChunkSize := fileSize - offset
		if chunkSizeInMB+offset > fileSize {
			currentChunkSize = (chunkSizeInMB + offset) - fileSize
		}

		chunkBytes := make([]byte, currentChunkSize)
		bytes, err := file.ReadAt(chunkBytes, offset)
		if err != nil {
			return err
		}
		log.Print(bytes)
	}

	return nil
}
