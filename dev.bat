@echo off

echo Restarting PostgreSQL...
docker-compose down
docker-compose up -d

echo Waiting for PostgreSQL to be ready...
:wait
docker-compose exec -T postgres pg_isready -U gatherer >nul 2>&1
if errorlevel 1 (
    timeout /t 1 /nobreak >nul
    goto wait
)

echo Killing any existing process on port 8080...
for /f "tokens=5" %%a in ('netstat -ano ^| findstr :8080 ^| findstr LISTENING') do taskkill /F /PID %%a >nul 2>&1

echo PostgreSQL ready. Starting server...
go run ./cmd/gatherer

pause
