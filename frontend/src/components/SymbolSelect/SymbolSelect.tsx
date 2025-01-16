import { Button, Popover, PopoverContent } from '@carbon/react';
import { css } from '@emotion/css';
import { useState } from 'react';
import { useAppData } from '../../hooks/useAppData';
import { SymbolList } from './SymbolList';

interface Props {
  value: string;
  onChange: (value: string) => void;
  disabled?: boolean;
}

export const SymbolSelect = ({ value, onChange, disabled = false }: Props): JSX.Element => {
  const { appData } = useAppData();
  const theme = appData.preferences.generalPreferences.theme;
  const [open, setOpen] = useState(false);

  return (
    <Popover
      align="bottom-left"
      open={open}
      onKeyDown={(evt) => {
        if (evt.key === 'Escape') {
          setOpen(false);
        }
      }}
      isTabTip
      onRequestClose={() => setOpen(false)}
    >
      <Button
        onClick={() => {
          setOpen(!open);
        }}
        aria-expanded={open}
        disabled={disabled}
        className={styles.content(theme, disabled)}
        size="sm"
        kind="ghost"
      >
        {value}
      </Button>
      <PopoverContent>
        <SymbolList onSelect={onChange} />
      </PopoverContent>
    </Popover>
  );
};

const styles = {
  content: (theme: string, isDisabled: boolean) => css`
    color: ${theme === 'dark' ? '#f4f4f4' : '#161616'} !important;
    opacity: ${isDisabled ? 0.3 : 1};
  `,
};
