import { defineConfig } from 'vitest/config'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  test: {
    environment: 'jsdom',
    setupFiles: ['./src/test/setup.ts'],
    clearMocks: true,
    coverage: {
      provider: 'v8',
      include: ['src/viewmodels/**/*.ts', 'src/views/**/*.tsx', 'src/services/**/*.ts', 'src/lib/**/*.ts'],
      exclude: ['src/**/*.test.*'],
      reporter: ['text', 'html', 'lcov'],
    },
  },
})
