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
              <Tag key={s.id} id={s.id} textValue={s.name}>
                {s.id.toUpperCase()}
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
              <ListBox.Item key={s.id} id={s.id} textValue={s.name}>
                {s.id.toUpperCase()} - {s.name}
              </ListBox.Item>
            ))}
          </ListBox>
        </ComboBox.Popover>
      </ComboBox>
    </div>
  )
}
