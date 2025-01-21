import { Loading } from '@carbon/react';
import { css } from '@emotion/css';
import { CandlestickData, ColorType, createChart, CrosshairMode, IChartApi, ISeriesApi, UTCTimestamp, WhitespaceData } from 'lightweight-charts';
import { useEffect, useRef, useState } from 'react';
import { GetCandlesticks } from '../../../../wailsjs/go/app/App';
import { app } from '../../../../wailsjs/go/models';
import { useAppData } from '../../../hooks/useAppData';
import { ChartConfig, Range, Timeframe } from '../../../types/tiles';
import { Toolbar } from './Toolbar';

interface Props {
  onDelete: () => void;
  config: ChartConfig;
  setClickedPrice: (price: number) => void;
  onConfigChange: (config: ChartConfig) => void;
}

export const Chart = ({ config, onConfigChange, setClickedPrice, onDelete }: Props): JSX.Element => {
  const [candlesticks, setCandlesticks] = useState<(CandlestickData<UTCTimestamp> | WhitespaceData<UTCTimestamp>)[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const chartContainerRef = useRef<HTMLDivElement>(null);
  const chartRef = useRef<IChartApi | null>(null);
  const candleSeriesRef = useRef<ISeriesApi<'Candlestick'> | null>(null);

  const { appData } = useAppData();

  const isDarkMode = appData.preferences.generalPreferences.theme === 'dark';

  const darkColors = {
    backgroundColor: 'transparent',
    textColor: 'white',
    candleUpColor: '#ffffff', // Couleur des chandeliers haussiers
    candleDownColor: '#5d606b', // Couleur des chandeliers baissiers
  };

  const lightColors = {
    backgroundColor: 'transparent',
    textColor: 'black',
    candleUpColor: '#4caf50', // Couleur des chandeliers haussiers
    candleDownColor: '#f44336', // Couleur des chandeliers baissiers
  };

  const colors = isDarkMode ? darkColors : lightColors;

  useEffect(() => {
    if (!chartContainerRef.current) {
      return;
    }

    const chart = createChart(chartContainerRef.current, {
      timeScale: {
        timeVisible:
          config.timeframe === Timeframe.oneMin ||
          config.timeframe === Timeframe.fiveMin ||
          config.timeframe === Timeframe.thirtyMin ||
          config.timeframe === Timeframe.oneHour ||
          config.timeframe === Timeframe.fourHour,
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
        // round to 2 decimal places

        setClickedPrice(price ? Math.round(price * 100) / 100 : 0);
        console.log('Prix cliqué (clic droit) :', price);
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
  }, [config]);

  useEffect(() => {
    const fetchData = async () => {
      const actualDate = new Date();
      const start = new Date();
      switch (config.range) {
        case Range.oneDay:
          start.setDate(actualDate.getDate() - 1); // 1 jour en arrière par rapport à la date du jour
          break;
        case Range.threeDays:
          start.setDate(actualDate.getDate() - 3); // 3 jours en arrière par rapport à la date du jour
          break;
        case Range.oneWeek:
          start.setDate(actualDate.getDate() - 7); // 1 semaine en arrière par rapport à la date du jour
          break;
        case Range.oneMonth:
          start.setMonth(actualDate.getMonth() - 1); // 1 mois en arrière par rapport à la date du jour
          break;
        case Range.threeMonths:
          start.setMonth(actualDate.getMonth() - 3); // 3 mois en arrière par rapport à la date du jour
          break;
        case Range.sixMonths:
          start.setMonth(actualDate.getMonth() - 6); // 6 mois en arrière par rapport à la date du jour
          break;
        case Range.oneYear:
          start.setFullYear(actualDate.getFullYear() - 1); // 1 an en arrière par rapport à la date du jour
          break;
        case Range.fiveYear:
          start.setFullYear(actualDate.getFullYear() - 5); // 5 ans en arrière par rapport à la date du jour
          break;
        default:
          break;
      }
      let tf: app.TimeFrame = {
        N: 1,
        Unit: 'Min',
      };
      switch (config.timeframe) {
        case Timeframe.oneMin:
          tf = {
            N: 1,
            Unit: 'Min',
          };
          break;
        case Timeframe.fiveMin:
          tf = {
            N: 5,
            Unit: 'Min',
          };
          break;
        case Timeframe.fifteenMin:
          tf = {
            N: 15,
            Unit: 'Min',
          };
          break;
        case Timeframe.thirtyMin:
          tf = {
            N: 30,
            Unit: 'Min',
          };
          break;
        case Timeframe.oneHour:
          tf = {
            N: 1,
            Unit: 'Hour',
          };
          break;
        case Timeframe.fourHour:
          tf = {
            N: 4,
            Unit: 'Hour',
          };
          break;
        case Timeframe.oneDay:
          tf = {
            N: 1,
            Unit: 'Day',
          };
          break;
        case Timeframe.oneWeek:
          tf = {
            N: 1,
            Unit: 'Week',
          };
          break;
        case Timeframe.oneMonth:
          tf = {
            N: 1,
            Unit: 'Month',
          };
          break;
        default:
          break;
      }
      const params: app.GetCandlesticksConfig = {
        Ticker: config.ticker,
        Start: start,
        End: actualDate.toISOString(),
        timeframe: {
          N: tf.N,
          Unit: tf.Unit,
        },
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

  return (
    <div className={styles.height100}>
      <Toolbar onDelete={onDelete} config={config} onConfigChange={onConfigChange} />
      {isLoading && <Loading />}
      <div ref={chartContainerRef} className={styles.chartContainer} />
    </div>
  );
};

const styles = {
  loader: css`
    /* vertically center */
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
