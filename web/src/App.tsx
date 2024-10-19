import '@carbon/charts-react/styles.css';
import { css } from '@emotion/css';
import { MarketMock } from './components/MarketMock/MarketMock';

export const App = () => {
  return (
    <>
      <div className={styles.container}>
        <MarketMock />
      </div>
    </>
  );
};

const styles = {
  container: css`
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 16px;
  `,
};
