package vector_store

import (
	"context"
	"databse-cluster-master-slave-architecture-golang/app/config/ai_config"
	"databse-cluster-master-slave-architecture-golang/app/config/db_config"
	"fmt"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

type AI_VectorStore struct {
	Store pgvector.Store
}

func NewVectorStore() (*AI_VectorStore, error) {

	ctx := context.Background()

	llm, errLLM := ollama.New(
		ollama.WithModel(ai_config.EMBEDDING_MODEL),
		ollama.WithServerURL(ai_config.OLLAMA_BASE_URL),
	)

	if errLLM != nil {
		return nil, fmt.Errorf("FAILED TO CREATE OLLAMA EMBEDDING : %v", errLLM)
	}

	embedder, errEmbedder := embeddings.NewEmbedder(llm)
	if errEmbedder != nil {
		return nil, errEmbedder
	}

	database := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().MASTER_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE)

	store, errStore := pgvector.New(
		ctx,
		pgvector.WithConnectionURL(database),
		pgvector.WithEmbedder(embedder),
		pgvector.WithCollectionTableName("ai_knowledge_vectors_langchain"),
	)

	if errStore != nil {
		return nil, errStore
	}

	return &AI_VectorStore{Store: store}, nil

}
