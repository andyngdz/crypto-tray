import { createTheme } from '@mui/material'

export const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#F7931A',
    },
    background: {
      default: '#141414',
      paper: '#1c1c1c',
    },
  },
  components: {
    MuiTextField: { defaultProps: { size: 'small' } },
    MuiSelect: { defaultProps: { size: 'small' } },
    MuiButton: { defaultProps: { size: 'small' } },
    MuiChip: { defaultProps: { size: 'small' } },
    MuiAutocomplete: {
      defaultProps: { size: 'small' },
      styleOverrides: {
        listbox: {
          borderRadius: '8px',
          '& .MuiAutocomplete-option': {
            minHeight: 'auto',
            padding: '4px 8px',
          },
        },
      },
    },
    MuiFormControl: { defaultProps: { size: 'small' } },
    MuiInputLabel: {
      styleOverrides: {
        root: { left: 0 },
      },
    },
    MuiMenuItem: { defaultProps: { dense: true } },
  },
})
