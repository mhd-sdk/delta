import { Edit, TrashCan } from '@carbon/icons-react';
import { Button, ContainedList, ContainedListItem, Modal, Search } from '@carbon/react';
import { css } from '@emotion/css';
import { createRef, useState } from 'react';
import { models } from '../../../wailsjs/go/models';
import { useAppData } from '../../hooks/useAppData';
import { EditableRow } from './EditableRow';

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

export const WorkspaceModal = ({ isOpen, onClose }: Props): JSX.Element => {
  const [workspaces, setWorkspaces] = useState<string[]>(['w1', 'w2', 'w3']);
  const [searchTerm, setSearchTerm] = useState('');
  const [editing, setEditing] = useState<string>();
  const [selected, setSelected] = useState<string>();
  const editRef = createRef<HTMLInputElement>();

  const handleEditWorkspace = (idx: number, newWorkspace: string) => {
    const newWorkspaces = [...workspaces];
    newWorkspaces[idx] = newWorkspace;
    setWorkspaces(newWorkspaces);
  };

  const { appData, onSave } = useAppData();
  const theme = appData.preferences.generalPreferences.theme;

  const results = workspaces.filter((item) => item.toLowerCase().includes(searchTerm.toLowerCase()));

  const handleSubmit = () => {
    onSave({
      ...appData,
    } as models.AppData);
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
  return (
    <Modal
      className={styles.modal}
      open={isOpen}
      onRequestClose={handleCancel}
      onRequestSubmit={handleSubmit}
      modalHeading="Workspaces"
      primaryButtonText="Open"
      primaryButtonDisabled={selected === undefined}
      secondaryButtonText={'Close'}
    >
      <ContainedList
        kind="on-page"
        action={
          <div className={styles.topbar}>
            <Search id="search-default-1" labelText={undefined} size="lg" />
            <Button size="lg">New workspace</Button>
          </div>
        }
        label={undefined}
      >
        {results.map((listItem, key) =>
          editing !== listItem ? (
            <ContainedListItem
              className={styles.item(isSelected(listItem), theme === 'dark')}
              key={key}
              onClick={() => {
                setSelected(listItem);
              }}
              action={
                <div className={styles.actions}>
                  <Button
                    onClick={() => handleOpenEditing(listItem)}
                    kind="ghost"
                    iconDescription="Edit"
                    hasIconOnly
                    renderIcon={Edit}
                    aria-label="Edit"
                  />
                  <Button kind="danger--ghost" iconDescription="Delete" hasIconOnly renderIcon={TrashCan} aria-label="Delete" />
                </div>
              }
            >
              <div>{listItem}</div>
            </ContainedListItem>
          ) : (
            <EditableRow
              key={key}
              previousName={listItem}
              onSave={(value) => {
                handleEditWorkspace(key, value);
                setEditing(undefined);
              }}
              onCancel={() => setEditing(undefined)}
            />
          )
        )}
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
    width: 100%;
    @media (min-width: 768px) {
      & .cds--modal-container {
        width: 75% !important;
      }
    }

    & .cds--modal-container {
      height: 100% !important;
    }

    & .cds--modal-content {
      padding-bottom: 1rem !important;
    }
  `,
};
