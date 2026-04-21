#!/bin/bash

# Ollama image berbasis ubuntu, pastikan curl ada
apt-get update -qq && apt-get install -y -qq curl 2>/dev/null || true

# Start ollama server di background
ollama serve &
OLLAMA_PID=$!

# Tunggu server ready
echo "⏳ Waiting for ollama server to start..."
until curl -s http://localhost:11434/ > /dev/null 2>&1; do
    sleep 2
done
echo "✅ Ollama server is up."

# Pull model
echo "📦 Pulling llama3.2:3b..."
ollama pull llama3.2:3b

echo "📦 Pulling nomic-embed-text..."
ollama pull nomic-embed-text

echo "✅ All models ready."
touch /tmp/ollama_ready
echo "✅ Marker /tmp/ollama_ready created."

wait $OLLAMA_PID