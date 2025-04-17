import { useTheme } from '@/context/theme-context';
import { CandlestickData, ColorType, createChart, CrosshairMode, IChartApi, ISeriesApi, UTCTimestamp, WhitespaceData } from 'lightweight-charts';
import { useEffect, useRef, useState } from 'react';
import { ChartConfig } from '../panel';
import { useGetCandlesticks } from '../queries';
import { isIntraday } from '../timerange';

interface Props {
  onDelete: () => void;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

interface ToolbarProps {
  onDelete: () => void;
  config: ChartConfig;
}

// Simplified Toolbar component since the original isn't accessible
const Toolbar = ({ onDelete, config }: ToolbarProps) => {
  return (
    <div style={{ height: '2rem', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <div>{config.ticker}</div>
      <button onClick={onDelete}>Delete</button>
    </div>
  );
};

export const Chart = ({ config, onDelete }: Props) => {
  const [candlesticks, setCandlesticks] = useState<(CandlestickData<UTCTimestamp> | WhitespaceData<UTCTimestamp>)[]>([]);
  console.log('toto');
  const chartContainerRef = useRef<HTMLDivElement>(null);
  const chartRef = useRef<IChartApi | null>(null);
  const candleSeriesRef = useRef<ISeriesApi<'Candlestick'> | null>(null);

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

  return (
    <div style={{ height: '100%', display: 'flex', flexDirection: 'column' }} ref={el}>
      {isLoading && <div>Loading...</div>}
      {error && <div>Error loading chart data</div>}
      <Toolbar onDelete={onDelete} config={config} />
      <div ref={chartContainerRef} style={{ height: 'calc(100% - 2rem)' }} />
    </div>
  );
};
