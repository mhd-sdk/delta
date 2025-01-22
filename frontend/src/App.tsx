import { Content, GlobalTheme, Modal } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { ToastContainer } from 'react-toastify';
import { Logout, SaveAppData } from '../wailsjs/go/app/App';
import { persistence } from '../wailsjs/go/models';
import { AuthModal } from './components/AuthModal/AuthModal';
import { Grid } from './components/Grid/Grid';
import { Headerbar } from './components/Headerbar/Headerbar';
import { PreferenceModal } from './components/PreferencesModal/PreferencesModal';
import { useAppData } from './hooks/useAppData';
import { NotificationInterface } from './types/notifications';
import { TileInterface, TileType } from './types/tiles';
import { calcOptimizedRange, defaultTimeframes } from './types/timeframe';
import { getThemeCode } from './utils/getThemeCode';

const genId = (Tiles: TileInterface[]): string => {
  let id = 0;
  while (Tiles.some((tile) => tile.id === id.toString())) {
    id++;
  }
  return id.toString();
};

function App() {
  const { appData } = useAppData();
  const theme = appData.preferences.generalPreferences.theme;
  const themeCode = getThemeCode(theme);

  const [isWorkspaceDirty, setIsWorkspaceDirty] = useState(false);

  const [isToolBoxOpen, setIsToolBoxOpen] = useState(false);

  const [isPreferencesOpen, setIsPreferencesOpen] = useState(false);

  const [isDatafeedOpen, setIsDatafeedOpen] = useState(false);

  const [notifications] = useState<NotificationInterface[]>([{ title: 'notif 1', type: 'info', subtitle: 'subtitle' }]);

  const [isLogoutModalOpen, setIsLogoutModalOpen] = useState(false);

  const [tiles, setTiles] = useState<TileInterface[]>([
    {
      id: '0',
      data: {
        type: TileType.Chart,
        config: {
          range: calcOptimizedRange(defaultTimeframes[4]),
          ticker: 'AAPL',
          timeframe: defaultTimeframes[4],
        },
      },
      x: 0,
      y: 0,
      w: 20,
      h: 20,
    },
  ]);

  const toggleToolBox = () => setIsToolBoxOpen(!isToolBoxOpen);

  const handleNewTile = (type: TileType) => {
    switch (type) {
      case TileType.Chart:
        setTiles([
          ...tiles,
          {
            id: genId(tiles),
            data: {
              type,
              config: {
                ticker: 'AAPL',
                timeframe: defaultTimeframes[4],
                range: calcOptimizedRange(defaultTimeframes[4]),
              },
            },
            x: 0,
            y: 0,
            w: 20,
            h: 20,
          },
        ]);
    }
  };

  const handleLogout = async () => {
    setIsLogoutModalOpen(false);
    setIsPreferencesOpen(false);
    await SaveAppData({
      ...appData,
      keys: {
        apiKey: '',
        secretKey: '',
      },
    } as persistence.AppData);
    setIsDatafeedOpen(true);
    Logout();
  };

  useEffect(() => {
    document.documentElement.dataset.carbonTheme = themeCode;
  }, [themeCode]);

  useEffect(() => {
    setIsWorkspaceDirty(true);
  }, [tiles]);

  // const notify = () => toast('Wow so easy !');

  return (
    <GlobalTheme theme={themeCode}>
      <div id="App" className={styles.app(theme)}>
        <Headerbar
          isWorkspaceDirty={isWorkspaceDirty}
          onSaveWorkspace={() => setIsWorkspaceDirty(false)}
          isToolBoxOpen={isToolBoxOpen}
          toggleToolBox={toggleToolBox}
          onNewTile={handleNewTile}
          onOpenPreferences={() => setIsPreferencesOpen(true)}
          notifications={notifications}
        />

        <div className={styles.toolBox(isToolBoxOpen, theme)}></div>

        <Content className={styles.content}>
          <Grid tiles={tiles} onChange={setTiles} />
        </Content>

        <PreferenceModal onLogout={() => setIsLogoutModalOpen(true)} isOpen={isPreferencesOpen} onClose={() => setIsPreferencesOpen(false)} />

        <AuthModal isOpen={isDatafeedOpen} setIsOpen={(isOpen) => setIsDatafeedOpen(isOpen)} />

        <Modal
          open={isLogoutModalOpen}
          onRequestClose={() => setIsLogoutModalOpen(false)}
          danger
          onRequestSubmit={handleLogout}
          modalHeading="Are you sure you want to logout ?"
          modalLabel="This action is irreversible"
          primaryButtonText="Logout"
          secondaryButtonText="Cancel"
        />

        <ToastContainer />
      </div>
    </GlobalTheme>
  );
}

const styles = {
  app: (theme: string) => css`
    scrollbar-color: ${theme === 'light' ? '#d0d0d0 #fff' : '#606060 #303030'};
  `,
  layout: css`
    height: calc(100vh - 3rem);
  `,
  content: css`
    transition: margin-left 0.3s;
    width: 100%;
    height: calc(100vh - 3rem);
    overflow-x: hidden;
    overflow-y: auto;
  `,
  toolBox: (isOpen: boolean, theme: string) => css`
    margin-top: ${isOpen ? '3rem' : '0rem'};
    height: 3rem;
    transition: margin-top 0.3s;
    border-bottom: 1px solid ${theme === 'light' ? '#d0d0d084' : '#333333'};
    padding: 0;
    background-color: ${theme === 'light' ? '#f4f4f4' : '#393939'};
  `,
  header: css``,
  ml: (number: number) => css`
    margin-left: ${number}rem;
  `,
  mr: (number: number) => css`
    margin-right: ${number}rem;
  `,
  favorites: css`
    display: flex;
    align-items: center;
    margin-left: auto;
    flex-grow: 1;
  `,
  rightActions: css`
    display: flex;
    align-items: center;
    margin-right: 0.5rem;
  `,
  settingsMenu: css`
    position: fixed !important;
    right: 0 !important;
    opacity: 0.5 !important;
  `,
};

export default App;
