import { Modal, TreeNode, TreeView } from '@carbon/react';
import { css } from '@emotion/css';
import { useState } from 'react';
import { AppearanceSettings, GeneralSettings } from '../../types/preferences';
import { Separator } from '../Separator';
import { Appearance } from './Appearance';
import { General } from './General';

interface Props {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

enum PreferenceTabs {
  General = 'general',
  Appearance = 'appearance',
}

export const PreferenceModal = ({ isOpen, setIsOpen }: Props): JSX.Element => {
  //   const menuTargetref = useRef(null);
  const [selected, setSelected] = useState(PreferenceTabs.General);

  const [generalSettings, setGeneralSettings] = useState<GeneralSettings>({} as GeneralSettings);
  const [appearanceSettings, setAppearanceSettings] = useState<AppearanceSettings>({} as AppearanceSettings);

  const handleChangeGeneral = (value: GeneralSettings) => {};
  const handleChangeAppearance = (value: AppearanceSettings) => {};

  // useEffect(() => {
  //   // const preferences = LoadAppData();
  //   GetProducts()
  //     .then((products) => {
  //       console.log(products);
  //     })
  //     .catch((err) => {
  //       console.log(err);
  //     });
  //   // setGeneralSettings();
  // }, []);

  const renderContent = (tab: PreferenceTabs) => {
    switch (tab) {
      case PreferenceTabs.General:
        return <General onChange={handleChangeGeneral} value={generalSettings} />;
      case PreferenceTabs.Appearance:
        return <Appearance onChange={handleChangeAppearance} value={appearanceSettings} />;
    }
  };
  return (
    <Modal
      className={styles.modal}
      open={isOpen}
      onRequestClose={() => setIsOpen(false)}
      modalHeading="Preferences"
      primaryButtonText="Add"
      secondaryButtonText="Cancel"
    >
      <div className={styles.modal}>
        <TreeView active={selected} selected={[selected]} label={''} className={styles.treeView}>
          <TreeNode id={PreferenceTabs.General} label="General" onClick={() => setSelected(PreferenceTabs.General)} />
          <TreeNode id={PreferenceTabs.Appearance} label="Appearance" onClick={() => setSelected(PreferenceTabs.Appearance)} />
        </TreeView>
        <Separator className={styles.divider} />
        <div className={styles.content}>
          {renderContent(selected)}
          {/* <TextInput
            data-modal-primary-focus
            id="text-input-1"
            labelText="Domain name"
            placeholder="e.g. github.com"
            style={{
              marginBottom: '1rem',
            }}
          />
          <Select id="select-1" defaultValue="us-south" labelText="Region">
            <SelectItem value="us-south" text="US South" />
            <SelectItem value="us-east" text="US East" />
          </Select>
          <Dropdown
            id="drop"
            label="Dropdown"
            titleText="Dropdown"
            items={[
              {
                id: 'one',
                label: 'one',
                name: 'one',
              },
              {
                id: 'two',
                label: 'two',
                name: 'two',
              },
            ]}
          />
          <br />
          <div ref={menuTargetref}>
            <MenuButton label="Actions" menuTarget={menuTargetref.current ?? undefined} menuAlignment={'top'}>
              <MenuItem label="First action" />
              <MenuItem label="Second action" />
              <MenuItem label="Third action" />
              <MenuItem label="Danger action" kind="danger" />
            </MenuButton>
          </div>
          <br />
          <MultiSelect
            id="test"
            label="Multiselect"
            titleText="Multiselect"
            items={[
              {
                id: 'downshift-1-item-0',
                text: 'Option 1',
              },
              {
                id: 'downshift-1-item-1',
                text: 'Option 2',
              },
            ]}
            itemToString={(item) => (item ? item.text : '')}
          /> */}
        </div>
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
    justify-content: flex-start;
    align-content: flex-start;
    height: 100%;
    @media (min-width: 768px) {
      /* Règles appliquées pour les écrans d'au moins 768px */
      & .cds--modal-container {
        width: 90% !important;
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
