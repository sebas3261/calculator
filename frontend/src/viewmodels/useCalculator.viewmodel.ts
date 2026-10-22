import { useRef, useState } from 'react'
import { calculatorService } from '../services/Calculator.service'
import { HttpError } from '../lib/http'

const operators: Record<string, string> = { '×': '*', '÷': '/', '−': '-' }

export function useCalculator() {
  const [expression, setExpression] = useState('')
  const [result, setResult] = useState(0)
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [isComplete, setIsComplete] = useState(false)
  const inputRef = useRef<HTMLInputElement>(null)
  const pending = useRef(false)

  function updateExpression(value: string) {
    setExpression(value)
    setError('')
    setIsComplete(false)
  }

  function clear() {
    updateExpression('')
    setResult(0)
  }

  async function calculate() {
    if (pending.current || !expression.trim()) return
    pending.current = true
    setIsLoading(true)
    setError('')
    try {
      const normalized = expression.replace(/[×÷−]/g, (symbol) => operators[symbol])
      const response = await calculatorService.calculate(normalized)
      if (!Number.isFinite(response.answer)) throw new Error('Invalid result')
      setResult(response.answer)
      setIsComplete(true)
    } catch (cause) {
      setError(cause instanceof HttpError ? cause.message : 'Unable to calculate. Check that the server is running and try again.')
    } finally {
      pending.current = false
      setIsLoading(false)
    }
  }

  function pressKey(key: string) {
    if (pending.current) return
    if (key === '=') { void calculate(); return }
    if (key === 'AC') { clear(); return }

    const input = inputRef.current
    const start = input?.selectionStart ?? expression.length
    const end = input?.selectionEnd ?? expression.length
    let next: string
    let caret: number

    if (key === '⌫') {
      const from = start === end ? Math.max(0, start - 1) : start
      next = expression.slice(0, from) + expression.slice(end)
      caret = from
    } else if (isComplete) {
      const continues = ['+', '−', '×', '÷', '^', '%'].includes(key)
      next = (continues ? (result < 0 ? `(${result})` : String(result)) : '') + key
      caret = next.length
    } else {
      next = expression.slice(0, start) + key + expression.slice(end)
      caret = start + key.length
    }
    updateExpression(next)
    requestAnimationFrame(() => {
      input?.focus({ preventScroll: true })
      input?.setSelectionRange(caret, caret)
    })
  }

  return { expression, result, error, isLoading, inputRef, updateExpression, pressKey, clear, calculate }
}
