import { Button } from '@/components/ui/button';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { cn } from '@/lib/utils';
import { Check, ChevronsUpDown } from 'lucide-react';
import { useState } from 'react';

// Mocked ticker options
const mockTickers = [
  { id: 'AAPL', name: 'Apple Inc.' },
  { id: 'MSFT', name: 'Microsoft Corporation' },
  { id: 'GOOGL', name: 'Alphabet Inc.' },
  { id: 'AMZN', name: 'Amazon.com, Inc.' },
  { id: 'TSLA', name: 'Tesla, Inc.' },
  { id: 'META', name: 'Meta Platforms, Inc.' },
  { id: 'NFLX', name: 'Netflix, Inc.' },
  { id: 'NVDA', name: 'NVIDIA Corporation' },
];

interface Props {
  value: string;
  onChange: (value: string) => void;
  disabled?: boolean;
}

export const TickerSelect = ({ value, onChange, disabled = false }: Props) => {
  const [open, setOpen] = useState(false);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button variant="outline" size="sm" disabled={disabled} className={cn('w-[180px] justify-between', disabled && 'opacity-50')}>
          {value || 'Select ticker'}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[180px] p-0">
        <div className="max-h-[300px] overflow-auto">
          {mockTickers.map((ticker) => (
            <div
              key={ticker.id}
              className={cn(
                'flex items-center justify-between px-2 py-1.5 text-sm cursor-pointer hover:bg-slate-100 dark:hover:bg-slate-800',
                value === ticker.id && 'bg-slate-100 dark:bg-slate-800'
              )}
              onClick={() => {
                onChange(ticker.id);
                setOpen(false);
              }}
            >
              <span>{ticker.id}</span>
              {value === ticker.id && <Check className="h-4 w-4" />}
            </div>
          ))}
        </div>
      </PopoverContent>
    </Popover>
  );
};
