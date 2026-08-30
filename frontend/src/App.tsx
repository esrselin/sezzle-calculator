import './App.css'

function App() {
  return (
    <main className="calculator">
      <h1>Calculator</h1>

      <div className="display">
        0
      </div>

      <div className="calculator-body">
        <input type="number" placeholder="First number" />

        <select defaultValue="+">
          <option value="+">Addition (+)</option>
          <option value="-">Subtraction (-)</option>
          <option value="*">Multiplication (×)</option>
          <option value="/">Division (÷)</option>
          <option value="^">Exponentiation (^)</option>
        </select>

        <input type="number" placeholder="Second number" />

        <button type="button">
          Calculate
        </button>
      </div>
    </main>
  )
}

export default App