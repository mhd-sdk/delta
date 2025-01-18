import { ChartCandlestick, Locked, Settings, Switcher, Unlocked } from '@carbon/icons-react';
import { Header as Carbonheader, HeaderGlobalAction, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { NotificationInterface } from '../../types/notifications';
import { TileEnum } from '../../types/tiles';
import { Notifications } from '../Notifications/Notifications';
import { Separator } from '../Separator';

interface Props {
  isToolBoxOpen: boolean;
  toggleToolBox: () => void;
  isLayoutLocked: boolean;
  toggleLock: () => void;
  onNewTile: (type: TileEnum) => void;
  onOpenPreferences: () => void;
  notifications: NotificationInterface[];
}

export const Headerbar = ({
  isLayoutLocked,
  isToolBoxOpen,
  toggleLock,
  toggleToolBox,
  onNewTile,
  onOpenPreferences,
  notifications,
}: Props): JSX.Element => {
  const lockLabel = isLayoutLocked ? 'Unlock layout' : 'Lock layout';
  return (
    <Carbonheader aria-label="Delta">
      <div className={styles.ml(0.5)}>
        <HeaderGlobalAction aria-label="Tiles" tooltipAlignment="end" isActive={isToolBoxOpen} onClick={toggleToolBox}>
          <Switcher />
        </HeaderGlobalAction>
      </div>
      <Separator />
      <div className={styles.favorites}>
        <HeaderGlobalAction aria-label="Chart" tooltipAlignment="center" onClick={() => onNewTile(TileEnum.Chart)}>
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
          <OverflowMenuItem itemText="Preferences" onClick={onOpenPreferences} />
          <OverflowMenuItem itemText="Workspaces" />
          <OverflowMenuItem hasDivider itemText="Quit" isDelete />
        </OverflowMenu>
      </div>
    </Carbonheader>
  );
};

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
