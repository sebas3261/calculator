import { afterEach, describe, expect, it, vi } from 'vitest'
import { calculatorService } from './Calculator.service'
import { HttpError } from '../lib/http'

afterEach(() => vi.unstubAllGlobals())

describe('calculator API integration', () => {
  it('sends the expression as JSON and reads the answer', async () => {
    const fetch = vi.fn().mockResolvedValue(new Response(JSON.stringify({ answer: 4 })))
    vi.stubGlobal('fetch', fetch)
    await expect(calculatorService.calculate('2+2')).resolves.toEqual({ answer: 4 })
    expect(fetch).toHaveBeenCalledWith('/calculate', {
      method: 'POST', headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ expression: '2+2' }),
    })
  })

  it('preserves validation messages and HTTP status', async () => {
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue(new Response(JSON.stringify({ error: 'division by zero' }), { status: 422 })))
    await expect(calculatorService.calculate('1/0')).rejects.toMatchObject({
      name: 'HttpError', status: 422, message: 'division by zero',
    })
  })

  it('handles non-JSON error responses', async () => {
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue(new Response('Bad Gateway', { status: 502 })))
    await expect(calculatorService.calculate('2+2')).rejects.toBeInstanceOf(HttpError)
  })
})
