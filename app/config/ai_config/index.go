package ai_config

import "os"

var OLLAMA_BASE_URL string
var LLM_MODEL string
var EMBEDDING_MODEL string

func AI_Config() {
	OLLAMA_BASE_URL = os.Getenv("OLLAMA_BASE_URL")
	LLM_MODEL = os.Getenv("LLM_MODEL")
	EMBEDDING_MODEL = os.Getenv("EMBEDDING_MODEL")
}
