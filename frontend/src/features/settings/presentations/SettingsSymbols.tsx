import { ComboBox, Input, Label, ListBox, Tag, TagGroup } from '@heroui/react'

import { useSettingsSymbols } from '@/features/settings/states/useSettingsSymbols'

export function SettingsSymbols() {
  const { inputValue, setInputValue, filteredSymbols, selectedSymbols, onSelect, onRemove } =
    useSettingsSymbols()

  return (
    <div className="flex flex-col gap-2">
      <Label>Currencies</Label>

      {selectedSymbols.length > 0 && (
        <TagGroup onRemove={onRemove}>
          <TagGroup.List>
            {selectedSymbols.map((s) => (
              <Tag key={s.symbol} id={s.symbol} textValue={s.name}>
                {s.symbol}
              </Tag>
            ))}
          </TagGroup.List>
        </TagGroup>
      )}

      <ComboBox
        allowsEmptyCollection
        className="w-full"
        inputValue={inputValue}
        selectedKey={null}
        onInputChange={setInputValue}
        onSelectionChange={onSelect}
      >
        <ComboBox.InputGroup>
          <Input placeholder="Search currencies..." />
          <ComboBox.Trigger />
        </ComboBox.InputGroup>
        <ComboBox.Popover>
          <ListBox>
            {filteredSymbols.map((s) => (
              <ListBox.Item key={s.symbol} id={s.symbol} textValue={s.name}>
                {s.symbol} - {s.name}
              </ListBox.Item>
            ))}
          </ListBox>
        </ComboBox.Popover>
      </ComboBox>
    </div>
  )
}
