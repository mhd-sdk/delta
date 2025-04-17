import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { CandlestickData, UTCTimestamp, WhitespaceData } from 'lightweight-charts';
import { ChartConfig } from './panel';
import { Unit } from './timerange';
import { rangeToDates } from './utils';

interface GetCandlesticksParams {
  config: ChartConfig;
  symbol: string;
}

type CandlestickResult = (CandlestickData<UTCTimestamp> | WhitespaceData<UTCTimestamp>)[];

interface BarData {
  timestamp: string;
  open: number;
  high: number;
  low: number;
  close: number;
  volume: number;
}

const fetchCandlesticks = async ({ config, symbol }: GetCandlesticksParams): Promise<CandlestickResult> => {
  const { start, end } = rangeToDates(config.range);

  const response = await axios.get<BarData[]>(`/api/market-data/bars`, {
    params: {
      symbol,
      timeframe: `${config.timeframe.n}${config.timeframe.unit}`,
      start: start.toISOString(),
      end: end.toISOString(),
    },
  });

  return response.data.map((bar) => ({
    time: (new Date(bar.timestamp).getTime() / 1000) as UTCTimestamp,
    open: bar.open,
    high: bar.high,
    low: bar.low,
    close: bar.close,
  }));
};

export const useGetCandlesticks = ({ config, symbol }: GetCandlesticksParams) => {
  const isOneMinute = config.timeframe.n === 1 && config.timeframe.unit === Unit.min;

  return useQuery({
    queryKey: ['candlesticks', symbol, config.timeframe.n, config.timeframe.unit, config.range],
    queryFn: () => fetchCandlesticks({ config, symbol }),
    refetchInterval: isOneMinute ? 60000 : false,
    staleTime: isOneMinute ? 60000 : 5 * 60000,
  });
};
