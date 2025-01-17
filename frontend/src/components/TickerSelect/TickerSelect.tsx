import { Button, Popover, PopoverContent } from '@carbon/react';
import { css } from '@emotion/css';
import { useState } from 'react';
import { useAppData } from '../../hooks/useAppData';
import { TickerList } from './TickerList';

interface Props {
  value: string;
  onChange: (value: string) => void;
  disabled?: boolean;
}

export const TickerSelect = ({ value, onChange, disabled = false }: Props): JSX.Element => {
  const { appData } = useAppData();
  const theme = appData.preferences.generalPreferences.theme;
  const [isOpen, setIsOpen] = useState(false);
  const handleOnKeyDown = (evt: React.KeyboardEvent) => {
    if (evt.key === 'Escape') {
      console.log('Escape');
      setIsOpen(false);
    }
  };

  const handleSelect = (ticker: string) => {
    onChange(ticker);
    console.log('handleSelect', ticker);
    setIsOpen(false);
  };

  return (
    <Popover align="bottom-left" open={isOpen} onKeyDown={handleOnKeyDown} isTabTip dropShadow>
      <Button
        className={styles.content(theme, disabled)}
        onClick={() => setIsOpen(!isOpen)}
        aria-expanded={isOpen}
        disabled={disabled}
        kind="ghost"
        size="sm"
      >
        {value}
      </Button>
      <PopoverContent>{isOpen && <TickerList onSelect={handleSelect} onClose={() => setIsOpen(false)} />}</PopoverContent>
    </Popover>
  );
};

const styles = {
  content: (theme: string, isDisabled: boolean) => css`
    color: ${theme === 'dark' ? '#f4f4f4' : '#161616'} !important;
    opacity: ${isDisabled ? 0.3 : 1};
  `,
};
