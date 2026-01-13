import { Box, IconButton } from '@mui/material'
import { Window } from '@wailsio/runtime'
import { Minus, Square, X } from 'lucide-react'

import { hideWindow } from '@/features/settings/services/configService'

export function TitleBar() {
  return (
    <Box
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-end',
        height: 32,
        px: 0.5,
        // Make title bar draggable
        '--wails-draggable': 'drag',
      }}
    >
      <IconButton
        size="small"
        onClick={() => Window.Minimise()}
        sx={{
          '--wails-draggable': 'no-drag',
          color: 'text.secondary',
          '&:hover': { color: 'text.primary' },
        }}
      >
        <Minus size={16} />
      </IconButton>
      <IconButton
        size="small"
        onClick={() => Window.ToggleMaximise()}
        sx={{
          '--wails-draggable': 'no-drag',
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
          '--wails-draggable': 'no-drag',
          color: 'text.secondary',
          '&:hover': { color: 'error.main' },
        }}
      >
        <X size={16} />
      </IconButton>
    </Box>
  )
}
