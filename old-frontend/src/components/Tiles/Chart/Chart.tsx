import { Loading, Menu, MenuItem, useContextMenu } from '@carbon/react';
import { css } from '@emotion/css';
import { CandlestickData, ColorType, createChart, CrosshairMode, IChartApi, ISeriesApi, UTCTimestamp, WhitespaceData } from 'lightweight-charts';
import { useEffect, useRef, useState } from 'react';
import { GetCandlesticks } from '../../../../wailsjs/go/app/App';
import { app } from '../../../../wailsjs/go/models';
import { ClipboardSetText } from '../../../../wailsjs/runtime/runtime';
import { useAppData } from '../../../hooks/useAppData';
import { rangeToDates } from '../../../types/range';
import { ChartConfig } from '../../../types/tiles';
import { isIntraday } from '../../../types/timeframe';
import { Toolbar } from './Toolbar';

interface Props {
  onDelete: () => void;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

export const Chart = ({ config, onConfigChange, onDelete }: Props) => {
  const [candlesticks, setCandlesticks] = useState<(CandlestickData<UTCTimestamp> | WhitespaceData<UTCTimestamp>)[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [clickedPrice, setClickedPrice] = useState<number>(0);

  const chartContainerRef = useRef<HTMLDivElement>(null);
  const chartRef = useRef<IChartApi | null>(null);
  const candleSeriesRef = useRef<ISeriesApi<'Candlestick'> | null>(null);

  const el = useRef<HTMLDivElement | null>(null);
  const menuProps = useContextMenu(el);

  const { appData } = useAppData();

  const isDarkMode = appData.preferences.generalPreferences.theme === 'dark';

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

  const colors = isDarkMode ? darkColors : lightColors;

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

    // Ajouter un gestionnaire pour le clic droit
    const handleRightClick = (event: MouseEvent) => {
      if (event.button === 2 && chartContainerRef.current && chartRef.current) {
        const rect = chartContainerRef.current.getBoundingClientRect();
        const y = event.clientY - rect.top;

        const price = candleSeriesRef.current?.coordinateToPrice(y);

        setClickedPrice(price ? Math.round(price * 100) / 100 : 0);
      }
    };

    // Ajouter l'écouteur d'événement au conteneur du graphique
    chartContainerRef.current.addEventListener('mousedown', handleRightClick);

    return () => {
      resizeObserver.disconnect();
      chart.remove();
      if (chartContainerRef.current) {
        chartContainerRef.current.removeEventListener('mousedown', handleRightClick);
      }
    };
  }, []);

  useEffect(() => {
    const timeout = setTimeout(() => {
      if (isLoading) {
        setIsLoading(false);
        console.error('Failed to load data');
      }
    }, 5000);
    const fetchData = async () => {
      const { start, end } = rangeToDates(config.range);

      const params: app.GetCandlesticksConfig = {
        Ticker: config.ticker,
        Start: start.toISOString(),
        End: end.toISOString(),
        timeframe: config.timeframe,
      } as app.GetCandlesticksConfig;
      const candles = await GetCandlesticks(params);

      const newData = candles.map((candle) => ({
        time: Math.floor(new Date(candle.t as string).getTime() / 1000) as UTCTimestamp,
        open: candle.o,
        high: candle.h,
        low: candle.l,
        close: candle.c,
      }));

      setCandlesticks(newData);
      setIsLoading(false);
    };
    setIsLoading(true);
    fetchData();
    return () => {
      clearTimeout(timeout);
    };
  }, [config]);

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

  const handleResetView = () => {
    if (chartRef.current) {
      chartRef.current.timeScale().resetTimeScale();
    }
  };

  const renderMenu = () => {
    return (
      <>
        <MenuItem label={`Copy price: ${clickedPrice}`} onClick={() => ClipboardSetText(clickedPrice.toString())} />
        <MenuItem label={`Reset view`} onClick={handleResetView} />
      </>
    );
  };

  return (
    <div className={styles.height100} ref={el}>
      {isLoading && <Loading />}
      <Toolbar onDelete={onDelete} config={config} onConfigChange={onConfigChange} />
      <div ref={chartContainerRef} className={styles.chartContainer} />
      <Menu label="Tile menu" {...menuProps} mode="full">
        {renderMenu()}
      </Menu>
    </div>
  );
};

const styles = {
  loader: css`
    margin-top: 20%;
  `,
  chartContainer: css`
    height: calc(100% - 2rem);
  `,
  height100: css`
    height: 100%;
    display: flex;
    flex-direction: column;
  `,
};
