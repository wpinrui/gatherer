# Gatherer - Intelligent Information Aggregator

## Overview

Gatherer is an intelligent information aggregation and timeline management system designed for students and professionals dealing with information overload from multiple sources. It provides a unified interface to capture, process, query, and act on information regardless of source format.

## Problem Statement

Students and professionals receive critical information through fragmented channels:
- Emails with attachments and embedded information
- PDFs on various microsites
- Screenshots from video calls containing QR codes or important visuals
- Audio recordings from meetings
- Web links scattered across platforms

Current solutions require manual organization, duplicate information exists across sources, and retrieving information chronologically or by relevance is difficult.

## Solution

A desktop-first application that:
1. Accepts information from multiple sources with minimal friction
2. Intelligently processes and indexes all content types
3. Provides semantic and temporal querying
4. Extracts and manages tasks and deadlines automatically
5. Learns from user feedback to improve time estimations
6. Eliminates duplicate information intelligently

## Architecture

### Technology Stack

**Backend:**
- Language: Go 1.21+
- Framework: Gin or Echo for HTTP API
- Database: PostgreSQL with pgvector extension
- Queue: Redis for background job processing
- Storage: Local filesystem with organized structure

**Frontend (Electron + React):**
- Shell: Electron for desktop packaging
- Framework: React 18+ with TypeScript
- State Management: React Query + Zustand
- UI Components: Tailwind CSS + shadcn/ui
- Timeline: react-calendar-timeline or custom implementation
- Document Viewer: PDF.js, image viewer, audio player
- Native Integration: Clipboard API, file system access via Electron

**External Services:**
- OpenAI API: Text embeddings, GPT-4 for summarization and extraction
- Whisper API: Audio transcription
- Tesseract OCR: Image text extraction (self-hosted via gosseract)

### System Components

#### 1. Electron Application Shell
- Self-contained Electron desktop application
- Spawns Docker container (PostgreSQL + pgvector)
- Spawns Go backend server
- Embeds React frontend in Electron renderer process
- Manages all process lifecycles (start/stop/restart)
- System tray integration for background operation
- Native OS integration (clipboard, file system, notifications)

#### 2. Ingestion Layer

**Entry Points:**
- **File Upload**: Direct file upload via drag-and-drop in the app
- **Clipboard Paste**: Paste anything from clipboard (images, text, URLs, files)
- **Folder Watching**: Monitor user-configured folder for automatic import
- **URL Submission**: Paste URL to fetch and archive web content

**Supported Formats:**
- Documents: PDF, DOCX, TXT, MD
- Images: PNG, JPG, JPEG, WebP (with OCR)
- Audio: MP3, WAV, M4A, OGG
- Video: MP4, WebM (extract audio for transcription)
- Web: HTML snapshots, bookmarks

#### 3. Processing Pipeline

Each ingested item goes through a processing pipeline:

```
Ingestion → Storage → Processing → Embedding → Indexing
```

**Processing Steps:**

1. **Storage**
   - Save original file to organized directory structure
   - Generate unique identifier (UUID)
   - Create database record with metadata

2. **Content Extraction**
   - PDF: Extract text using pdfcpu/unidoc
   - Images: OCR via Tesseract, QR code detection via gozxing
   - Audio: Transcribe via Whisper API
   - Web: Extract main content, strip ads/navigation
   - DOCX: Extract text and formatting

3. **Intelligence Extraction** (via LLM)
   - Summarization: Generate concise summary
   - Task Extraction: Identify action items, deadlines, requirements
   - Date Parsing: Extract all mentioned dates and deadlines
   - Entity Recognition: People, places, topics
   - Categorization: Automatic tagging based on content

4. **Embedding Generation**
   - Split content into chunks (500-1000 tokens)
   - Generate embeddings via OpenAI API
   - Store in pgvector for semantic search

5. **Deduplication**
   - Compare embeddings with existing content
   - Flag similar items (>0.85 similarity)
   - Suggest keeping best version (prefer: direct PDF > screenshot > email forward)

#### 4. Database Schema

**Core Tables:**

```sql
-- Items: All ingested content
items (
  id UUID PRIMARY KEY,
  title TEXT,
  source_type TEXT, -- email, upload, extension, url
  content_type TEXT, -- pdf, image, audio, web
  original_filename TEXT,
  file_path TEXT,
  ingested_at TIMESTAMP,
  status TEXT, -- pending, processing, completed, error
  metadata JSONB
)

-- Processed content
content (
  id UUID PRIMARY KEY,
  item_id UUID REFERENCES items(id),
  extracted_text TEXT,
  summary TEXT,
  processed_at TIMESTAMP
)

-- Embeddings for semantic search
embeddings (
  id UUID PRIMARY KEY,
  item_id UUID REFERENCES items(id),
  chunk_index INTEGER,
  chunk_text TEXT,
  embedding vector(1536), -- OpenAI embedding size
  created_at TIMESTAMP
)

-- Extracted tasks and deadlines
tasks (
  id UUID PRIMARY KEY,
  item_id UUID REFERENCES items(id),
  description TEXT,
  deadline TIMESTAMP,
  estimated_duration_minutes INTEGER,
  actual_duration_minutes INTEGER, -- user feedback
  priority TEXT, -- high, medium, low
  status TEXT, -- pending, in_progress, completed
  created_at TIMESTAMP
)

-- Tags and categories
tags (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE,
  color TEXT
)

item_tags (
  item_id UUID REFERENCES items(id),
  tag_id UUID REFERENCES tags(id),
  PRIMARY KEY (item_id, tag_id)
)

-- Duplicates tracking
duplicates (
  id UUID PRIMARY KEY,
  item_id_1 UUID REFERENCES items(id),
  item_id_2 UUID REFERENCES items(id),
  similarity_score FLOAT,
  preferred_item_id UUID, -- user choice
  created_at TIMESTAMP
)

-- User feedback for time estimation learning
time_feedback (
  id UUID PRIMARY KEY,
  task_id UUID REFERENCES tasks(id),
  estimated_minutes INTEGER,
  actual_minutes INTEGER,
  task_category TEXT,
  recorded_at TIMESTAMP
)
```

#### 5. Query Engine

**Search Capabilities:**

1. **Semantic Search**
   - Natural language queries: "What do I need to do this week?"
   - Embedding-based similarity search
   - Hybrid search combining keyword and semantic

2. **Temporal Queries**
   - Timeline view: See items chronologically
   - Deadline filtering: "Next 2 days", "This week"
   - Date range queries

3. **Filtered Search**
   - By source type (email, PDF, screenshot)
   - By content type (documents, audio, images)
   - By tags or categories
   - By completion status

4. **Smart Queries**
   - "What do I need to prepare for tomorrow?"
   - "Show me all information about [topic]"
   - "What was mentioned in last week's meetings?"

**Query Processing:**
```
User Query → Intent Classification → Query Expansion → 
Vector Search + SQL Filters → Ranking → Results
```

#### 6. Task Management

**Task Extraction:**
- LLM identifies action items from content
- Extracts: description, deadline, priority signals
- Links to source document

**Time Estimation:**
- Initial: Rule-based heuristics (reading time, complexity signals)
- Learning: Track actual vs estimated time
- Categories: reading, assignment, preparation, admin
- Adjust estimates based on historical accuracy

**Task Workflow:**
- Auto-create from extracted information
- User can edit, split, merge tasks
- Mark complete with actual time spent
- System learns from feedback

**Proactive Suggestions:**
- "You have 3 items due in 2 days, estimated 4 hours total"
- "Based on past tasks, you might want to start this 2 days early"
- "This overlaps with your existing commitments on [date]"

#### 7. Frontend Interface

**Main Views:**

1. **Dashboard**
   - Upcoming deadlines widget
   - Recent items
   - Quick search bar
   - Suggested actions ("Review these 3 items for tomorrow")

2. **Timeline View**
   - Chronological display of all items
   - Filter by date range
   - Visual deadline indicators
   - Drag items to adjust dates

3. **Search Interface**
   - Natural language search
   - Advanced filters panel
   - Results grouped by relevance/date
   - Preview pane for quick viewing

4. **Item Detail View**
   - Full content display (PDF viewer, audio player, image viewer)
   - Extracted summary
   - Associated tasks
   - Related items (based on similarity)
   - Source information
   - Edit metadata, add notes

5. **Tasks View**
   - Kanban board or list view
   - Filter by deadline, priority, status
   - Time tracking interface
   - Bulk actions

6. **Settings**
   - API key configuration (OpenAI)
   - Watched folder configuration
   - Clipboard monitoring preferences
   - Data management (export, backup)

#### 8. Clipboard Integration

**Capabilities:**
- Monitor clipboard for new content (opt-in)
- Manual paste via Ctrl+V or button in app
- Support all clipboard content types:
  - Images (screenshots, copied images)
  - Text (plain text, rich text)
  - URLs (automatically fetch and archive)
  - Files (copied from Explorer)
- Automatic content type detection
- Duplicate detection before import

**Implementation:**
- Electron's clipboard API for cross-platform support
- File handling via Node.js fs module
- Image processing via sharp or native APIs

#### 9. Folder Watching

**Capabilities:**
- User-configurable watched folder (e.g., `~/Gatherer/inbox/`)
- Automatic import when files are added
- Support for all file types in Supported Formats
- Option to move or copy files after import
- Configurable polling interval or native fs events

**Implementation:**
- Node.js `chokidar` for cross-platform file watching
- Queue new files for processing pipeline
- Handle rapid file additions gracefully

## Implementation Phases

### Phase 1: Core Infrastructure (Weeks 1-2)
**Goal:** Basic application with file upload and viewing

**Tasks:**
1. Project setup (Go modules, directory structure, Git)
2. Database schema and migrations
3. Basic Gin API server with health check
4. File upload endpoint
5. PostgreSQL connection and basic CRUD
6. React app scaffold with TypeScript
7. File upload UI component
8. Items list view
9. Simple file viewer (PDF, image, text)
10. Desktop launcher script (auto-start backend)

**Deliverables:**
- Can upload PDF, image, text files
- View uploaded files in list
- Click to view content
- Application launches with one command

### Phase 2: Content Processing (Weeks 3-4)
**Goal:** Extract and process content from files

**Tasks:**
1. PDF text extraction (pdfcpu)
2. Image OCR integration (Tesseract)
3. Background job queue setup (Redis)
4. Processing worker implementation
5. LLM integration for summarization
6. Content storage and retrieval
7. Processing status tracking
8. Error handling and retries
9. Display extracted text and summary in UI

**Deliverables:**
- Uploaded files automatically processed
- Text extracted from PDFs and images
- Summaries generated
- View processing status

### Phase 3: Search Foundation (Weeks 5-6)
**Goal:** Basic search with embeddings

**Tasks:**
1. Install pgvector extension
2. Embedding generation via OpenAI
3. Content chunking strategy
4. Vector storage and indexing
5. Basic semantic search endpoint
6. Search UI component
7. Results display with relevance
8. Keyword search fallback

**Deliverables:**
- Search across all content
- Semantic understanding of queries
- Ranked results display

### Phase 4: Task Extraction (Weeks 7-8)
**Goal:** Automatic task and deadline extraction

**Tasks:**
1. LLM prompt for task extraction
2. Task database schema
3. Task extraction worker
4. Date parsing and normalization
5. Task CRUD endpoints
6. Tasks view UI
7. Timeline visualization
8. Task editing and completion

**Deliverables:**
- Tasks auto-extracted from content
- View all tasks with deadlines
- Mark tasks complete
- Timeline showing upcoming items

### Phase 5: Electron Shell (Weeks 9-10)
**Goal:** Package app as self-contained Electron application

**Tasks:**
1. Electron project setup with electron-builder
2. Main process to spawn Docker container
3. Main process to spawn Go backend
4. IPC communication between main and renderer
5. Process health monitoring and restart
6. System tray integration
7. Window state persistence
8. App packaging and distribution setup

**Deliverables:**
- Single executable launches entire stack
- Docker + Go backend auto-start
- System tray for background operation
- Proper shutdown handling

### Phase 6: Audio Processing (Weeks 11-12)
**Goal:** Transcribe and process audio content

**Tasks:**
1. Audio upload support
2. Whisper API integration
3. Audio transcription worker
4. Audio player UI component
5. Transcription display
6. Search transcriptions
7. Support video files (extract audio)

**Deliverables:**
- Upload audio/video files
- Automatic transcription
- Search audio content by transcript
- Play audio with transcript

### Phase 7: Clipboard & Folder Watching (Weeks 13-14)
**Goal:** Easy content capture via clipboard and folder monitoring

**Tasks:**
1. Clipboard paste handler (Ctrl+V in app)
2. Image clipboard support (screenshots)
3. Text/URL clipboard support
4. File clipboard support (Explorer copy)
5. Folder watcher setup with chokidar
6. Settings UI for watched folder path
7. Import queue for folder additions
8. Duplicate detection on import

**Deliverables:**
- Paste anything into app from clipboard
- Configure watched folder in settings
- Files auto-import when added to folder
- Duplicates detected before import

### Phase 8: Intelligence & Learning (Weeks 15-16)
**Goal:** Smart features and learning system

**Tasks:**
1. Deduplication algorithm
2. Similarity detection
3. Time estimation heuristics
4. Time tracking for tasks
5. Learning model for time estimation
6. Feedback loop implementation
7. Proactive suggestions engine
8. Smart query understanding

**Deliverables:**
- Detect duplicate content
- Suggest keeping best version
- Time estimates for tasks
- Improve estimates from feedback
- Smart suggestions ("Do this next")

### Phase 9: Advanced Features (Weeks 17-18)
**Goal:** Polish and advanced capabilities

**Tasks:**
1. Advanced filtering in search
2. Custom tags and categories
3. Manual categorization
4. Export functionality
5. Backup and restore
6. Settings management UI
7. Keyboard shortcuts
8. Dark mode

**Deliverables:**
- Rich filtering options
- Tag and organize content
- Export data
- Backup system
- Polished UX

### Phase 10: Optimization & Deployment (Weeks 19-20)
**Goal:** Performance and distribution

**Tasks:**
1. Query optimization
2. Caching strategy
3. Background processing optimization
4. Frontend performance (lazy loading, virtualization)
5. Build pipeline
6. Installer creation (Windows/Mac/Linux)
7. Documentation
8. Demo video and screenshots

**Deliverables:**
- Fast search and UI
- Production-ready build
- Installable application
- User documentation
- Portfolio-ready presentation

## Technical Considerations

### Self-Hosting & Local Development

**Single Computer Setup:**
- PostgreSQL: Docker container managed by Electron
- Redis: Docker container managed by Electron
- Backend: Go server spawned by Electron main process
- Frontend: React app in Electron renderer process

**Application Architecture:**
```
Gatherer.exe
    └── Electron Main Process
            ├── Spawns: docker-compose up (PostgreSQL + Redis)
            ├── Spawns: Go backend server
            ├── Manages: Process lifecycle
            └── Renders: React frontend (renderer process)
```

### Data Storage

**File Organization:**
```
~/Gatherer/
  data/
    files/
      2024/
        12/
          [uuid].pdf
          [uuid].png
    database/
      postgres/ (if embedded)
    logs/
  config/
    config.yaml
```

### API Design

**RESTful Endpoints:**
```
POST   /api/items              # Upload item
GET    /api/items              # List items
GET    /api/items/:id          # Get item details
DELETE /api/items/:id          # Delete item

POST   /api/search             # Search query
GET    /api/search/suggestions # Autocomplete

GET    /api/tasks              # List tasks
POST   /api/tasks              # Create task
PUT    /api/tasks/:id          # Update task
DELETE /api/tasks/:id          # Delete task

POST   /api/process/:id        # Trigger reprocessing
GET    /api/process/:id/status # Processing status

POST   /api/clipboard          # Import from clipboard
POST   /api/folder/watch       # Configure watched folder
GET    /api/folder/watch       # Get watched folder config

GET    /api/timeline           # Timeline data
GET    /api/duplicates         # Duplicate detection results
```

### Security Considerations

**Local-First Security:**
- No external authentication needed (local user)
- API keys stored securely in config
- CORS restricted to localhost
- All communication via localhost only

**Data Privacy:**
- All data stays on user's machine
- Cloud APIs only for processing (OpenAI, Whisper)
- Option to use local models (future: Ollama integration)

### Performance Targets

- Upload processing: Start within 1 second
- Search latency: <200ms for semantic search
- UI responsiveness: 60fps scrolling
- Large file handling: 100MB+ PDFs without blocking
- Concurrent processing: Handle 10+ items simultaneously

### Future Enhancements (Post-MVP)

1. **Mobile Companion App**: Quick capture from phone
2. **Local LLM Support**: Ollama for offline processing
3. **Collaboration**: Share items with others
4. **Cloud Sync**: Optional sync across devices
5. **Advanced Analytics**: Time spent, productivity insights
6. **Integration API**: Connect to calendar, todo apps
7. **Voice Input**: Dictate notes directly
8. **Smart Notifications**: Remind about upcoming deadlines
9. **Browser Extension**: Optional extension for web content capture
10. **Spaced Repetition**: Study mode for educational content

## Success Metrics

**Portfolio Perspective:**
- Demonstrates full-stack capability (Go, React, TypeScript, Electron)
- Shows ML/AI integration (embeddings, LLM, OCR, speech-to-text)
- Complex system design (multi-stage pipeline, async processing)
- Desktop application with process orchestration
- Production-quality code and architecture

**User Perspective:**
- Reduces time to find information by 80%
- Eliminates duplicate information
- Never miss a deadline
- Spend less time organizing, more time acting
- Feel in control of information flow

## Project Timeline

**Total Duration:** 20 weeks (~5 months)
**Effort:** Part-time (15-20 hours/week)

This is an aggressive but achievable timeline for a motivated developer with code agent assistance. Adjust phases as needed based on actual progress.

## Conclusion

Gatherer solves a real problem for a niche audience (students, professionals with information overload) while showcasing impressive technical skills. The combination of multi-modal processing, intelligent search, and learning systems makes this a standout portfolio project.

The self-hosted, local-first approach ensures privacy and control while the gradual enhancement path allows for continuous improvement and learning throughout development.
