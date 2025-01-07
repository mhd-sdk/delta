import { Checkbox, Link, Modal } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { GetAppData } from '../../../wailsjs/go/main/App';
import { persistence } from '../../../wailsjs/go/models';

interface Props {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

export const AuthModal = ({ isOpen, setIsOpen }: Props): JSX.Element => {
  const [appData, setAppData] = useState<persistence.AppData>();
  const [rememberMe, setRememberMe] = useState(false);

  useEffect(() => {
    GetAppData().then((data) => {
      setRememberMe(data.user.rememberMe);
      setAppData(data);
    });
  }, []);
  console.log({ appData });

  const handleChangeRememberMe = () => {
    setRememberMe(!rememberMe);
    setAppData({ ...appData } as persistence.AppData);
  };

  return (
    <Modal
      preventCloseOnClickOutside
      className={styles.modal}
      open={isOpen}
      modalHeading="Broker Authentication"
      primaryButtonText="Connect"
      onRequestClose={() => setIsOpen(false)}
    >
      <p
        style={{
          marginBottom: '1rem',
        }}
      >
        Delta software use{' '}
        <Link inline href="#">
          Alpaca
        </Link>{' '}
        services for market data and trading. To use the application, you need to connect your Alpaca account.
      </p>
      <Checkbox labelText="Remember me" id="remember-me" value="remember-me" checked={rememberMe} onChange={() => setRememberMe(!rememberMe)} />
    </Modal>
  );
};

const styles = {
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
};
