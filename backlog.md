# Gatherer Development Backlog

**Last Updated:** 2025-12-04

## Project Status

**Current Phase:** Phase 1 - Core Infrastructure
**Active Branch:** main
**Latest Commit:** feat: add database connection and item repository (PR #5 pending)

---

## Next Immediate Task

**Priority:** Frontend setup

**Suggested Next Task:**
- [ ] React app initialization with TypeScript and Vite

**Estimated Effort:** ~30 minutes, scaffolding + basic config

---

## Phase 1: Core Infrastructure (Current)

### Not Started
- [ ] React app initialization with TypeScript and Vite
- [ ] Basic file upload UI component
- [ ] Items list view component
- [ ] Simple file viewer (PDF.js for PDF, img tag for images)
- [ ] Application launcher script (start backend + open browser)

### In Progress
- [ ] PostgreSQL setup + database connection - PR #5 pending review

### Completed
- [x] Initialize Go module and project structure - PR #1
- [x] Basic Gin HTTP server with health check endpoint - PR #2
- [x] File upload endpoint (multipart/form-data) - PR #3
- [x] File storage service with interface - PR #4

---

## Phase 2: Content Processing (Upcoming)

### Planned Tasks
- [ ] Add pdfcpu dependency for PDF processing
- [ ] Create PDF text extraction function
- [ ] Add unit tests for PDF extraction
- [ ] Set up Redis for job queue
- [ ] Implement background job queue structure
- [ ] Create processing worker
- [ ] Add Tesseract OCR integration for images
- [ ] Integrate OpenAI API for summarization
- [ ] Store extracted content in database
- [ ] Display processing status in UI
- [ ] Show extracted text and summaries

---

## Phase 3: Search Foundation (Future)

### Planned Tasks
- [ ] Install pgvector extension in PostgreSQL
- [ ] Create embeddings table schema
- [ ] Implement content chunking logic
- [ ] Integrate OpenAI embedding API
- [ ] Store embeddings in database
- [ ] Create semantic search endpoint
- [ ] Build search UI component
- [ ] Implement result ranking and display
- [ ] Add keyword search fallback

---

## Phase 4-10: Later Phases

See `proposal.md` for detailed breakdown of:
- Phase 4: Task Extraction
- Phase 5: Browser Extension
- Phase 6: Audio Processing
- Phase 7: Email Integration
- Phase 8: Intelligence & Learning
- Phase 9: Advanced Features
- Phase 10: Optimization & Deployment

---

## Rejected Tasks (Too Large)

None yet

---

## Technical Debt / Improvements

None yet

---

## Questions / Decisions Needed

1. **Database Setup:** Local PostgreSQL or embedded (e.g., using Docker)?
2. **Frontend Routing:** React Router or Next.js?
3. **Desktop Launcher:** Simple shell script or Electron wrapper?
4. **API Keys:** How should OpenAI API key be configured? (Environment variable, config file?)

---

## Notes

- All tasks should be broken down into <300 line implementations
- Each task should result in a separate PR
- Follow the git workflow in claude.md
- Update this backlog after each PR is completed
- Refer to proposal.md for architectural guidance

---

## Task Size Guidelines

**Small Task (Ideal):**
- 50-150 lines of code
- 1-3 files changed
- Can be completed in 1-2 hours
- Single clear purpose

**Medium Task (Acceptable):**
- 150-300 lines of code
- 3-5 files changed
- Can be completed in half a day
- May need to be split if complex

**Large Task (Must Reject):**
- 300+ lines of code
- 5+ files significantly changed
- Takes multiple days
- Multiple concerns mixed together

**When in doubt, split it smaller.**
