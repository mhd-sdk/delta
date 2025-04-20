import { Header } from '@/components/layout/header';
import { Main } from '@/components/layout/main';
import { useState } from 'react';
import { TickerSelect } from '../components/chart/tickerSelect';

const ChartsPage = () => {
  const [selectedTicker, setSelectedTicker] = useState('AAPL');

  return (
    <>
      <Header>
        <div className="ml-auto flex items-center space-x-4">
          <div className="mr-4">
            <span className="font-medium mr-2">Current Ticker:</span>
            <TickerSelect value={selectedTicker} onChange={setSelectedTicker} />
          </div>
        </div>
      </Header>

      <Main>
        <div className="mb-4 flex items-center justify-between">
          <h1 className="text-2xl font-bold tracking-tight">Charts Dashboard</h1>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="p-6 bg-card rounded-lg border shadow-sm">
            <h2 className="text-xl font-semibold mb-4">{selectedTicker} Chart</h2>
            <div className="h-64 bg-muted/20 rounded flex items-center justify-center">Chart visualization will be displayed here</div>
          </div>

          <div className="p-6 bg-card rounded-lg border shadow-sm">
            <h2 className="text-xl font-semibold mb-4">{selectedTicker} Performance</h2>
            <div className="h-64 bg-muted/20 rounded flex items-center justify-center">Performance metrics will be displayed here</div>
          </div>

          <div className="p-6 bg-card rounded-lg border shadow-sm">
            <h2 className="text-xl font-semibold mb-4">{selectedTicker} Details</h2>
            <div className="h-64 bg-muted/20 rounded flex items-center justify-center">Ticker details will be displayed here</div>
          </div>
        </div>
      </Main>
    </>
  );
};

export default ChartsPage;
