import { CandlestickData, createChart, IChartApi, ISeriesApi, Time } from 'lightweight-charts';
import { useEffect, useRef } from 'react';
import { VolumeProfile } from './VolumeProfile';

interface VolumeProfileItem {
  price: number;
  vol: number;
}

interface VolumeProfileData {
  time: Time;
  profile: VolumeProfileItem[];
  width: number;
}

const generateMockCandlestickData = (): CandlestickData[] => {
  const data: CandlestickData[] = [];
  const startPrice = 100;
  let currentPrice = startPrice;

  for (let i = 0; i < 100; i++) {
    const open = currentPrice;
    const close = open + (Math.random() - 0.5) * 10;
    const high = Math.max(open, close) + Math.random() * 5;
    const low = Math.min(open, close) - Math.random() * 5;
    data.push({
      time: (i + 1) as Time,
      open,
      high,
      low,
      close,
    });
    currentPrice = close;
  }

  return data;
};

const generateMockVolumeProfileData = (basePrice: number): VolumeProfileItem[] => {
  const priceStep = Math.round(basePrice * 0.1);
  const profile: VolumeProfileItem[] = [];
  for (let i = 0; i < 15; i++) {
    profile.push({
      price: basePrice + i * priceStep,
      vol: Math.round(Math.random() * 20),
    });
  }
  return profile;
};

export const Graph = () => {
  const chartContainerRef = useRef<HTMLDivElement | null>(null);
  const chartRef = useRef<IChartApi | null>(null);
  const candlestickSeriesRef = useRef<ISeriesApi<'Candlestick'> | null>(null);

  useEffect(() => {
    if (!chartContainerRef.current) return;

    // Create the chart
    const chart = createChart(chartContainerRef.current, {
      width: chartContainerRef.current.offsetWidth,
      height: 400,
    });
    chartRef.current = chart;

    // Add candlestick series
    const candlestickSeries = chart.addCandlestickSeries();
    candlestickSeriesRef.current = candlestickSeries;

    // Set mock candlestick data
    const candlestickData = generateMockCandlestickData();
    candlestickSeries.setData(candlestickData);

    // Generate and add volume profile
    const basePrice = candlestickData[candlestickData.length - 50].close;
    const profileData = generateMockVolumeProfileData(basePrice);
    const volumeProfileData: VolumeProfileData = {
      time: candlestickData[candlestickData.length - 50].time,
      profile: profileData,
      width: 10,
    };

    const volumeProfile = new VolumeProfile(chart, candlestickSeries, volumeProfileData);
    candlestickSeries.attachPrimitive(volumeProfile);

    // Clean up on component unmount
    return () => {
      chart.remove();
    };
  }, []);

  return <div ref={chartContainerRef} style={{ width: '100%', height: '100%' }} />;
};
