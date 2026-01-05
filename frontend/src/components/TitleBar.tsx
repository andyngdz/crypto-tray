import { Box, IconButton } from '@mui/material'
import { Quit, WindowMinimise, WindowToggleMaximise } from '@wailsjs/runtime/runtime'
import { Minus, Square, X } from 'lucide-react'

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
        onClick={WindowMinimise}
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
        onClick={WindowToggleMaximise}
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
        onClick={Quit}
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
