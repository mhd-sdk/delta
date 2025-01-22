import { ChartCandlestick, Save, Settings, Switcher } from '@carbon/icons-react';
import { Header as Carbonheader, HeaderGlobalAction, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { NotificationInterface } from '../../types/notifications';
import { TileType } from '../../types/tiles';
import { Notifications } from '../Notifications/Notifications';
import { Separator } from '../Separator';

interface Props {
  isToolBoxOpen: boolean;
  toggleToolBox: () => void;
  onNewTile: (type: TileType) => void;
  onOpenPreferences: () => void;
  notifications: NotificationInterface[];
  isWorkspaceDirty: boolean;
  onSaveWorkspace: () => void;
}

export const Headerbar = ({
  isToolBoxOpen,
  toggleToolBox,
  onNewTile,
  onOpenPreferences,
  notifications,
  isWorkspaceDirty,
  onSaveWorkspace,
}: Props): JSX.Element => {
  return (
    <Carbonheader aria-label="Delta">
      <div className={styles.ml(0.5)}>
        <HeaderGlobalAction aria-label="Tiles" tooltipAlignment="end" isActive={isToolBoxOpen} onClick={toggleToolBox}>
          <Switcher />
        </HeaderGlobalAction>
      </div>
      <Separator />
      <div className={styles.favorites}>
        <HeaderGlobalAction aria-label="Chart" tooltipAlignment="center" onClick={() => onNewTile(TileType.Chart)}>
          <ChartCandlestick />
        </HeaderGlobalAction>
      </div>
      <Separator />
      <div className={styles.rightActions}>
        <HeaderGlobalAction
          className={styles.saveAction(isWorkspaceDirty)}
          aria-label="Save workspace"
          onClick={onSaveWorkspace}
          tooltipAlignment="center"
        >
          <Save />
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
  saveAction: (isDirty: boolean) => css`
    color: ${isDirty ? '#f4f4f4' : '#161616'};
    opacity: ${isDirty ? 1 : 0.3};
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
