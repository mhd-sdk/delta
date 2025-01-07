import { ChartCandlestick, Locked, Settings, Switcher, Unlocked } from '@carbon/icons-react';
import { Button, Content, GlobalTheme, Header, HeaderGlobalAction, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { ToastContainer, toast } from 'react-toastify';
import { AuthModal } from './components/AuthModal/AuthModal';
import { Grid } from './components/Grid/Grid';
import { Notifications } from './components/Notifications/Notifications';
import { PreferenceModal } from './components/PreferencesModal/PreferencesModal';
import { Separator } from './components/Separator';
import { NotificationInterface } from './types/notifications';
import { TileEnum, TileInterface } from './types/tiles';

const genId = (Tiles: TileInterface[]): string => {
  let id = 0;
  while (Tiles.some((tile) => tile.content.id === id.toString())) {
    id++;
  }
  return id.toString();
};

function App() {
  const [theme, setTheme] = useState<'light' | 'dark'>('dark');

  const themeCode = theme === 'light' ? 'g10' : 'g90';

  const [isToolBoxOpen, setIsToolBoxOpen] = useState(false);

  const [isPreferencesOpen, setIsPreferencesOpen] = useState(false);

  const [isDatafeedOpen, setIsDatafeedOpen] = useState(true);

  const [notifications] = useState<NotificationInterface[]>([{ title: 'notif 1', type: 'info', subtitle: 'subtitle' }]);

  const [isLayoutLocked, setIsLayoutLocked] = useState(false);

  const [tiles, setTiles] = useState<TileInterface[]>([
    {
      content: {
        id: '0',
        type: TileEnum.Chart,
      },
      x: 0,
      y: 0,
      w: 20,
      h: 20,
    },
  ]);

  const lockLabel = isLayoutLocked ? 'Unlock layout' : 'Lock layout';

  const handleToggleTheme = () => {
    const newTheme = theme === 'light' ? 'dark' : 'light';
    setTheme(newTheme);
  };

  const toggleLock = () => setIsLayoutLocked(!isLayoutLocked);

  const toggleToolBox = () => setIsToolBoxOpen(!isToolBoxOpen);

  const handleNewTile = (type: TileEnum) => {
    const newTile: TileInterface = {
      content: {
        id: genId(tiles),
        type,
      },
      x: 0,
      y: 0,
      w: 20,
      h: 20,
    };
    setTiles([...tiles, newTile]);
  };

  useEffect(() => {
    document.documentElement.dataset.carbonTheme = themeCode;
  }, [themeCode]);

  const notify = () => toast('Wow so easy !');

  return (
    <GlobalTheme theme={themeCode}>
      <div id="App">
        <Header aria-label="Delta">
          <Button onClick={notify}>Notify !</Button>
          <div className={styles.ml(0.5)}>
            <HeaderGlobalAction aria-label="Tiles" tooltipAlignment="end" isActive={isToolBoxOpen} onClick={toggleToolBox}>
              <Switcher />
            </HeaderGlobalAction>
          </div>
          <Separator />
          <div className={styles.favorites}>
            <HeaderGlobalAction aria-label="Chart" tooltipAlignment="center" onClick={() => handleNewTile(TileEnum.Chart)}>
              <ChartCandlestick />
            </HeaderGlobalAction>
          </div>
          <Separator />
          <div className={styles.rightActions}>
            <HeaderGlobalAction onClick={toggleLock} aria-label={lockLabel} tooltipAlignment="center">
              {isLayoutLocked ? <Locked /> : <Unlocked />}
            </HeaderGlobalAction>
            <Notifications notifications={notifications} />

            <OverflowMenu renderIcon={Settings} size="lg" flipped aria-label="overflow-menu">
              <OverflowMenuItem itemText="Preferences" onClick={() => setIsPreferencesOpen(true)} />
              <OverflowMenuItem itemText="Switch theme" onClick={handleToggleTheme} />
              <OverflowMenuItem itemText="Account infos" />
              <OverflowMenuItem hasDivider itemText="Disconnect" isDelete />
            </OverflowMenu>
          </div>
        </Header>

        <div className={styles.toolBox(isToolBoxOpen)}></div>

        <Content className={styles.content}>
          <Grid isLocked={isLayoutLocked} tiles={tiles} onChange={setTiles} />
        </Content>
        <PreferenceModal isOpen={isPreferencesOpen} setIsOpen={(isOpen) => setIsPreferencesOpen(isOpen)} />
        <AuthModal isOpen={isDatafeedOpen} setIsOpen={(isOpen) => setIsDatafeedOpen(isOpen)} />
        <ToastContainer />
      </div>
    </GlobalTheme>
  );
}

const styles = {
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
  toolBox: (isOpen: boolean) => css`
    margin-top: ${isOpen ? '3rem' : '0rem'};
    height: 3rem;
    transition: margin-top 0.3s;
    border-bottom: 1px solid #d0d0d084;
    padding: 0;
    background-color: #f4f4f4;
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
