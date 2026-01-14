// Number format options for the settings dropdown
export const NUMBER_FORMAT_OPTIONS = [
  { value: 'us', label: 'US (1,234.56)' },
  { value: 'european', label: 'European (1.234,56)' },
  { value: 'asian', label: 'Asian (1,234.56)' },
] as const

export type NumberFormatOption = (typeof NUMBER_FORMAT_OPTIONS)[number]
