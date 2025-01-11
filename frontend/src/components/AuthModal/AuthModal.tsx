import { ArrowUpRight } from '@carbon/icons-react';
import { Checkbox, Link, Modal } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { GetAppData } from '../../../wailsjs/go/main/App';
import { persistence } from '../../../wailsjs/go/models';
import { BrowserOpenURL } from '../../../wailsjs/runtime/runtime';

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
      passiveModal
      modalHeading="Broker Authentication"
      onRequestClose={() => setIsOpen(false)}
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
      <Checkbox labelText="Keep logged in" id="remember-me" value="remember-me" checked={rememberMe} onChange={handleChangeRememberMe} />
    </Modal>
  );
};

const styles = {
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
