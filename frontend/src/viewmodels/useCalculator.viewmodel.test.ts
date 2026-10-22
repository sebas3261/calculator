import { act, renderHook } from '@testing-library/react'
import { describe, expect, it, vi } from 'vitest'
import { useCalculator } from './useCalculator.viewmodel'
import { calculatorService } from '../services/Calculator.service'
import { HttpError, NetworkError } from '../lib/http'

vi.mock('../services/Calculator.service', () => ({ calculatorService: { calculate: vi.fn() } }))
const calculate = vi.mocked(calculatorService.calculate)

function deferred<T>() {
  let resolve!: (value: T) => void
  let reject!: (reason: unknown) => void
  const promise = new Promise<T>((res, rej) => { resolve = res; reject = rej })
  return { promise, resolve, reject }
}

describe('useCalculator', () => {
  it('starts empty and ignores blank expressions', async () => {
    const { result } = renderHook(useCalculator)
    expect(result.current.result).toBe(0)
    expect(result.current.expression).toBe('')
    act(() => result.current.updateExpression('  '))
    await act(() => result.current.calculate())
    expect(calculate).not.toHaveBeenCalled()
  })

  it('normalizes visual operators and keeps root and power syntax', async () => {
    calculate.mockResolvedValueOnce({ answer: 11 })
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('√9+2^3×4÷4−0'))
    await act(() => result.current.calculate())
    expect(calculate).toHaveBeenCalledWith('√9+2^3*4/4-0')
    expect(result.current.result).toBe(11)
    expect(result.current.isLoading).toBe(false)
  })

  it('prevents duplicate requests before React renders and blocks button edits while pending', async () => {
    const request = deferred<{ answer: number }>()
    calculate.mockReturnValueOnce(request.promise)
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('2+2'))
    let first!: Promise<void>
    act(() => {
      first = result.current.calculate()
      void result.current.calculate()
      result.current.pressKey('9')
    })
    expect(calculate).toHaveBeenCalledTimes(1)
    expect(result.current.isLoading).toBe(true)
    expect(result.current.expression).toBe('2+2')
    await act(async () => { request.resolve({ answer: 4 }); await first })
    expect(result.current.result).toBe(4)
    expect(result.current.isLoading).toBe(false)
  })

  it('shows API errors, releases the request lock, and allows retry', async () => {
    calculate.mockRejectedValueOnce(new HttpError(422, 'division by zero'))
    calculate.mockResolvedValueOnce({ answer: 2 })
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('1/0'))
    await act(() => result.current.calculate())
    expect(result.current.error).toBe('division by zero')
    expect(result.current.isLoading).toBe(false)
    act(() => result.current.updateExpression('4/2'))
    expect(result.current.error).toBe('')
    await act(() => result.current.calculate())
    expect(result.current.result).toBe(2)
  })

  it('reports a network failure and stops loading', async () => {
    calculate.mockRejectedValueOnce(new NetworkError())
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('2+2'))
    await act(() => result.current.calculate())
    expect(result.current.error).toContain('Unable to connect')
    expect(result.current.isLoading).toBe(false)
  })

  it('rejects non-finite results without replacing the previous answer', async () => {
    calculate.mockResolvedValueOnce({ answer: Infinity })
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('2^9999'))
    await act(() => result.current.calculate())
    expect(result.current.error).toBe('The server returned an invalid calculation result.')
    expect(result.current.result).toBe(0)
  })

  it.each(['+', '−', '×', '÷', '^', '%'])('continues a negative answer with %s and preserves precedence', async (operator) => {
    calculate.mockResolvedValueOnce({ answer: -3 })
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('2-5'))
    await act(() => result.current.calculate())
    act(() => result.current.pressKey(operator))
    expect(result.current.expression).toBe(`(-3)${operator}`)
  })

  it.each(['7', '√', '('])('starts a fresh expression with %s after a result', async (key) => {
    calculate.mockResolvedValueOnce({ answer: 5 })
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('2+3'))
    await act(() => result.current.calculate())
    act(() => result.current.pressKey(key))
    expect(result.current.expression).toBe(key)
  })

  it('clears both the expression and result', async () => {
    calculate.mockResolvedValueOnce({ answer: 4 })
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('2+2'))
    await act(() => result.current.calculate())
    act(() => result.current.pressKey('AC'))
    expect(result.current.expression).toBe('')
    expect(result.current.result).toBe(0)
    expect(result.current.error).toBe('')
  })
})

describe('input validation and scientific result continuation', () => {
  it('rejects unsupported characters before calling the API', async () => {
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('2+alert(1)'))
    await act(() => result.current.calculate())
    expect(calculate).not.toHaveBeenCalled()
    expect(result.current.error).toContain('supported mathematical operators')
  })

  it.each([1e21, 1e-7, -1e21])('can reuse the result %s', async (answer) => {
    calculate.mockResolvedValueOnce({ answer }).mockResolvedValueOnce({ answer: answer * 2 })
    const { result } = renderHook(useCalculator)
    act(() => result.current.updateExpression('10^21'))
    await act(() => result.current.calculate())
    act(() => result.current.pressKey('×'))
    act(() => result.current.pressKey('2'))
    await act(() => result.current.calculate())
    expect(calculate).toHaveBeenLastCalledWith(`${answer < 0 ? `(${answer})` : answer}*2`)
    expect(result.current.result).toBe(answer * 2)
  })
})
