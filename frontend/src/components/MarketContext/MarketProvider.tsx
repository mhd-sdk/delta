import { createContext, PropsWithChildren, useEffect, useState } from 'react';

export interface MarketMockParams {
  averageBuy: number; // Average volume for buy orders
  buyFrequency: number; // Frequency of buy orders (milliseconds)
  averageSell: number; // Average volume for sell orders
  sellFrequency: number; // Frequency of sell orders (milliseconds)
  duration: number; // x last seconds to analyze
}

export interface Order {
  side: OrderSide;
  size: number;
  time: Date;
  price: number;
}
export type Tape = Order[];

export enum OrderSide {
  Buy = 'buy',
  Sell = 'sell',
}

function generateRandomVolume(mean: number, deviation: number): number {
  const randomOffset = Math.random() * deviation * 2 - deviation;
  return mean + randomOffset;
}

export interface IMarketContext {
  tape: Tape;
  settings: {
    averageBuy: number;
    buyFrequency: number;
    averageSell: number;
    sellFrequency: number;
    setAverageBuy: (value: number) => void;
    setBuyFrequency: (value: number) => void;
    setAverageSell: (value: number) => void;
    setSellFrequency: (value: number) => void;
  };
}

export const MarketContext = createContext({} as IMarketContext);

const generateRandomOrder = (side: OrderSide, averageSize: number): Order => {
  const size = generateRandomVolume(averageSize, 2);
  const price = Math.random() * 100;
  const order: Order = {
    side,
    size,
    time: new Date(),
    price,
  };
  return order;
};

export const MarketProvider = ({ children }: PropsWithChildren): JSX.Element => {
  const [averageBuy, setAverageBuy] = useState(50);
  const [buyFrequency, setBuyFrequency] = useState(1000);
  const [averageSell, setAverageSell] = useState(50);
  const [sellFrequency, setSellFrequency] = useState(1000);

  const [tape, setTape] = useState<Tape>([]);

  useEffect(() => {
    // Mock buy orders
    const buyInterval = setInterval(() => {
      const buyOrder = generateRandomOrder(OrderSide.Buy, averageBuy);
      setTape((prevTape) => [...prevTape, buyOrder]);
    }, buyFrequency);

    // Mock sell orders
    const sellInterval = setInterval(() => {
      const sellOrder = generateRandomOrder(OrderSide.Sell, averageSell);
      setTape((prevTape) => [...prevTape, sellOrder]);
    }, sellFrequency);

    return () => {
      clearInterval(buyInterval);
      clearInterval(sellInterval);
    };
  }, [averageBuy, buyFrequency, averageSell, sellFrequency]);

  const value: IMarketContext = {
    tape,
    settings: {
      averageBuy,
      buyFrequency,
      averageSell,
      sellFrequency,
      setAverageBuy,
      setAverageSell,
      setBuyFrequency,
      setSellFrequency,
    },
  };

  return <MarketContext.Provider value={value}>{children}</MarketContext.Provider>;
};
