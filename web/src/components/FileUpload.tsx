import { useState, useRef } from 'react';
import type { FormEvent, ChangeEvent } from 'react';
import './FileUpload.css';

interface UploadResponse {
  id: string;
  filename: string;
  size: number;
  path: string;
  created_at: string;
}

type UploadStatus = 'idle' | 'uploading' | 'success' | 'error';

export function FileUpload() {
  const [status, setStatus] = useState<UploadStatus>('idle');
  const [message, setMessage] = useState('');
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0] || null;
    setSelectedFile(file);
    setStatus('idle');
    setMessage('');
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    if (!selectedFile) {
      setStatus('error');
      setMessage('Please select a file');
      return;
    }

    setStatus('uploading');
    setMessage('Uploading...');

    const formData = new FormData();
    formData.append('file', selectedFile);

    try {
      const response = await fetch('/upload', {
        method: 'POST',
        body: formData,
      });

      const text = await response.text();
      if (!text) {
        throw new Error('Server returned empty response. Is the backend running?');
      }

      let parsed: unknown;
      try {
        parsed = JSON.parse(text);
      } catch {
        throw new Error('Invalid response from server');
      }

      if (!response.ok) {
        const errorData = parsed as { error?: string };
        throw new Error(errorData.error || 'Upload failed');
      }

      const data = parsed as UploadResponse;
      setStatus('success');
      setMessage(`Uploaded: ${data.filename} (${formatBytes(data.size)})`);

      // Reset file input
      setSelectedFile(null);
      if (fileInputRef.current) {
        fileInputRef.current.value = '';
      }
    } catch (err) {
      setStatus('error');
      setMessage(err instanceof Error ? err.message : 'Upload failed');
    }
  };

  return (
    <div className="file-upload">
      <h2>Upload File</h2>
      <form onSubmit={handleSubmit}>
        <div className="file-input-wrapper">
          <input
            ref={fileInputRef}
            type="file"
            onChange={handleFileChange}
            disabled={status === 'uploading'}
          />
        </div>
        {selectedFile && (
          <p className="selected-file">
            Selected: {selectedFile.name} ({formatBytes(selectedFile.size)})
          </p>
        )}
        <button
          type="submit"
          disabled={!selectedFile || status === 'uploading'}
        >
          {status === 'uploading' ? 'Uploading...' : 'Upload'}
        </button>
      </form>
      {message && (
        <p className={`message ${status}`}>{message}</p>
      )}
    </div>
  );
}

function formatBytes(bytes: number): string {
  if (bytes <= 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.min(Math.floor(Math.log(bytes) / Math.log(k)), sizes.length - 1);
  return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i];
}
