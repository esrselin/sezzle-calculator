import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { describe, expect, it, vi } from 'vitest'
import App from './App'

describe('Calculator', () => {
  it('renders the calculator', () => {
    render(<App />)

    expect(screen.getByText('Calculator')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('First number')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('Second number')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'Calculate' })).toBeInTheDocument()
  })

  it('calculates and displays the result', async () => {
    const fetchMock = vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => ({ result: 15 }),
    } as Response)

    render(<App />)

    fireEvent.change(screen.getByPlaceholderText('First number'), {
      target: { value: '10' },
    })

    fireEvent.change(screen.getByPlaceholderText('Second number'), {
      target: { value: '5' },
    })

    fireEvent.click(screen.getByRole('button', { name: 'Calculate' }))

    await waitFor(() => {
      expect(screen.getByText('15')).toBeInTheDocument()
    })

    expect(fetchMock).toHaveBeenCalledWith(
      'http://localhost:8080/calculate',
      expect.objectContaining({
        method: 'POST',
      }),
    )

    fetchMock.mockRestore()
  })
})
it('shows an error when numbers are missing', () => {
  render(<App />)

  fireEvent.click(screen.getByRole('button', { name: 'Calculate' }))

  expect(
    screen.getByText('Please enter a number.')
  ).toBeInTheDocument()
})
it('shows an error when the API returns an error', async () => {
  const fetchMock = vi.spyOn(globalThis, 'fetch').mockResolvedValue({
    ok: false,
    json: async () => ({ error: 'Cannot divide by zero.' }),
  } as Response)

  render(<App />)

  fireEvent.change(screen.getByPlaceholderText('First number'), {
    target: { value: '10' },
  })

  fireEvent.change(screen.getByPlaceholderText('Second number'), {
    target: { value: '0' },
  })

  fireEvent.click(screen.getByRole('button', { name: 'Calculate' }))

  await waitFor(() => {
    expect(
      screen.getByText('Cannot divide by zero.')
    ).toBeInTheDocument()
  })

  fetchMock.mockRestore()
})
it('calculates square root', async () => {
  const fetchMock = vi.spyOn(globalThis, 'fetch').mockResolvedValue({
    ok: true,
    json: async () => ({ result: 5 }),
  } as Response)

  render(<App />)

  fireEvent.change(screen.getByPlaceholderText('First number'), {
    target: { value: '25' },
  })

  fireEvent.change(screen.getByRole('combobox'), {
    target: { value: 'sqrt' },
  })

  expect(
    screen.queryByPlaceholderText('Second number')
  ).not.toBeInTheDocument()

  fireEvent.click(screen.getByRole('button', { name: 'Calculate' }))

  await waitFor(() => {
    expect(screen.getByText('5')).toBeInTheDocument()
  })

  fetchMock.mockRestore()
})
it('calculates percentage', async () => {
  const fetchMock = vi.spyOn(globalThis, 'fetch').mockResolvedValue({
    ok: true,
    json: async () => ({ result: 5 }),
  } as Response)

  render(<App />)

  fireEvent.change(screen.getByPlaceholderText('First number'), {
    target: { value: '50' },
  })

  fireEvent.change(screen.getByRole('combobox'), {
    target: { value: '%' },
  })

  fireEvent.change(screen.getByPlaceholderText('Second number'), {
    target: { value: '10' },
  })

  fireEvent.click(screen.getByRole('button', { name: 'Calculate' }))

  await waitFor(() => {
    expect(screen.getByText('5')).toBeInTheDocument()
  })

  fetchMock.mockRestore()
})