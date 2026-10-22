import './Calculator.view.css'
import { useCalculator } from '../viewmodels/useCalculator.viewmodel'

const keys = [
  'AC', '(', ')', '⌫',
  '√', '^', '%', '÷',
  '7', '8', '9', '×',
  '4', '5', '6', '−',
  '1', '2', '3', '+',
  '0', '.', '=',
]
const names: Record<string, string> = {
  AC: 'Clear all', '(': 'Open parenthesis', ')': 'Close parenthesis',
  '⌫': 'Delete last character', '√': 'Square root', '^': 'Power',
  '÷': 'Divide', '×': 'Multiply', '−': 'Subtract', '+': 'Add',
  '%': 'Percent', '.': 'Decimal point', '=': 'Calculate',
}

export default function Calculator() {
  const { expression, result, error, isLoading, inputRef, updateExpression, pressKey, clear, calculate } = useCalculator()
  return (
    <main className="calculator-page">
      <div className="calculator-layout">
      <section className="calculator-intro" aria-labelledby="intro-title">
        <span className="intro-eyebrow">A LITTLE LESS COMPLICATED</span>
        <h1 id="intro-title">Make room<br />for clarity.</h1>
        <p>From a quick sum to a bigger question.<br />A simple space to work things out.</p>
        <span className="intro-note">Every answer starts somewhere.</span>
      </section>
      <section className="calculator" aria-labelledby="calculator-title">
        <header className="calculator-header">
          <h2 id="calculator-title">Calculator</h2>
          <span>EXPRESSIONS</span>
        </header>
        <div className="calculator-display">
          <label htmlFor="expression">Expression</label>
          <input
            ref={inputRef}
            id="expression" type="text" placeholder="Type an expression"
            autoComplete="off" spellCheck={false} value={expression}
            readOnly={isLoading} aria-invalid={Boolean(error)}
            aria-describedby={error ? 'calculation-error' : undefined}
            onChange={(event) => updateExpression(event.target.value)}
            onKeyDown={(event) => {
              if (event.key === 'Enter' || event.key === '=') {
                event.preventDefault()
                void calculate()
              }
              if (event.key === 'Escape' && !isLoading) clear()
            }}
          />
          <output aria-label="Result" aria-live="polite" aria-busy={isLoading}>{isLoading ? '…' : result}</output>
          {error && <p id="calculation-error" className="calculator-error" role="alert">{error}</p>}
        </div>
        <div className="calculator-keypad">
          {keys.map((key) => (
            <button
              type="button"
              key={key}
              disabled={isLoading || (key === '=' && !expression.trim())}
              onMouseDown={(event) => event.preventDefault()}
              onClick={() => pressKey(key)}
              aria-label={names[key] ?? key}
              className={`calculator-key ${key === '=' ? 'equals' : ['÷', '×', '−', '+'].includes(key) ? 'operator' : names[key] && key !== '.' ? 'utility' : ''}`}
            >
              {key === '^' ? <span aria-hidden="true">x<sup>y</sup></span> : key}
            </button>
          ))}
        </div>
      </section>
      </div>
    </main>
  )
}
