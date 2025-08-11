import { Header } from '@/components/layout/header';
import { Main } from '@/components/layout/main';
import { ProfileDropdown } from '@/components/profile-dropdown';
import { Search } from '@/components/search';
import { ThemeSwitch } from '@/components/theme-switch';
import { useEffect, useState } from 'react';
import { Grid } from './components/grid';
import { Panel, PanelType } from './panel';
import { Range, Unit } from './timerange';

const PANELS_STORAGE_KEY = 'dashboard-panels';

const defaultPanels: Panel[] = [
  {
    data: {
      type: PanelType.Chart,
      config: {
        range: Range.oneDay,
        ticker: 'NVDA',
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
        ticker: 'NVDA',
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
        ticker: 'NVDA',
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
];

export default function Dashboard() {
  const [panels, setPanels] = useState<Panel[]>([]);

  // Load panels from localStorage on mount
  useEffect(() => {
    try {
      const savedPanels = localStorage.getItem(PANELS_STORAGE_KEY);
      if (savedPanels) {
        setPanels(JSON.parse(savedPanels));
      } else {
        setPanels(defaultPanels);
      }
    } catch (error) {
      console.error('Failed to load panels from localStorage:', error);
      setPanels(defaultPanels);
    }
  }, []);

  // Save panels to localStorage whenever they change
  useEffect(() => {
    if (panels.length > 0) {
      try {
        localStorage.setItem(PANELS_STORAGE_KEY, JSON.stringify(panels));
      } catch (error) {
        console.error('Failed to save panels to localStorage:', error);
      }
    }
  }, [panels]);

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
