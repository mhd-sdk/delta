import { Search } from '@carbon/icons-react';
import { Button } from '@carbon/react';
import { useEffect, useState } from 'react';
import { GetSymbols } from '../../../wailsjs/go/main/App';

interface Props {
  value: string;
  onChange: (value: string) => void;
}

export const SymbolSelect = ({ onChange, value }: Props): JSX.Element => {
  const [symbols, setSymbols] = useState<string[]>([]);
  useEffect(() => {
    const fetchSymbols = async () => {
      const res = await GetSymbols();
      setSymbols(res.map((s) => s.symbol));
    };
    fetchSymbols();
  }, []);
  return (
    <>
      <Button renderIcon={Search} size="sm" kind="ghost">
        {value}
      </Button>
    </>
  );
};
