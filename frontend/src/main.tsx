import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { HeroUIProvider } from '@heroui/react'
import React from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'
import './style.css'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60, // 1 minute
      retry: 1,
    },
  },
})

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <HeroUIProvider>
        <main className="dark text-foreground bg-background">
          <App />
        </main>
      </HeroUIProvider>
    </QueryClientProvider>
  </React.StrictMode>
)
