import { Card, CardContent, CardHeader } from '@/components/ui/card';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { useTheme } from '@/context/theme-context';
import { CandlestickData, ColorType, createChart, CrosshairMode, IChartApi, ISeriesApi, UTCTimestamp, WhitespaceData } from 'lightweight-charts';
import { useEffect, useRef, useState } from 'react';
import { ChartConfig } from '../../panel';
import { useGetCandlesticks } from '../../queries';
import { defaultTimeframes, isIntraday, Timeframe, Unit } from '../../timerange';
import { calcOptimizedRange } from '../../utils';

interface Props {
  onDelete: () => void;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

interface ToolbarProps {
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

const popularTickers = ['AAPL', 'MSFT', 'GOOGL', 'AMZN', 'NVDA', 'META', 'TSLA', 'NFLX'];

const Toolbar = ({ config, onConfigChange }: ToolbarProps) => {
  const handleTickerChange = (value: string) => {
    onConfigChange({ ...config, ticker: value });
  };

  const handleTimeframeChange = (value: string) => {
    const [n, unit] = value.split('-');
    const timeframe: Timeframe = {
      n: parseInt(n),
      unit: unit as Unit,
    };
    onConfigChange({
      ...config,
      timeframe,
      range: calcOptimizedRange(timeframe),
    });
  };

  const { theme } = useTheme();
  const isDark = theme === 'dark';

  return (
    <div
      className="drag-handle flex flex-row items-center justify-between w-full space-x-2"
      style={{
        backgroundColor: isDark ? '#333' : '#f0f0f0',
        borderBottom: `1px solid ${isDark ? '#444' : '#ddd'}`,
        borderRadius: '0.375rem 0.375rem 0 0',
        padding: '0.5rem',
      }}
    >
      <div className="flex items-center space-x-2">
        <Select defaultValue={config.ticker} onValueChange={handleTickerChange}>
          <SelectTrigger className="drag-cancel w-24 h-8">
            <SelectValue placeholder="Ticker" />
          </SelectTrigger>
          <SelectContent>
            {popularTickers.map((ticker) => (
              <SelectItem key={ticker} value={ticker}>
                {ticker}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>

        <Select defaultValue={`${config.timeframe.n}-${config.timeframe.unit}`} onValueChange={handleTimeframeChange}>
          <SelectTrigger className="drag-cancel w-24 h-8">
            <SelectValue placeholder="Timeframe" />
          </SelectTrigger>
          <SelectContent>
            {defaultTimeframes.map((tf) => (
              <SelectItem key={`${tf.n}-${tf.unit}`} value={`${tf.n}-${tf.unit}`}>
                {`${tf.n} ${tf.unit}`}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>
    </div>
  );
};

export const Chart = ({ config, onDelete, onConfigChange }: Props) => {
  const [candlesticks, setCandlesticks] = useState<(CandlestickData<UTCTimestamp> | WhitespaceData<UTCTimestamp>)[]>([]);
  const chartContainerRef = useRef<HTMLDivElement>(null);
  const chartRef = useRef<IChartApi | null>(null);
  const candleSeriesRef = useRef<ISeriesApi<'Candlestick'> | null>(null);
  const cardRef = useRef<HTMLDivElement>(null);

  const el = useRef<HTMLDivElement | null>(null);

  const { theme } = useTheme();

  const {
    data: candlestickData,
    isLoading,
    error,
  } = useGetCandlesticks({
    config,
    symbol: config.ticker,
  });

  const darkColors = {
    backgroundColor: 'transparent',
    textColor: 'white',
    candleUpColor: '#ffffff',
    candleDownColor: '#5d606b',
  };

  const lightColors = {
    backgroundColor: 'transparent',
    textColor: 'black',
    candleUpColor: '#4caf50',
    candleDownColor: '#f44336',
  };

  const colors = theme === 'dark' ? darkColors : lightColors;

  useEffect(() => {
    if (!chartContainerRef.current) {
      return;
    }

    const chart = createChart(chartContainerRef.current, {
      timeScale: {
        timeVisible: isIntraday(config.timeframe),
      },
      grid: {
        horzLines: {
          visible: false,
        },
        vertLines: {
          visible: false,
        },
      },
      autoSize: true,
      layout: {
        background: { type: ColorType.Solid, color: colors.backgroundColor },
        textColor: colors.textColor,
      },
      crosshair: {
        mode: CrosshairMode.Normal,
      },
      width: chartContainerRef.current.clientWidth,
      height: chartContainerRef.current.clientHeight,
    });
    chartRef.current = chart;

    chart.timeScale().resetTimeScale();

    const candleSeries = chart.addCandlestickSeries({
      upColor: colors.candleUpColor,
      downColor: colors.candleDownColor,
      borderVisible: false,
      wickUpColor: colors.candleUpColor,
      wickDownColor: colors.candleDownColor,
    });
    candleSeriesRef.current = candleSeries;

    const resizeObserver = new ResizeObserver(() => {
      if (chartContainerRef.current) {
        chart.applyOptions({
          width: chartContainerRef.current.clientWidth,
          height: chartContainerRef.current.clientHeight,
        });
      }
      if (candleSeriesRef.current) {
        candleSeriesRef.current.applyOptions({
          borderVisible: false,
          wickUpColor: colors.candleUpColor,
          wickDownColor: colors.candleDownColor,
        });
      }
    });

    resizeObserver.observe(chartContainerRef.current);

    return () => {
      resizeObserver.disconnect();
      chart.remove();
    };
  }, []);

  useEffect(() => {
    if (candlestickData) {
      setCandlesticks(candlestickData);
    }
  }, [candlestickData]);

  useEffect(() => {
    if (candleSeriesRef.current) {
      candleSeriesRef.current.setData(candlesticks);
      chartRef.current?.timeScale().resetTimeScale();
    }
  }, [candlesticks]);

  useEffect(() => {
    if (chartRef.current) {
      chartRef.current.applyOptions({
        layout: {
          background: { type: ColorType.Solid, color: colors.backgroundColor },
          textColor: colors.textColor,
        },
      });
    }
    if (candleSeriesRef.current) {
      candleSeriesRef.current.applyOptions({
        upColor: colors.candleUpColor,
        downColor: colors.candleDownColor,
        borderVisible: false,
        wickUpColor: colors.candleUpColor,
        wickDownColor: colors.candleDownColor,
      });
    }
  }, [colors]);

  const handleContextMenu = (e: React.MouseEvent) => {
    e.preventDefault();
  };

  return (
    <Card className="h-full w-full py-0" ref={cardRef} onContextMenu={handleContextMenu}>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <div className="h-full w-full" onContextMenu={(e) => e.preventDefault()}>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 p-0">
              <Toolbar config={config} onConfigChange={onConfigChange} />
            </CardHeader>
            <CardContent className="h-[calc(100%-3rem)] pt-2">
              <div style={{ height: '100%', display: 'flex', flexDirection: 'column' }} ref={el}>
                {isLoading && <div>Loading...</div>}
                {error && <div>Error loading chart data</div>}
                <div ref={chartContainerRef} style={{ height: 'calc(100% - 2rem)' }} />
              </div>
            </CardContent>
          </div>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <DropdownMenuItem onClick={onDelete}>Delete Panel</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </Card>
  );
};
