import { Modal, TreeNode, TreeView } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { models } from '../../../wailsjs/go/models';
import { useAppData } from '../../hooks/useAppData';
import { Separator } from '../Separator';
import { Account } from './Account';
import { General } from './General';

interface Props {
  isOpen: boolean;
  onClose: () => void;
  onLogout: () => void;
}

enum PreferenceTabs {
  General = 'general',
  Account = 'account',
}

export const PreferenceModal = ({ isOpen, onClose, onLogout }: Props): JSX.Element => {
  const { appData, onSave } = useAppData();
  const [selected, setSelected] = useState(PreferenceTabs.General);
  const [isDirty, setIsDirty] = useState(false);
  const [generalPreferences, setGeneralPreferences] = useState<models.GeneralPreferences>({
    language: 'en',
    theme: 'dark',
  });

  useEffect(() => {
    const fetchAppData = async () => {
      setGeneralPreferences(appData.preferences.generalPreferences);
    };
    fetchAppData();
  }, [isOpen, appData]);

  const handleChangeGeneral = (value: models.GeneralPreferences) => {
    setIsDirty(true);
    setGeneralPreferences(value);
  };

  const handleSubmit = () => {
    onSave({
      ...appData,
      preferences: {
        ...appData.preferences,
        generalPreferences,
      },
    } as models.AppData);
    setIsDirty(false);
  };

  const handleCancel = async () => {
    setGeneralPreferences(appData.preferences.generalPreferences);
    setIsDirty(false);
    onClose();
  };

  const renderContent = (tab: PreferenceTabs) => {
    switch (tab) {
      case PreferenceTabs.General:
        return <General onChange={handleChangeGeneral} value={generalPreferences} />;
      case PreferenceTabs.Account:
        return <Account onLogout={onLogout} />;
    }
  };
  return (
    <Modal
      className={styles.modal}
      open={isOpen}
      onRequestClose={handleCancel}
      onRequestSubmit={handleSubmit}
      modalHeading="Preferences"
      primaryButtonText="Apply"
      primaryButtonDisabled={!isDirty}
      secondaryButtonText={isDirty ? 'Cancel' : 'Close'}
    >
      <div className={styles.modal}>
        <TreeView active={selected} selected={[selected]} label={''} className={styles.treeView}>
          <TreeNode id={PreferenceTabs.General} label="General" onClick={() => setSelected(PreferenceTabs.General)} />
          <TreeNode id={PreferenceTabs.Account} label="Account" onClick={() => setSelected(PreferenceTabs.Account)} />
        </TreeView>
        <Separator className={styles.divider} />
        {isOpen && <div className={styles.content}>{renderContent(selected)}</div>}
      </div>
    </Modal>
  );
};

const styles = {
  divider: css`
    height: 100%;
  `,
  treeView: css`
    width: 200px;
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
