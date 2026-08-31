import { useState } from 'react'
import './App.css'

function App() {
  const [firstNumber, setFirstNumber] = useState('')
  const [secondNumber, setSecondNumber] = useState('')
  const [operation, setOperation] = useState('+')
  const [result, setResult] = useState('')

  const handleCalculate = async () => {
    const first = Number(firstNumber)
    const second = Number(secondNumber)

    const response = await fetch('http://localhost:8080/calculate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        firstNumber: first,
        operation: operation,
        secondNumber: second,
      }),
    })

    const data = await response.json()

    setResult(String(data.result))
  }

  return (
    <main className="calculator">
      <h1>Calculator</h1>

      <div className="display">
        {result || 0}
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
          onChange={(event) => setOperation(event.target.value)}
        >
          <option value="+">Addition (+)</option>
          <option value="-">Subtraction (-)</option>
          <option value="*">Multiplication (×)</option>
          <option value="/">Division (÷)</option>
          <option value="^">Exponentiation (^)</option>
        </select>

        <input
          type="number"
          placeholder="Second number"
          value={secondNumber}
          onChange={(event) => setSecondNumber(event.target.value)}
        />

        <button
          type="button"
          onClick={handleCalculate}
        >
          Calculate
        </button>
      </div>
    </main>
  )
}

export default App