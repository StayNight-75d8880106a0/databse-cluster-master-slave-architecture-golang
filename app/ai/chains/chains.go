package chains

import (
	"context"
	"databse-cluster-master-slave-architecture-golang/app/ai/prompt"
	"databse-cluster-master-slave-architecture-golang/app/ai/vector_store"
	"databse-cluster-master-slave-architecture-golang/app/config/ai_config"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type AI_Chains struct {
	LLM         *ollama.LLM
	VectorStore *vector_store.AI_VectorStore
}

func NewAiChain(vs *vector_store.AI_VectorStore) (*AI_Chains, error) {

	llm, errLLM := ollama.New(
		ollama.WithModel(ai_config.LLM_MODEL),
		ollama.WithServerURL(ai_config.OLLAMA_BASE_URL),
	)

	if errLLM != nil {
		return nil, errLLM
	}

	result := &AI_Chains{
		LLM:         llm,
		VectorStore: vs,
	}

	return result, nil

}

func (a *AI_Chains) GenerateResponse(ctx context.Context, userInput string, streamingFunction func(ctx context.Context, chunk []byte) error) (string, error) {

	searchRes, errSearch := a.VectorStore.Store.SimilaritySearch(ctx, userInput, 11)

	var contextDocs string

	if errSearch == nil {
		for _, doc := range searchRes {
			fmt.Printf("🔍 [SEARCH RESULT] source=%v | content=%s\n",
				doc.Metadata["source"], doc.PageContent[:min(80, len(doc.PageContent))])
			contextDocs += doc.PageContent + "\n"
		}
	}

	fullPrompt := fmt.Sprintf(`Instructions: %s 

		Available Documentation:
		%s

		User Question: %s`,
		prompt.SystemPrompt, contextDocs, userInput,
	)

	chunks, err := a.LLM.GenerateContent(ctx,
		[]llms.MessageContent{
			{
				Role:  llms.ChatMessageTypeHuman,
				Parts: []llms.ContentPart{llms.TextPart(fullPrompt)},
			},
		},
		llms.WithStreamingFunc(streamingFunction),
	)
	if err != nil {
		return "", err
	}

	return chunks.Choices[0].Content, nil

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
