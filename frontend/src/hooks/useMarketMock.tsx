import { useContext } from 'react';
import { IMarketContext, MarketContext } from '../components/MarketContext/MarketProvider';

export const useMarketMock = (): IMarketContext => {
  const marketMock = useContext(MarketContext);
  return marketMock;
};
