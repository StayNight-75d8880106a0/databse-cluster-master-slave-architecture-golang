package vector_store

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/langchaingo/textsplitter"
)

func (v *AI_VectorStore) LoadKnowledgeBase(docsPath string) error {

	ctx := context.Background()

	files, errFiles := os.ReadDir(docsPath)

	if errFiles != nil {
		return errFiles
	}

	splitter := textsplitter.NewRecursiveCharacter()
	splitter.ChunkSize = 1001
	splitter.ChunkOverlap = 201

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			path := filepath.Join(docsPath, file.Name())
			content, _ := os.ReadFile(path)

			hash := md5.Sum(content)
			fileHash := hex.EncodeToString(hash[:])

			docs, err := textsplitter.CreateDocuments(splitter, []string{string(content)}, []map[string]any{
				{"source": file.Name(), "hash": fileHash},
			})

			if err != nil {
				continue
			}

			_, err = v.Store.AddDocuments(ctx, docs)
			if err != nil {
				fmt.Printf("Failed to save the %s document: %v\n", file.Name(), err)
			}
		}
	}

	fmt.Println("✅ Knowledge base updated successfully!")
	return nil

}
