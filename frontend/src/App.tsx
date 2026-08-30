import { useState } from 'react'
import './App.css'

function App() {
  const [firstNumber, setFirstNumber] = useState('')
  const [secondNumber, setSecondNumber] = useState('')
  const [operation, setOperation] = useState('+')
  return (
    <main className="calculator">
      <h1>Calculator</h1>

      <div className="display">
        0
      </div>

      <div className="calculator-body">
        <input 
        type="number" 
        placeholder="First number" 
        value={firstNumber}
        onChange={(event) => setFirstNumber(event.target.value)}
        />

        <select
        value={operation}
        onChange={(event) => setOperation(event.target.value)}>
        <option value="+">Addition (+)</option>
        <option value="-">Subtraction (-)</option>
        <option value="*">Multiplication (×)</option>
        <option value="/">Division (÷)</option>
        <option value="^">Exponentiation (^)</option>
      </select>

        <input 
        type="number" 
        placeholder="Second number" 
        value = {secondNumber}
        onChange={(event) => setSecondNumber(event.target.value)}
        />

        <button 
        type="button"
        onClick={() => {
          console.log("clicked")
        }}
        >
          Calculate
        </button>
      </div>
    </main>
  )
}

export default App