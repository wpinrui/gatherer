import './App.css'
import { FileUpload } from './components/FileUpload'
import { ItemsList } from './components/ItemsList'

function App() {
  return (
    <div className="app">
      <h1>Gatherer</h1>
      <p>Intelligent information aggregator</p>
      <FileUpload />
      <ItemsList />
    </div>
  )
}

export default App
