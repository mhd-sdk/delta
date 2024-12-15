import { Settings, Switcher } from '@carbon/icons-react';
import { Content, Header, HeaderGlobalAction, HeaderGlobalBar, HeaderName, PopoverContent, SkipToContent, Tile } from '@carbon/react';
import { css } from '@emotion/css';
import { useState } from 'react';
import GridLayout from 'react-grid-layout';
import { Notification, Notifications } from './components/Notifications/Notifications';
import { Separator } from './components/Separator';

function App() {
  const [isToolBoxOpen, setIsToolBoxOpen] = useState(false);

  const [notifications, setNotifications] = useState<Notification[]>([{ title: 'notif 1', type: 'info', subtitle: 'subtitle' }]);

  const layout: GridLayout.Layout[] = [
    { i: 'a', x: 0, y: 0, w: 1, h: 2, static: true },
    { i: 'b', x: 1, y: 0, w: 3, h: 2, minW: 2, maxW: 4 },
    { i: 'c', x: 4, y: 0, w: 1, h: 2 },
  ];
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
      <Tile className={styles.toolBox(isToolBoxOpen)}>tiles</Tile>

      <Content className={styles.content}>
        <GridLayout className="layout" layout={layout} cols={12} rowHeight={30} width={1200}>
          <div key="a">a</div>
          <div key="b">b</div>
          <div key="c">c</div>
        </GridLayout>
      </Content>
    </div>
  );
}

const styles = {
  content: css`
    transition: margin-left 0.3s;
  `,
  toolBox: (isOpen: boolean) => css`
    margin-top: ${isOpen ? '3rem' : '-4rem'};
    transition: margin-top 0.3s;
    border-bottom: 1px solid #d0d0d084;
  `,
  notifs: (isOpen: boolean) => css`
    margin-top: 3rem;
    position: absolute;
    right: 0;
    margin-right: ${isOpen ? '0rem' : '-20rem'};
    width: 20rem;
    transition: margin-right 0.3s;
    height: calc(100% - 3rem);
  `,
  header: css``,
  headerActions: css`
    display: flex;
    align-items: center;
  `,
};

export default App;
