# Gatherer - Intelligent Information Aggregator

## Overview

Gatherer is an intelligent information aggregation and timeline management system designed for students and professionals dealing with information overload from multiple sources. It provides a unified interface to capture, process, query, and act on information regardless of source format.

## Problem Statement

Students and professionals receive critical information through fragmented channels:
- PDFs scattered across learning portals and microsites
- Screenshots from video calls containing important visuals
- Photos of whiteboards and slides taken on phones
- Audio recordings from meetings and lectures
- Quick notes jotted down without context
- Web links saved "for later" and forgotten
- Files saved in random folders across the system

Current solutions require manual organization, force users to tag and categorize everything themselves, and make it nearly impossible to find information when you need it. The cognitive overhead of organizing information often exceeds the effort of just doing the work.

## Solution

A **smart knowledge base** that builds itself. Gatherer is not just storage with search—it's an intelligent system that:

1. **Captures with zero friction**: Paste from clipboard, drop files in a folder, drag-and-drop—content flows in effortlessly
2. **Understands context automatically**: Analyzes content to infer module codes, topics, and relationships without manual tagging
3. **Prompts only when necessary**: When context is ambiguous, asks for minimal input (one click to confirm a suggestion)
4. **Connects related information**: Links items by topic, time, source, and semantic similarity
5. **Extracts actionable intelligence**: Identifies tasks, deadlines, and priorities from content
6. **Learns and improves**: Gets smarter about context inference based on user corrections

The core principle: **Gatherer does the organizing so you don't have to.**

---

## User Scenario: A Day in the Life

*Meet Alex, a university student taking CS2030S (Programming Methodology II), CS2040S (Data Structures), and IS1108 (Digital Ethics). Here's how Gatherer fits into a typical day.*

### Morning: Online Lecture (CS2030S)

**9:00 AM** — Alex joins a Zoom lecture. The professor shares slides about Java Streams.

- Alex takes a screenshot of a complex diagram → **Win+Shift+S, then Ctrl+V into Gatherer**
- A toast notification appears: *"Screenshot added. CS2030S?"* (Gatherer recognized "Stream" and recent CS2030S activity)
- Alex clicks **✓** to confirm. Done in 1 second.

**9:45 AM** — Professor mentions "Assignment 3 due Friday 11:59 PM, covers Streams and Optional."

- Alex screenshots the slide → pastes into Gatherer
- Gatherer auto-extracts: *"Task detected: Assignment 3, deadline Fri 11:59 PM, topic: Streams, Optional"*
- Task appears in Alex's task list, linked to the screenshot.

**10:00 AM** — Lecture ends. Alex downloads the PDF slides from LumiNUS.

- Alex saves PDF to `~/Downloads/Gatherer/` (watched folder)
- Gatherer auto-imports, detects "CS2030S" in the filename and content
- Processes: extracts text, generates summary, links to this morning's screenshots
- Alex never touches the file again—it's searchable and connected.

### Midday: Face-to-Face Lecture (CS2040S)

**12:00 PM** — Alex attends a lecture on graph algorithms. Takes photos of the whiteboard on phone.

- Later, Alex transfers photos to laptop → drops into watched folder
- Gatherer OCRs the whiteboard images, detects "BFS", "DFS", "CS2040S" from content
- Auto-tags and links to previous CS2040S materials.

**12:30 PM** — Alex jots quick notes in Notepad during lecture:
```
- BFS uses queue
- DFS uses stack
- time complexity O(V+E)
```
- Saves as `notes.txt` to watched folder
- Gatherer imports, but context is unclear (no module code in text)
- **Prompt appears**: *"What's this about?"* with suggestions: `CS2040S` (recent), `CS2030S`, `Other`
- Alex clicks `CS2040S`. Gatherer links it to today's photos and previous graph materials.

### Afternoon: Break Between Classes

**2:00 PM** — Alex has 30 minutes before tutorial. Opens Gatherer.

**Dashboard shows:**
- "3 items from this morning need review" (the screenshots)
- "CS2030S Assignment 3 due in 4 days" (auto-extracted task)
- "You added 5 items about graphs today" (temporal clustering)

**Alex searches**: *"what's the difference between BFS and DFS?"*
- Gatherer returns: today's whiteboard photo, last week's lecture PDF, a linked YouTube video Alex saved last month
- All from different times, all semantically connected.

### Late Afternoon: Tutorial (IS1108)

**4:00 PM** — Tutorial for Digital Ethics. Tutor discusses the essay assignment.

- Tutor writes on whiteboard: "Essay due Week 10. Topic: AI bias in hiring. 2000 words. Use at least 5 academic sources."
- Alex photos the whiteboard → transfers to watched folder
- Gatherer OCRs, extracts task: *"Essay, Week 10, AI bias in hiring, 2000 words, 5 sources"*
- Auto-tagged `IS1108` (detected from "ethics" and recent IS1108 pattern)

**4:30 PM** — Tutor shares useful links verbally. Alex copies them one by one:
- Copies URL → **Ctrl+V** in Gatherer → Gatherer fetches page, archives, extracts content
- Repeat 3 times. All auto-linked to IS1108 essay task.

### Evening: Working on Assignment

**7:00 PM** — Alex starts CS2030S Assignment 3.

**Searches**: *"CS2030S streams examples"*
- Results: this morning's screenshots, lecture PDF, tutorial worksheet from 2 weeks ago, a Stack Overflow link Alex saved

**While working**, Alex finds a useful blog post:
- Copies URL → Ctrl+V → Gatherer archives it
- Gatherer suggests: *"Link to CS2030S Assignment 3?"* (based on current activity)
- Alex confirms. The resource is now connected to the assignment.

**8:00 PM** — Alex edits the assignment PDF with annotations (using external PDF editor).
- Gatherer detects the file changed in watched folder
- Re-processes, extracts new annotations separately from original content
- Annotations searchable: *"my notes on assignment 3"* finds them.

### Night: Planning Tomorrow

**10:00 PM** — Alex opens Gatherer's **Timeline View**.

**Sees:**
- Tomorrow: CS2040S tutorial (needs to review graphs)
- Friday: CS2030S Assignment 3 due
- Week 10: IS1108 Essay due

**Tasks View shows:**
- "Review BFS/DFS for tutorial" (auto-suggested based on tomorrow's schedule)
- "Complete CS2030S Assignment 3" — estimated 4 hours (based on past assignments)
- "Start IS1108 essay research" — 5 sources needed

Alex marks "Review BFS/DFS" as priority for tomorrow morning.

### The Intelligence in Action

Throughout the day, Gatherer demonstrated:

| Feature | Example |
|---------|---------|
| **Context inference** | Recognized "CS2030S" from content without manual tagging |
| **Minimal prompts** | Only asked for context when truly ambiguous (the quick notes) |
| **Smart suggestions** | Offered likely tags based on recent activity and content |
| **Temporal awareness** | Grouped morning items together, knew "today's" context |
| **Task extraction** | Found deadlines in screenshots and created tasks automatically |
| **File change detection** | Noticed PDF annotations and re-processed |
| **Semantic linking** | Connected BFS/DFS content across photos, notes, PDFs, and videos |
| **One-click capture** | Clipboard paste with single confirmation click |

**What Alex never did:**
- Manually create folders or organize files
- Type out task details or deadlines
- Tag items with categories
- Remember where anything was saved
- Worry about duplicates or versions

---

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
Ingestion → Storage → Context Inference → Processing → Embedding → Linking
```

**Processing Steps:**

1. **Storage**
   - Save original file to organized directory structure
   - Generate unique identifier (UUID)
   - Create database record with metadata
   - Record timestamp and source (clipboard, folder, drag-drop)

2. **Content Extraction**
   - PDF: Extract text using pdfcpu/unidoc
   - Images: OCR via Tesseract, QR code detection via gozxing
   - Audio: Transcribe via Whisper API
   - Web: Extract main content, strip ads/navigation
   - DOCX: Extract text and formatting
   - Detect file changes: If file was seen before, extract only new/changed content (e.g., annotations)

3. **Context Inference** (the intelligence layer)

   This is the core of Gatherer's smart organization:

   - **Module/Topic Detection**: Scan content for module codes (CS2030S), course names, topic keywords
   - **Temporal Context**: What was the user doing in the last 30 minutes? Group with recent items
   - **Pattern Matching**: "Last 5 items were tagged CS2040S, this one probably is too"
   - **Confidence Scoring**: High confidence → auto-tag silently. Low confidence → prompt user

   **Prompt Strategy:**
   - Only prompt when confidence < 70%
   - Show top 2-3 suggestions as clickable buttons
   - "Other" option opens minimal text input
   - Learn from corrections to improve future inference

4. **Intelligence Extraction** (via LLM)
   - Summarization: Generate concise summary
   - Task Extraction: Identify action items, deadlines, requirements
   - Date Parsing: Extract all mentioned dates and deadlines (convert "Friday" to actual date)
   - Entity Recognition: People, places, topics
   - Requirement Extraction: "2000 words", "5 sources", specific deliverables

5. **Embedding Generation**
   - Split content into chunks (500-1000 tokens)
   - Generate embeddings via OpenAI API
   - Store in pgvector for semantic search

6. **Linking & Deduplication**
   - Compare embeddings with existing content
   - Flag duplicates (>0.90 similarity) → suggest keeping best version
   - Find related items (0.70-0.90 similarity) → create soft links
   - Link to active tasks if content is related (e.g., resource linked to assignment)
   - Temporal clustering: Items added within 15 minutes likely belong together

#### 4. Database Schema

**Core Tables:**

```sql
-- Items: All ingested content
items (
  id UUID PRIMARY KEY,
  title TEXT,
  source_type TEXT, -- clipboard, folder, upload, url
  content_type TEXT, -- pdf, image, audio, web, text
  original_filename TEXT,
  file_path TEXT,
  file_hash TEXT, -- for detecting changes
  ingested_at TIMESTAMP,
  status TEXT, -- pending, processing, completed, error
  context_id UUID REFERENCES contexts(id), -- linked context (module, project, etc.)
  context_confidence FLOAT, -- how confident was the auto-inference
  metadata JSONB
)

-- Contexts: Modules, projects, categories
contexts (
  id UUID PRIMARY KEY,
  name TEXT, -- "CS2030S", "IS1108 Essay", "Hall Application"
  context_type TEXT, -- module, project, personal, work
  keywords TEXT[], -- terms that suggest this context
  color TEXT,
  created_at TIMESTAMP,
  last_used_at TIMESTAMP -- for suggesting recent contexts
)

-- Context inference learning
context_corrections (
  id UUID PRIMARY KEY,
  item_id UUID REFERENCES items(id),
  suggested_context_id UUID REFERENCES contexts(id),
  actual_context_id UUID REFERENCES contexts(id),
  corrected_at TIMESTAMP
  -- Used to improve inference: "When I suggested X but user chose Y, what was different?"
)

-- Processed content
content (
  id UUID PRIMARY KEY,
  item_id UUID REFERENCES items(id),
  extracted_text TEXT,
  summary TEXT,
  is_annotation BOOLEAN DEFAULT FALSE, -- true if this is user-added content (annotations, highlights)
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

-- Item relationships (semantic links, temporal clusters)
item_links (
  id UUID PRIMARY KEY,
  item_id_1 UUID REFERENCES items(id),
  item_id_2 UUID REFERENCES items(id),
  link_type TEXT, -- 'similar', 'temporal', 'task_resource', 'duplicate'
  similarity_score FLOAT,
  created_at TIMESTAMP
)

-- Task resources (items linked to tasks)
task_resources (
  task_id UUID REFERENCES tasks(id),
  item_id UUID REFERENCES items(id),
  added_at TIMESTAMP,
  PRIMARY KEY (task_id, item_id)
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
   - By context (CS2030S, IS1108 Essay, etc.)
   - By source type (clipboard, folder, upload)
   - By content type (documents, audio, images)
   - By time range ("today", "this week", "last month")
   - By linked task or deadline

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
   - Recent items grouped by context
   - Quick search bar (always visible)
   - Pending context confirmations ("3 items need tagging")
   - Suggested actions ("Review these items for tomorrow's tutorial")
   - Activity summary ("You added 12 items about CS2040S today")

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
   - Extracted summary and key points
   - Context badge (CS2030S) with one-click change
   - Associated tasks (auto-linked or manually added)
   - Related items (semantically similar, same context, same time)
   - Annotations section (if file was modified externally)
   - Source information and processing history

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

**Always-On Monitoring:**
- Clipboard is monitored continuously while Gatherer runs
- When content is detected, a subtle indicator shows in system tray/app
- User can add content with a single hotkey (e.g., Ctrl+Shift+G) or click

**Quick Capture Flow:**
```
User copies something → Gatherer detects content →
Shows toast: "Add to Gatherer? [CS2030S] [CS2040S] [Other]" →
User clicks one button → Done
```

**Supported Content Types:**
- Images: Screenshots (Win+Shift+S), copied images from web/apps
- Text: Plain text, rich text, code snippets
- URLs: Automatically fetch and archive the page
- Files: Files copied from Explorer (Ctrl+C on files)

**Smart Context Suggestions:**
- Suggest context based on clipboard content analysis
- Suggest recent contexts ("You've been adding CS2030S items")
- Learn from corrections to improve suggestions

**Implementation:**
- Electron's clipboard API for cross-platform monitoring
- Polling interval: 500ms (configurable)
- Deduplication: Don't prompt for same content twice
- Background processing: Don't block user, process after capture

#### 9. Folder Watching

**Watched Folder Setup:**
- User configures one or more watched folders (e.g., `~/Downloads/Gatherer/`)
- Can be any folder—Downloads, Desktop, OneDrive sync folder, etc.
- Great for phone photo sync (photos land in folder → auto-import)

**Auto-Import Flow:**
```
File appears in folder → Gatherer detects →
Analyze content for context → High confidence? Auto-tag silently →
Low confidence? Queue for user prompt next time app is focused
```

**File Change Detection:**
- Track file hashes to detect modifications
- If a previously-imported file changes (e.g., PDF annotated), re-process
- Extract only new content (annotations, highlights)
- Don't create duplicate entries—update existing item

**Context Inference for Folder Items:**
- Analyze filename: `CS2030S_Lecture5.pdf` → tag CS2030S
- Analyze content: OCR finds "Data Structures" → suggest CS2040S
- Temporal grouping: 5 photos added in 2 minutes → same context

**Implementation:**
- Node.js `chokidar` for cross-platform file watching
- File hash comparison (MD5/SHA256) for change detection
- Queue system for batch processing (don't overwhelm during bulk imports)
- Notification when items need context confirmation

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
