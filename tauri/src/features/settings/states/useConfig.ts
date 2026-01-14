import {
  fetchAvailableSymbols,
  fetchConfig,
  fetchProviders,
  saveConfig,
} from '@/features/settings/services/configService'
import type { Config } from '@/features/settings/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'

const QUERY_KEYS = {
  config: ['config'] as const,
  providers: ['providers'] as const,
  symbols: ['symbols'] as const,
}

export function useConfig() {
  const queryClient = useQueryClient()

  const configQuery = useQuery({
    queryKey: QUERY_KEYS.config,
    queryFn: fetchConfig,
  })

  const providersQuery = useQuery({
    queryKey: QUERY_KEYS.providers,
    queryFn: fetchProviders,
  })

  const symbolsQuery = useQuery({
    queryKey: QUERY_KEYS.symbols,
    queryFn: fetchAvailableSymbols,
  })

  const saveMutation = useMutation({
    mutationFn: saveConfig,
    onMutate: async (newConfig: Config) => {
      await queryClient.cancelQueries({ queryKey: QUERY_KEYS.config })
      const previousConfig = queryClient.getQueryData<Config>(QUERY_KEYS.config)
      queryClient.setQueryData(QUERY_KEYS.config, newConfig)
      return { previousConfig }
    },
    onSuccess: (data, newConfig, context) => {
      // If provider changed, invalidate symbols and refetch config to get reset symbols
      if (context?.previousConfig?.provider_id !== newConfig.provider_id) {
        queryClient.invalidateQueries({ queryKey: QUERY_KEYS.symbols })
        queryClient.invalidateQueries({ queryKey: QUERY_KEYS.config })
      }
    },
    onError: (err, _newConfig, context) => {
      if (context?.previousConfig) {
        queryClient.setQueryData(QUERY_KEYS.config, context.previousConfig)
      }
    },
  })

  const updateConfig = (updates: Partial<Config>) => {
    const currentConfig = configQuery.data

    if (!currentConfig) return

    saveMutation.mutate({ ...currentConfig, ...updates })
  }

  const loading =
    configQuery.isLoading || providersQuery.isLoading || symbolsQuery.isLoading
  const error =
    configQuery.error?.message ||
    providersQuery.error?.message ||
    symbolsQuery.error?.message ||
    saveMutation.error?.message

  return {
    config: configQuery.data,
    providers: providersQuery.data ?? [],
    availableSymbols: symbolsQuery.data ?? [],
    saving: saveMutation.isPending,
    loading,
    error,
    updateConfig,
  }
}
