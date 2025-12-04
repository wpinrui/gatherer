# Gatherer

**Intelligent Information Aggregator for Managing Information Overload**

Gatherer is a desktop application that helps students and professionals capture, process, and query information from multiple sources (emails, PDFs, screenshots, audio recordings, web links) in one unified system with intelligent search and task management.

## The Problem

Information comes at you from everywhere:
- Emails with important attachments
- PDFs scattered across microsites
- Screenshots from video calls with QR codes
- Audio recordings from meetings
- Links buried in chat messages

Current solutions force you to manually organize everything, create duplicates across systems, and make it hard to find what you need when you need it.

## The Solution

Gatherer provides:
- **Universal Capture**: Add information via file upload, email forwarding, or browser extension
- **Intelligent Processing**: Automatic text extraction, OCR, transcription, and summarization
- **Semantic Search**: Find anything with natural language queries
- **Task Extraction**: Automatically identify deadlines and action items
- **Timeline View**: See everything chronologically or by deadline
- **Learning System**: Improves time estimates based on your feedback

## Tech Stack

- **Backend**: Go with Gin framework
- **Database**: PostgreSQL with pgvector for semantic search
- **Frontend**: React with TypeScript
- **AI/ML**: OpenAI API (embeddings, GPT-4, Whisper)
- **OCR**: Tesseract
- **Deployment**: Self-hosted desktop application

## Project Status

🚧 **In Development** - Phase 1: Core Infrastructure

See `proposal.md` for full design document and implementation roadmap.

## Development

This project follows a disciplined development approach with small, focused PRs. See `claude.md` for development guidelines.

### Quick Start

```cmd
# Start PostgreSQL and run the server
dev.bat

# In another terminal, upload a file
curl -X POST http://localhost:8080/upload -F "file=@path\to\document.pdf"

# Query the database
psql.bat
```

**Available endpoints:**
- `GET /health` - Health check
- `POST /upload` - Upload a file (max 50MB)

**Scripts:**
- `dev.bat` - Start PostgreSQL + run server
- `psql.bat` - Open PostgreSQL shell

### Documentation

- **[proposal.md](./proposal.md)**: Complete system design and architecture
- **[claude.md](./claude.md)**: Development workflow and AI agent instructions  
- **[backlog.md](./backlog.md)**: Current tasks and project status


## License

TBD
