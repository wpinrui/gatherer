import { useState, useEffect } from 'react';
import './ItemsList.css';

interface Item {
  id: string;
  original_name: string;
  file_size: number;
  mime_type?: string;
  created_at: string;
}

type LoadStatus = 'loading' | 'success' | 'error';

export function ItemsList() {
  const [items, setItems] = useState<Item[]>([]);
  const [status, setStatus] = useState<LoadStatus>('loading');
  const [error, setError] = useState('');

  const fetchItems = async () => {
    setStatus('loading');
    try {
      const response = await fetch('/items');
      if (!response.ok) {
        throw new Error('Failed to fetch items');
      }
      const data = await response.json();
      setItems(data);
      setStatus('success');
    } catch (err) {
      setStatus('error');
      setError(err instanceof Error ? err.message : 'Failed to load items');
    }
  };

  useEffect(() => {
    fetchItems();
  }, []);

  if (status === 'loading') {
    return (
      <div className="items-list">
        <h2>Your Items</h2>
        <p className="loading">Loading...</p>
      </div>
    );
  }

  if (status === 'error') {
    return (
      <div className="items-list">
        <h2>Your Items</h2>
        <p className="error">{error}</p>
        <button onClick={fetchItems}>Retry</button>
      </div>
    );
  }

  return (
    <div className="items-list">
      <h2>Your Items</h2>
      {items.length === 0 ? (
        <p className="empty">No items yet. Upload a file to get started.</p>
      ) : (
        <ul>
          {items.map((item) => (
            <li key={item.id} className="item">
              <div className="item-name">{item.original_name}</div>
              <div className="item-meta">
                <span>{formatBytes(item.file_size)}</span>
                <span>{formatDate(item.created_at)}</span>
              </div>
            </li>
          ))}
        </ul>
      )}
      <button onClick={fetchItems} className="refresh-btn">
        Refresh
      </button>
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

function formatDate(isoString: string): string {
  const date = new Date(isoString);
  return date.toLocaleDateString(undefined, {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
}
