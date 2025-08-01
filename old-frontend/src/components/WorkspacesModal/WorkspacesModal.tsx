import { Edit, TrashCan } from '@carbon/icons-react';
import { Button, ContainedList, ContainedListItem, Modal } from '@carbon/react';
import { css } from '@emotion/css';
import { createRef, useEffect, useState } from 'react';
import { models } from '../../../wailsjs/go/models';
import { useAppData } from '../../hooks/useAppData';
import { EditableRow } from './EditableRow';

interface Props {
  isOpen: boolean;
  onClose: () => void;
  onSelect: (workspace: string) => void;
  selectedWorkspace: string;
}

export const WorkspacesModal = ({ isOpen, onClose, onSelect, selectedWorkspace }: Props): JSX.Element => {
  const { appData, onSave } = useAppData();
  const [searchTerm, setSearchTerm] = useState('');
  const [editing, setEditing] = useState<string>();
  const [creating, setCreating] = useState<boolean>(false);
  const [selected, setSelected] = useState<string>();
  const editRef = createRef<HTMLInputElement>();

  const handleEditWorkspace = (idx: number, newWorkspace: string) => {
    onSave({
      ...appData,
      workspaces: appData.workspaces.map((w, i) => (i === idx ? { ...w, name: newWorkspace } : w)),
    } as models.AppData);
  };

  const handleDeleteWorkspace = (name: string) => {
    onSave({
      ...appData,
      workspaces: appData.workspaces.filter((w) => w.name !== name),
    } as models.AppData);
  };

  const handleCreateWorkspace = (name: string) => {
    onSave({
      ...appData,
      workspaces: [
        ...appData.workspaces,
        {
          name,
          tiles: [],
          layout: [],
        },
      ],
    } as models.AppData);
    setCreating(false);
  };

  const theme = appData.preferences.generalPreferences.theme;

  const results = appData.workspaces.filter((item) => item.name.toLowerCase().includes(searchTerm.toLowerCase()));

  const handleSelect = () => {
    if (selected) {
      onSelect(selected);
      onClose();
    }
  };

  const handleCancel = async () => {
    onClose();
  };

  const isSelected = (item: string) => item === selected;

  const handleOpenEditing = (item: string) => {
    setEditing(item);
    setSelected(undefined);
    // focus
    setTimeout(() => {
      editRef.current?.focus();
    }, 100);
  };

  useEffect(() => {
    if (isOpen) {
      setSelected(undefined);
    }
  }, [isOpen]);

  return (
    <Modal
      className={styles.modal}
      open={isOpen}
      onRequestClose={handleCancel}
      onRequestSubmit={handleSelect}
      modalHeading="Workspaces"
      primaryButtonText="Open"
      primaryButtonDisabled={selected === undefined}
      secondaryButtonText={'Close'}
    >
      <ContainedList
        action={
          <div className={styles.topbar}>
            <Button size="lg" onClick={() => setCreating(true)}>
              New workspace
            </Button>
          </div>
        }
        label=""
      >
        {results.map((wp, key) =>
          editing !== wp.name ? (
            <ContainedListItem
              disabled={selectedWorkspace === wp.name}
              className={styles.item(isSelected(wp.name), theme === 'dark')}
              key={key}
              onClick={() => {
                setSelected(wp.name);
              }}
              action={
                <div className={styles.actions}>
                  <Button
                    onClick={() => handleOpenEditing(wp.name)}
                    kind="ghost"
                    iconDescription="Edit"
                    hasIconOnly
                    renderIcon={Edit}
                    aria-label="Edit"
                    disabled={selectedWorkspace === wp.name}
                  />
                  <Button
                    kind="danger--ghost"
                    iconDescription="Delete"
                    hasIconOnly
                    renderIcon={TrashCan}
                    aria-label="Delete"
                    disabled={selectedWorkspace === wp.name}
                    onClick={() => handleDeleteWorkspace(wp.name)}
                  />
                </div>
              }
            >
              <div>
                {wp.name} {selectedWorkspace === wp.name && '(Currently selected)'}
              </div>
            </ContainedListItem>
          ) : (
            <EditableRow
              uniqueMode="edit"
              key={key}
              previousName={wp.name}
              onSave={(value) => {
                handleEditWorkspace(key, value);
                setEditing(undefined);
              }}
              onCancel={() => setEditing(undefined)}
            />
          )
        )}
        {creating && <EditableRow uniqueMode="new" previousName="New Workspace" onSave={handleCreateWorkspace} onCancel={() => setCreating(false)} />}
      </ContainedList>
    </Modal>
  );
};

const styles = {
  topbar: css`
    display: flex;
    flex-direction: row;
    width: 100%;
  `,
  item: (isSelected: boolean, isDarkMode: boolean) => css`
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
    ${isSelected && `background-color: ${isDarkMode ? '#606060' : '#e4e4e4'};`}
  `,
  actions: css`
    z-index: 1;
    margin-left: auto;
    display: flex;
  `,
  divider: css`
    height: 100%;
  `,
  content: css`
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    padding: 1rem;
  `,
  modal: css`
    display: flex;
    flex-direction: row;
    height: 100%;
    @media (min-width: 768px) {
      & .cds--modal-container {
        width: 55% !important;
      }
    }

    & .cds--modal-container {
      height: 60% !important;
    }

    & .cds--modal-content {
      padding-bottom: 1rem !important;
    }
  `,
};
