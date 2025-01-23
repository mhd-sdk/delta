import { Close, Save } from '@carbon/icons-react';
import { Button, TextInput } from '@carbon/react';
import { createRef, useEffect, useState } from 'react';
import { useClickOutside } from '../../hooks/useClickOutside';

interface Props {
  previousName: string;
  onSave: (name: string) => void;
  onCancel: () => void;
}

export const EditableRow = ({ onCancel, onSave, previousName }: Props): JSX.Element => {
  const editRef = createRef<HTMLInputElement>();
  const wrapperRef = createRef<HTMLDivElement>();
  const [value, setValue] = useState(previousName);
  useEffect(() => {
    editRef.current?.focus();
    // set value
  }, []);

  // onClick outside cancel
  useClickOutside(wrapperRef, onCancel);

  return (
    <div ref={wrapperRef}>
      <TextInput
        value={value}
        onChange={(e) => setValue(e.target.value)}
        size="lg"
        id="edit"
        ref={editRef}
        labelText={undefined}
        decorator={
          <div style={{ display: 'flex', alignItems: 'center', right: '0' }}>
            <Button kind="ghost" size="lg" hasIconOnly renderIcon={Save} iconDescription="Save" onClick={() => onSave(value)} />
            <Button size="lg" kind="danger--ghost" iconDescription="Cancel" hasIconOnly renderIcon={Close} aria-label="Cancel" onClick={onCancel} />
          </div>
        }
      />
    </div>
  );
};
