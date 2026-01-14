import { Box, IconButton } from '@mui/material'
import { getCurrentWindow } from '@tauri-apps/api/window'
import { Minus, Square, X } from 'lucide-react'

import { hideWindow } from '@/features/settings/services/configService'

export function TitleBar() {
  const appWindow = getCurrentWindow()

  return (
    <Box
      data-tauri-drag-region
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-end',
        height: 32,
        px: 0.5,
      }}
    >
      <IconButton
        size="small"
        onClick={() => appWindow.minimize()}
        sx={{
          color: 'text.secondary',
          '&:hover': { color: 'text.primary' },
        }}
      >
        <Minus size={16} />
      </IconButton>
      <IconButton
        size="small"
        onClick={() => appWindow.toggleMaximize()}
        sx={{
          color: 'text.secondary',
          '&:hover': { color: 'text.primary' },
        }}
      >
        <Square size={14} />
      </IconButton>
      <IconButton
        size="small"
        onClick={hideWindow}
        sx={{
          color: 'text.secondary',
          '&:hover': { color: 'error.main' },
        }}
      >
        <X size={16} />
      </IconButton>
    </Box>
  )
}
