import { Card, CardBody, CardHeader, Select, SelectItem, Input, Button, Divider, type SharedSelection } from '@heroui/react';
import { useConfig } from '../hooks/useConfig';
import { HideWindow } from '../../wailsjs/go/main/App';
import { StatusDisplay } from './StatusDisplay';
import { REFRESH_OPTIONS } from '../constants/settings';

// Generic handler factory for Select components - reduces duplication
function createSelectionHandler<T>(
  transform: (value: string) => T,
  update: (value: T) => void
) {
  return (keys: SharedSelection) => {
    if (keys === 'all') return;
    const value = Array.from(keys)[0];
    if (value) {
      update(transform(String(value)));
    }
  };
}

export function Settings() {
  const { config, providers, loading, saving, error, updateConfig } = useConfig();

  if (loading) {
    return <StatusDisplay type="loading" />;
  }

  if (!config) {
    return <StatusDisplay type="error" />;
  }

  const currentProvider = providers.find(p => p.id === config.provider_id);

  const handleProviderChange = createSelectionHandler(
    (id) => id,
    (providerId) => updateConfig({ provider_id: providerId })
  );

  const handleAPIKeyChange = (value: string) => {
    updateConfig({
      api_keys: {
        ...config.api_keys,
        [config.provider_id]: value,
      },
    });
  };

  const handleRefreshIntervalChange = createSelectionHandler(
    Number,
    (seconds) => updateConfig({ refresh_seconds: seconds })
  );

  const handleClose = () => {
    HideWindow();
  };

  return (
    <div className="min-h-screen p-6 bg-background">
      <Card className="max-w-md mx-auto">
        <CardHeader className="flex flex-col items-start gap-1">
          <h1 className="text-xl font-bold">Crypto Tray Settings</h1>
          <p className="text-small text-default-500">Configure your price tracker</p>
        </CardHeader>
        <Divider />
        <CardBody className="gap-4">
          {error && (
            <div className="p-3 text-sm text-white bg-danger rounded-lg">
              {error}
            </div>
          )}

          <Select
            label="API Provider"
            placeholder="Select a provider"
            selectedKeys={[config.provider_id]}
            onSelectionChange={handleProviderChange}
          >
            {providers.map((p) => (
              <SelectItem key={p.id}>
                {p.name}
              </SelectItem>
            ))}
          </Select>

          {currentProvider?.requiresApiKey && (
            <Input
              label="API Key"
              type="password"
              placeholder="Enter API key"
              value={config.api_keys[config.provider_id] || ''}
              onValueChange={handleAPIKeyChange}
            />
          )}

          <Select
            label="Refresh Interval"
            placeholder="Select interval"
            selectedKeys={[String(config.refresh_seconds)]}
            onSelectionChange={handleRefreshIntervalChange}
          >
            {REFRESH_OPTIONS.map((opt) => (
              <SelectItem key={String(opt.value)}>
                {opt.label}
              </SelectItem>
            ))}
          </Select>

          <Divider className="my-2" />

          <Button
            color="primary"
            className="w-full"
            onPress={handleClose}
            isLoading={saving}
          >
            {saving ? 'Saving...' : 'Close'}
          </Button>
        </CardBody>
      </Card>
    </div>
  );
}
