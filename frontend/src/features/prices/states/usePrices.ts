import {
  fetchPrices,
  refreshPrices,
} from '@/features/prices/services/priceService'
import type { PriceData } from '@/features/prices/types'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { EventsOff, EventsOn } from '@wailsjs/runtime/runtime'
import { useCallback, useEffect, useMemo } from 'react'

interface UsePricesOptions {
  symbols: string[]
  enabled?: boolean
}

const QUERY_KEY = ['prices'] as const

export function usePrices({ symbols, enabled = true }: UsePricesOptions) {
  const queryClient = useQueryClient()

  const query = useQuery({
    queryKey: [...QUERY_KEY, symbols],
    queryFn: () => fetchPrices(symbols),
    enabled: enabled && symbols.length > 0,
    staleTime: Infinity,
    gcTime: 5 * 60 * 1000,
    refetchOnWindowFocus: false,
    retry: false,
  })

  const handlePriceUpdate = useCallback(
    (data: PriceData[]) => {
      queryClient.setQueryData([...QUERY_KEY, symbols], data)
    },
    [queryClient, symbols]
  )

  const priceMap = useMemo(() => {
    const map = new Map<string, PriceData>()
    query.data?.forEach((p) => map.set(p.symbol, p))
    return map
  }, [query.data])

  useEffect(() => {
    EventsOn('price:update', handlePriceUpdate)
    return () => EventsOff('price:update')
  }, [handlePriceUpdate])

  return {
    prices: query.data ?? [],
    priceMap,
    isLoading: query.isLoading,
    error: query.error?.message,
    refresh: refreshPrices,
  }
}
