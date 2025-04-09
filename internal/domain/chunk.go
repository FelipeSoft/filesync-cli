package chunk

import (
	"log"
	"os"
)

func ProcessFileInChunks(filepath string, chunkSizeInMB int64) (chan []byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	totalChunks := (fileSize / chunkSizeInMB)

	if fileSize%chunkSizeInMB != 0 {
		totalChunks++
	}

	chunks := make(chan []byte, chunkSizeInMB+1)

	go func() {
		defer file.Close()
		defer close(chunks)

		for c := int64(0); c < totalChunks; c++ {
			startOffset := c * chunkSizeInMB
			endOffset := startOffset + chunkSizeInMB

			if endOffset > fileSize {
				endOffset = fileSize
			}

			chunkSize := endOffset - startOffset
			bytes := make([]byte, chunkSize)

			n, err := file.ReadAt(bytes, startOffset)
			if err != nil && err.Error() != "EOF" {
				log.Printf("Error reading chunk %d: %v", c+1, err)
				return
			}

			chunks <- bytes[:n]
		}
	}()

	return chunks, nil
}
