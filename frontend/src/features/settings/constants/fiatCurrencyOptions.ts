// Fiat currency options for the display currency dropdown
export const FIAT_CURRENCY_OPTIONS = [
  // Major currencies
  { value: 'usd', label: 'USD ($)' },
  { value: 'eur', label: 'EUR (\u20AC)' },
  { value: 'gbp', label: 'GBP (\u00A3)' },
  { value: 'jpy', label: 'JPY (\u00A5)' },
  { value: 'chf', label: 'CHF (Fr)' },
  { value: 'cad', label: 'CAD ($)' },
  { value: 'aud', label: 'AUD ($)' },
  { value: 'nzd', label: 'NZD ($)' },
  { value: 'cny', label: 'CNY (\u00A5)' },
  { value: 'hkd', label: 'HKD ($)' },
  { value: 'sgd', label: 'SGD ($)' },

  // European currencies
  { value: 'sek', label: 'SEK (kr)' },
  { value: 'nok', label: 'NOK (kr)' },
  { value: 'dkk', label: 'DKK (kr)' },

  // Asian currencies
  { value: 'krw', label: 'KRW (\u20A9)' },
  { value: 'twd', label: 'TWD (NT$)' },
  { value: 'inr', label: 'INR (\u20B9)' },
  { value: 'thb', label: 'THB (\u0E3F)' },
  { value: 'vnd', label: 'VND (\u20AB)' },

  // Americas
  { value: 'brl', label: 'BRL (R$)' },
  { value: 'mxn', label: 'MXN ($)' },

  // Other
  { value: 'zar', label: 'ZAR (R)' },
] as const

export type FiatCurrencyOption = (typeof FIAT_CURRENCY_OPTIONS)[number]
