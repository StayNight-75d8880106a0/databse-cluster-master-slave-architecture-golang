package ai_registry

import (
	"databse-cluster-master-slave-architecture-golang/app/ai/ai_controller"
	"databse-cluster-master-slave-architecture-golang/app/ai/chains"
	"databse-cluster-master-slave-architecture-golang/app/ai/vector_store"
	"log"

	"gorm.io/gorm"
)

type AIModule struct {
	AI_Controller *ai_controller.AI_Controller
	VectorStore   *vector_store.AI_VectorStore
}

func AI_Registry(db *gorm.DB) *AIModule {

	vectorStore, errVectorStore := vector_store.NewVectorStore()

	if errVectorStore != nil {
		log.Fatalf("Failed to initialize vector store: %v", errVectorStore.Error())
	}

	errVectorStore = vectorStore.LoadKnowledgeBase("app/docs")

	if errVectorStore != nil {
		log.Fatalf("Failed to load knowledge base: %v", errVectorStore.Error())
	}

	ErrLoadDatabaseSnapshot := vectorStore.LoadDatabaseSnapshot(db)

	if ErrLoadDatabaseSnapshot != nil {
		log.Fatalf("Failed to load database snapshot: %v", ErrLoadDatabaseSnapshot.Error())
	}

	chain, errChain := chains.NewAiChain(vectorStore)

	if errChain != nil {
		log.Fatalf("Failed to initialize AI chain: %v", errChain.Error())
	}

	controller := ai_controller.NewAiControllerRegistry(chain)

	return &AIModule{
		AI_Controller: controller,
		VectorStore:   vectorStore,
	}
}
