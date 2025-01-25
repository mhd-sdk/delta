import { Close, Save } from '@carbon/icons-react';
import { Button, TextInput } from '@carbon/react';
import { createRef, useEffect, useState } from 'react';
import { useAppData } from '../../hooks/useAppData';
import { useClickOutside } from '../../hooks/useClickOutside';

interface Props {
  previousName: string;
  onSave: (name: string) => void;
  onCancel: () => void;
  uniqueMode: 'new' | 'edit';
}

export const EditableRow = ({ onCancel, onSave, previousName, uniqueMode }: Props): JSX.Element => {
  const { appData, onSave: onSaveAppdata } = useAppData();

  const editRef = createRef<HTMLInputElement>();
  const wrapperRef = createRef<HTMLDivElement>();
  const [value, setValue] = useState(previousName);
  useEffect(() => {
    editRef.current?.focus();
    // set value
  }, []);

  // onClick outside cancel
  useClickOutside(wrapperRef, onCancel);

  const isUnique =
    uniqueMode === 'new'
      ? appData.workspaces.some((w) => w.name === value)
      : appData.workspaces.some((w) => w.name === value && w.name !== previousName);

  return (
    <div ref={wrapperRef}>
      <TextInput
        value={value}
        onChange={(e) => setValue(e.target.value)}
        size="lg"
        id="edit"
        ref={editRef}
        invalid={isUnique}
        invalidText="Name must be unique"
        labelText={undefined}
        decorator={
          <div style={{ display: 'flex', alignItems: 'center', right: isUnique ? '40px' : '0' }}>
            <Button kind="ghost" size="lg" hasIconOnly renderIcon={Save} iconDescription="Save" onClick={() => onSave(value)} disabled={isUnique} />

            <Button size="lg" kind="danger--ghost" iconDescription="Cancel" hasIconOnly renderIcon={Close} aria-label="Cancel" onClick={onCancel} />
          </div>
        }
      />
    </div>
  );
};
