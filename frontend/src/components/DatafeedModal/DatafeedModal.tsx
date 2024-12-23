import { Checkbox, Dropdown, DropdownSkeleton, Modal, PasswordInput, Stack, TextInput } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { GetSystems, LoadAppData } from '../../../wailsjs/go/main/App';

interface Props {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

export const DatafeedModal = ({ isOpen, setIsOpen }: Props): JSX.Element => {
  const [isSystemsLoading, setIsSystemsLoading] = useState(true);
  const [isSystemsError, setIsSystemsError] = useState(false);
  const [systems, setSystems] = useState<string[]>([]);

  const [userId, setUserId] = useState('');
  const [password, setPassword] = useState('');
  const [rememberMe, setRememberMe] = useState(false);
  const [selectedSystem, setSelectedSystem] = useState('');

  useEffect(() => {
    LoadAppData().then((data) => {});
    setIsSystemsLoading(true);
    GetSystems().then((systems) => {
      if (!systems.system_name) {
        return setIsSystemsError(true);
      }
      setSystems(systems.system_name);
      setIsSystemsLoading(false);
    });
  }, []);

  return (
    <Modal preventCloseOnClickOutside className={styles.modal} open={isOpen} modalHeading="Datafeed authentification" primaryButtonText="Connect">
      <div>
        <Stack gap={5} className={styles.fullWidth}>
          <Stack gap={3}>
            <TextInput id="user-id" type="text" labelText="User ID" value={userId} onChange={(e) => setUserId(e.target.value)} />
            <PasswordInput id="password" labelText="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
          </Stack>
          <div className={styles.stack}>
            <Checkbox labelText="Remember me" id="remember-me" value="remember-me" checked={rememberMe} onChange={() => setRememberMe(!rememberMe)} />
            {isSystemsLoading ? (
              <DropdownSkeleton />
            ) : (
              <Dropdown
                className={styles.system}
                autoAlign
                id="system-dropdown"
                invalid={isSystemsError}
                invalidText="Could not load available systems"
                titleText="System"
                initialSelectedItem={systems[0]}
                label="Rithmic System name"
                items={[...systems, 'oko', 'oko', 'oko', 'oko', 'oko', 'oko', 'oko']}
                onChange={(e) => {
                  if (e.selectedItem) {
                    setSelectedSystem(e.selectedItem);
                  }
                }}
              />
            )}
          </div>
        </Stack>
      </div>
    </Modal>
  );
};

const styles = {
  fullWidth: css`
    width: 100%;
  `,
  modal: css`
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-content: flex-start;
    overflow: hidden;
    @media (min-width: 768px) {
      /* Règles appliquées pour les écrans d'au moins 768px */
      & .cds--modal-container {
        width: 600px !important;
      }
    }
  `,
  stack: css`
    display: flex;
    align-items: end;
    overflow: none;
  `,
  system: css`
    width: 50% !important;
  `,
};
