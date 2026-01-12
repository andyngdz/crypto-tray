import { Box, Typography } from '@mui/material'
import type { ReactNode } from 'react'

interface SettingsSectionProps {
  title: string
  children: ReactNode
}

export function SettingsSection({ title, children }: SettingsSectionProps) {
  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3, p: 3 }}>
      <Typography color="text.secondary">{title}</Typography>
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          gap: 2,
        }}
      >
        {children}
      </Box>
    </Box>
  )
}
