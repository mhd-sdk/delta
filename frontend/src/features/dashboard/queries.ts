import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { ChartConfig } from './panel';
import { Unit } from './timerange';
import { rangeToDates } from './utils';

interface GetCandlesticksParams {
  config: ChartConfig;
  symbol: string;
}

const fetchCandlesticks = async ({
  config,
  symbol,
}: GetCandlesticksParams): Promise<
  {
    time: string; // timestamp ISO
    open: number; // open
    high: number; // high
    low: number; // low
    close: number; // close
    volume: number; // volume
  }[]
> => {
  const { start, end } = rangeToDates(config.range);

  const response = await axios.post(`http://localhost:3000/api/market-data/bars`, {
    data: {
      symbol,
      timeframe: config.timeframe,
      start: start.toISOString(),
      end: end.toISOString(),
    },
  });
  return response.data.map((bar: any) => ({
    // transform time to unix timestamp
    time: Math.floor(new Date(bar.t).getTime() / 1000),
    open: bar.o,
    high: bar.h,
    low: bar.l,
    close: bar.c,
    volume: bar.v,
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
