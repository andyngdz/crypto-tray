import { Avatar, AvatarProps } from '@mui/material'

import logoFallback from '@/assets/images/logo.png'

const CRYPTO_ICONS = import.meta.glob<{ default: string }>(
  '@/assets/crypto-icons/*.svg',
  { eager: true }
)

interface CryptoIconProps extends AvatarProps {
  symbol: string
  size?: number
}

export function CryptoIcon({
  symbol,
  size = 20,
  ...restProps
}: CryptoIconProps) {
  const iconName = symbol.toLowerCase()
  const iconPath = `/src/assets/crypto-icons/${iconName}.svg`
  const icon = CRYPTO_ICONS[iconPath]
  const iconUrl = icon ? icon.default : logoFallback

  return (
    <Avatar
      src={iconUrl}
      alt={symbol}
      sx={{
        width: size,
        height: size,
        bgcolor: 'transparent',
      }}
      {...restProps}
    />
  )
}
