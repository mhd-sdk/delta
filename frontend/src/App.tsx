import { Settings, Switcher } from '@carbon/icons-react';
import { Content, Header, HeaderGlobalAction, HeaderGlobalBar, HeaderName, PopoverContent, SkipToContent } from '@carbon/react';
import { css } from '@emotion/css';
import { useState } from 'react';
import { Notification, Notifications } from './components/Notifications/Notifications';
import { Separator } from './components/Separator';
import { Grid } from './components/Grid/Grid';

function App() {
  const [isToolBoxOpen, setIsToolBoxOpen] = useState(false);

  const [notifications, setNotifications] = useState<Notification[]>([{ title: 'notif 1', type: 'info', subtitle: 'subtitle' }]);

  const [tiles, setTiles] = useState<{ id: string; content: JSX.Element; x: number; y: number; w: number; h: number }[]>([
    { id: '1', content: <div>1</div>, x: 0, y: 0, w: 20, h: 20 },
    { id: '2', content: <div>2</div>, x: 20, y: 0, w: 20, h: 20 },
    { id: '3', content: <div>3</div>, x: 40, y: 0, w: 20, h: 20 },
    { id: '4', content: <div>4</div>, x: 60, y: 0, w: 20, h: 20 },
  ]);

  return (
    <div
      id="App"
      className={css`
        display: flex;
        flex-direction: column;
        height: 100vh;
        overflow: none;
      `}
    >
      <Header aria-label="Delta">
        <SkipToContent />

        <HeaderName prefix="" href="#" className={styles.header}>
          DeltΔ
        </HeaderName>
        <Separator />
        <div className={styles.headerActions}>
          <HeaderGlobalAction aria-label="Tiles" isActive={isToolBoxOpen} onClick={() => setIsToolBoxOpen(!isToolBoxOpen)}>
            <Switcher size={20} />
          </HeaderGlobalAction>
          <PopoverContent className="p-3">coucou</PopoverContent>
        </div>

        <HeaderGlobalBar>
          <Notifications notifications={notifications} />

          <HeaderGlobalAction aria-label="Settings" tooltipAlignment="end">
            <Settings size={20} />
          </HeaderGlobalAction>
        </HeaderGlobalBar>
      </Header>
      <div className={styles.toolBox(isToolBoxOpen)}>tiles</div>

      <Content className={styles.content}>
        <Grid tiles={tiles} />
      </Content>
    </div>
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
    background-color: #f4f4f4;
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
  headerActions: css`
    display: flex;
    align-items: center;
  `,
};

export default App;
