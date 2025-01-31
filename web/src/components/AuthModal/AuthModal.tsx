import { ArrowUpRight } from '@carbon/icons-react';
import { ComposedModal, Link, ModalBody, ModalFooter, ModalHeader, TextInput, Toggle } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { SaveAppData, TestCredentials } from '../../../wailsjs/go/app/App';
import { models } from '../../../wailsjs/go/models';
import { BrowserOpenURL } from '../../../wailsjs/runtime/runtime';
import { useAppData } from '../../hooks/useAppData';

interface Props {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

export const AuthModal = ({ isOpen, setIsOpen }: Props): JSX.Element => {
  const [useKeys, setUseKeys] = useState(false);
  const [key, setKey] = useState('');
  const [secret, setSecret] = useState('');
  const [loadingState, setLoadingState] = useState<'inactive' | 'active' | 'finished' | 'error'>('inactive');

  const { appData } = useAppData();

  const handleToggle = () => {
    setUseKeys(!useKeys);
  };

  useEffect(() => {
    const loadKeys = async () => {
      if (appData.keys.apiKey === '' || appData.keys.secretKey === '') {
        setIsOpen(true);
      }
    };
    loadKeys();
  }, [isOpen]);

  const handleLogin = async () => {
    SaveAppData({
      ...appData,
      keys: {
        apiKey: useKeys ? key : appData.keys.apiKey,
        secretKey: useKeys ? secret : appData.keys.secretKey,
      },
    } as models.AppData);
    setIsOpen(false);
    setLoadingState('inactive');
    setSecret('');
    setKey('');
    setUseKeys(false);
    setIsOpen(false);
  };

  const handleRequest = async () => {
    setLoadingState('active');
    const res = await TestCredentials(key, secret);
    if (res === true) {
      handleLogin();
      setLoadingState('finished');
    } else {
      setLoadingState('error');
    }
  };

  const description = loadingState === 'error' ? 'Invalid credentials' : 'Logging in...';

  return (
    <ComposedModal preventCloseOnClickOutside className={styles.modal} open={isOpen} onClose={() => false}>
      <ModalHeader title="Broker Authentication" />
      <ModalBody>
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
            <TextInput
              id={'API-KEY'}
              labelText="Key"
              value={key}
              onChange={(e) => {
                setLoadingState('inactive');
                setKey(e.target.value);
              }}
            />
            <TextInput
              id={'API-SECRET'}
              labelText="Secret"
              value={secret}
              onChange={(e) => {
                setLoadingState('inactive');
                setSecret(e.target.value);
              }}
            />
          </div>
        )}
      </ModalBody>
      {useKeys && (
        // @ts-expect-error: idk why but typescript is not happy with this
        <ModalFooter primaryButtonText="Login" loadingStatus={loadingState} loadingDescription={description} onRequestSubmit={handleRequest} />
      )}
    </ComposedModal>
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
