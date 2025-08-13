import { Header } from '@/components/layout/header';
import { Main } from '@/components/layout/main';
import { ProfileDropdown } from '@/components/profile-dropdown';
import { Search } from '@/components/search';
import { ThemeSwitch } from '@/components/theme-switch';
import { useState } from 'react';
import { Grid } from './components/grid';
import { Panel, PanelType } from './panel';
import { Range, Unit } from './timerange';

export default function Dashboard() {
  const [panels, setPanels] = useState<Panel[]>([
    {
      data: {
        type: PanelType.Chart,
        config: {
          range: Range.oneDay,
          ticker: 'nvda',
          timeframe: {
            n: 1,
            unit: Unit.hour,
          },
        },
      },
      height: 5,
      width: 3,
      id: '1',
      x: 0,
      y: 0,
    },
    {
      data: {
        type: PanelType.Chart,
        config: {
          range: Range.oneDay,
          ticker: 'nvda',
          timeframe: {
            n: 1,
            unit: Unit.hour,
          },
        },
      },
      height: 5,
      width: 3,
      id: '2',
      x: 1,
      y: 0,
    },
    {
      data: {
        type: PanelType.Chart,
        config: {
          range: Range.oneDay,
          ticker: 'nvda',
          timeframe: {
            n: 1,
            unit: Unit.hour,
          },
        },
      },
      height: 5,
      width: 3,
      id: '3',
      x: 2,
      y: 0,
    },
  ]);
  console.log(panels);

  return (
    <>
      {/* ===== Top Heading ===== */}
      <Header>
        <div className="ml-auto flex items-center space-x-4">
          <Search />
          <ThemeSwitch />
          <ProfileDropdown />
        </div>
      </Header>

      {/* ===== Main ===== */}
      <Main className="h-screen">
        <div className="mb-2 flex items-center justify-between space-y-2">
          <h1 className="text-2xl font-bold tracking-tight">Dashboard</h1>
        </div>
        <div
          style={{
            transition: 'margin-left 0.3s',
            width: '100%',
            position: 'absolute',
          }}
        >
          <Grid panels={panels} onChange={setPanels} />
        </div>
      </Main>
    </>
  );
}
