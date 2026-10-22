import { act, fireEvent, render, screen, waitFor } from '@testing-library/react'
import { describe, expect, it, vi } from 'vitest'
import Calculator from './Calculator.view'
import { calculatorService } from '../services/Calculator.service'
import { HttpError } from '../lib/http'

vi.mock('../services/Calculator.service', () => ({ calculatorService: { calculate: vi.fn() } }))
const calculate = vi.mocked(calculatorService.calculate)
const button = (name: string) => screen.getByRole('button', { name })
const input = () => screen.getByRole('textbox', { name: 'Expression' }) as HTMLInputElement
const click = (name: string) => fireEvent.click(button(name))

describe('Calculator view', () => {
  it('calculates root and power using the keypad', async () => {
    calculate.mockResolvedValueOnce({ answer: 11 })
    render(<Calculator />)
    expect(button('Calculate')).toBeDisabled()
    for (const key of ['Square root', '9', 'Add', '2', 'Power', '3', 'Calculate']) click(key)
    await waitFor(() => expect(screen.getByLabelText('Result')).toHaveTextContent('11'))
    expect(calculate).toHaveBeenCalledWith('√9+2^3')
  })

  it('replaces selected text and restores the cursor after insertion', async () => {
    render(<Calculator />)
    fireEvent.change(input(), { target: { value: '1234' } })
    input().setSelectionRange(1, 3)
    click('9')
    expect(input()).toHaveValue('194')
    await waitFor(() => expect(input().selectionStart).toBe(2))
    expect(input().selectionEnd).toBe(2)
    expect(input()).toHaveFocus()
  })

  it('deletes a selection, backspaces at the cursor, and does nothing at the beginning', () => {
    render(<Calculator />)
    fireEvent.change(input(), { target: { value: '1234' } })
    input().setSelectionRange(1, 3)
    click('Delete last character')
    expect(input()).toHaveValue('14')
    input().setSelectionRange(1, 1)
    click('Delete last character')
    expect(input()).toHaveValue('4')
    input().setSelectionRange(0, 0)
    click('Delete last character')
    expect(input()).toHaveValue('4')
  })

  it.each(['Enter', '='])('calculates using %s and clears using Escape', async (key) => {
    calculate.mockResolvedValueOnce({ answer: 10 })
    render(<Calculator />)
    fireEvent.change(input(), { target: { value: '5+5' } })
    fireEvent.keyDown(input(), { key })
    await waitFor(() => expect(screen.getByLabelText('Result')).toHaveTextContent('10'))
    fireEvent.keyDown(input(), { key: 'Escape' })
    expect(input()).toHaveValue('')
    expect(screen.getByLabelText('Result')).toHaveTextContent('0')
  })

  it('locks input and buttons while waiting, then restores editing', async () => {
    let resolve!: (value: { answer: number }) => void
    calculate.mockReturnValueOnce(new Promise((done) => { resolve = done }))
    render(<Calculator />)
    fireEvent.change(input(), { target: { value: '2+2' } })
    click('Calculate')
    expect(input()).toHaveAttribute('readonly')
    for (const key of screen.getAllByRole('button')) expect(key).toBeDisabled()
    await act(async () => resolve({ answer: 4 }))
    expect(input()).not.toHaveAttribute('readonly')
    expect(button('Calculate')).toBeEnabled()
  })

  it('announces calculation errors and clears them when editing', async () => {
    calculate.mockRejectedValueOnce(new HttpError(422, 'division by zero'))
    render(<Calculator />)
    fireEvent.change(input(), { target: { value: '1/0' } })
    click('Calculate')
    expect(await screen.findByRole('alert')).toHaveTextContent('division by zero')
    expect(input()).toHaveAttribute('aria-invalid', 'true')
    fireEvent.change(input(), { target: { value: '1/2' } })
    expect(screen.queryByRole('alert')).not.toBeInTheDocument()
    expect(input()).toHaveAttribute('aria-invalid', 'false')
  })
})
