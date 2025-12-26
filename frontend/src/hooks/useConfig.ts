import { useState, useEffect, useCallback } from 'react';
import { GetConfig, SaveConfig, GetAvailableProviders } from '../../wailsjs/go/main/App';
import { isConfig, isProviderInfo } from '../types/config';
import type { Config, ProviderInfo } from '../types/config';

export function useConfig() {
  const [config, setConfig] = useState<Config | null>(null);
  const [providers, setProviders] = useState<ProviderInfo[]>([]);
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function load() {
      try {
        const [cfg, provs] = await Promise.all([
          GetConfig(),
          GetAvailableProviders(),
        ]);

        // Validate responses with type guards
        if (!isConfig(cfg)) {
          throw new Error('Invalid config response from backend');
        }

        if (!Array.isArray(provs) || !provs.every(isProviderInfo)) {
          throw new Error('Invalid providers response from backend');
        }

        setConfig(cfg);
        setProviders(provs);
      } catch (err) {
        setError(err instanceof Error ? err.message : String(err));
      } finally {
        setLoading(false);
      }
    }
    load();
  }, []);

  const updateConfig = useCallback(async (updates: Partial<Config>) => {
    if (!config) return;

    const newConfig = { ...config, ...updates };
    setConfig(newConfig);
    setSaving(true);

    try {
      await SaveConfig(newConfig);
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err));
      setConfig(config); // Revert on error
    } finally {
      setSaving(false);
    }
  }, [config]);

  return { config, providers, loading, saving, error, updateConfig };
}
