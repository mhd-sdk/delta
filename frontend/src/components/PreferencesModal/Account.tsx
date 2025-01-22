import { Information } from '@carbon/icons-react';
import { Button, CopyButton, Link, TextInput, TextInputSkeleton, Toggletip, ToggletipButton, ToggletipContent, ToggletipLabel } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useState } from 'react';
import { GetAccount } from '../../../wailsjs/go/app/App';
import { alpaca } from '../../../wailsjs/go/models';
import { BrowserOpenURL } from '../../../wailsjs/runtime/runtime';
interface Props {
  onLogout: () => void;
}

export const Account = ({ onLogout }: Props): JSX.Element => {
  const [account, setAccount] = useState<alpaca.Account>();
  const isLoading = account === undefined;
  useEffect(() => {
    const fetchAccount = async () => {
      setTimeout(async () => {
        const acc = await GetAccount();
        setAccount(acc);
      }, 500);
    };
    fetchAccount();
  }, []);

  return (
    <>
      <div className={styles.wrapper}>
        <div className={styles.field}>
          {isLoading ? (
            <TextInputSkeleton />
          ) : (
            <>
              <TextInput id="account-number" labelText="Account number" value={account?.account_number} />
              <CopyButton autoAlign className={styles.copyButton} onClick={() => navigator.clipboard.writeText(account?.account_number ?? '')} />
            </>
          )}
        </div>
        <div className={styles.field}>
          {isLoading ? <TextInputSkeleton /> : <TextInput id="currency" labelText="Currency" value={account?.currency} />}
        </div>
        <div className={styles.field}>{isLoading ? <TextInputSkeleton /> : <TextInput id="status" labelText="Status" value={account?.status} />}</div>

        <div className={styles.field}>{isLoading ? <TextInputSkeleton /> : <TextInput id="equity" labelText="Equity" value={account?.equity} />}</div>

        <div className={styles.field}>{isLoading ? <TextInputSkeleton /> : <TextInput id="balance" labelText="Balance" value={account?.cash} />}</div>

        <div className={styles.field}>
          {isLoading ? (
            <TextInputSkeleton />
          ) : (
            <TextInput
              id="multiplier"
              labelText={
                <>
                  <ToggletipLabel>Multiplier</ToggletipLabel>
                  <Toggletip align={'right-top'}>
                    <ToggletipButton label="Show information">
                      <Information />
                    </ToggletipButton>
                    <ToggletipContent>
                      <p>
                        Correspond to the allowed{' '}
                        <Link onClick={() => BrowserOpenURL('https://en.wikipedia.org/wiki/Leverage_(finance)')} href="#">
                          leverage
                        </Link>{' '}
                        for the account.
                      </p>
                    </ToggletipContent>
                  </Toggletip>
                </>
              }
              value={account?.multiplier}
            />
          )}
        </div>

        <div className={styles.field}>
          {isLoading ? <TextInputSkeleton /> : <TextInput id="buyingpower" labelText="Buying power" value={account?.buying_power} />}
        </div>
        <div className={styles.logoutButton}>
          <Button kind="danger--ghost" onClick={onLogout}>
            Logout
          </Button>
        </div>
      </div>
    </>
  );
};

const styles = {
  logoutButton: css`
    margin-left: auto;
    margin-top: 1rem;
  `,
  fieldLabel: css`
    display: flex;
    gap: 0.2rem;
  `,
  copyButton: css`
    top: 0.75rem;
  `,
  wrapper: css`
    display: flex;
    flex-direction: column;
    gap: 1rem;
  `,
  mb: (value: number) => css`
    margin-bottom: ${value}rem !important;
  `,
  field: css`
    display: flex;
    justify-content: space-between;
    align-items: center;
  `,
};
