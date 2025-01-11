import { ArrowUpRight } from '@carbon/icons-react';
import { Link, Modal, TextInput, Toggle } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { GetAppData, SaveAppData } from '../../../wailsjs/go/main/App';
import { persistence } from '../../../wailsjs/go/models';
import { BrowserOpenURL } from '../../../wailsjs/runtime/runtime';

interface Props {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

export const AuthModal = ({ isOpen, setIsOpen }: Props): JSX.Element => {
  const [useKeys, setUseKeys] = useState(false);
  const [key, setKey] = useState('');
  const [secret, setSecret] = useState('');
  const handleToggle = () => {
    setUseKeys(!useKeys);
  };

  useEffect(() => {
    const loadKeys = async () => {
      const data = await GetAppData();
      if (data.keys.apiKey === '' || data.keys.secretKey === '') {
        setIsOpen(true);
      }
    };
    loadKeys();
  }, []);

  const handleSubmit = async () => {
    const data = await GetAppData();
    SaveAppData({
      ...data,
      keys: {
        apiKey: useKeys ? key : data.keys.apiKey,
        secretKey: useKeys ? secret : data.keys.secretKey,
      },
    } as persistence.AppData);
    setIsOpen(false);
  };

  return (
    <Modal
      preventCloseOnClickOutside
      className={styles.modal}
      open={isOpen}
      passiveModal={!useKeys}
      modalHeading="Broker Authentication"
      primaryButtonText="Connect"
      closeButtonLabel="You cannot close this modal"
      onRequestSubmit={handleSubmit}
    >
      <p
        style={{
          marginBottom: '1rem',
        }}
      >
        Delta software use{' '}
        <Link inline href="#" onClick={() => BrowserOpenURL('https://alpaca.markets/about-us')}>
          Alpaca
        </Link>{' '}
        services for market data and trading. To use the application, you need to{' '}
        <Link
          data-modal-primary-focus
          href="#"
          onClick={() =>
            BrowserOpenURL(
              'https://app.alpaca.markets/oauth/authorize?response_type=code&client_id=6f5ab5debd23bc6da75c5f87a5fe3f58&redirect_uri=delta:&scope=account:write&scope=data&scope=trading'
            )
          }
          renderIcon={() => <ArrowUpRight />}
        >
          connect your Alpaca account
        </Link>
        .
      </p>
      <Toggle
        size="sm"
        labelText=""
        labelA="Use API Keys instead"
        labelB="Use API Keys instead"
        id="togglekeys"
        toggled={useKeys}
        onToggle={handleToggle}
      />

      {useKeys && (
        <div className={styles.flex}>
          <TextInput id={'API-KEY'} labelText="Key" value={key} onChange={(e) => setKey(e.target.value)} />
          <TextInput id={'API-SECRET'} labelText="Secret" value={secret} onChange={(e) => setSecret(e.target.value)} />
        </div>
      )}
    </Modal>
  );
};

const styles = {
  flex: css`
    display: flex;
    flex-direction: row;
    gap: 1rem;
  `,
  modal: css`
    display: flex;
    flex-direction: row;
    align-content: flex-start;
    overflow: hidden;
    @media (min-width: 768px) {
      & .cds--modal-container {
        width: 600px !important;
      }
    }
  `,
};
