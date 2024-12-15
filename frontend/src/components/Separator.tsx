import { css } from '@emotion/css';

export const Separator = (): JSX.Element => {
  return <div className={styles.separator}></div>;
};

const styles = {
  separator: css`
    border-left: 1px solid #d0d0d0;
    height: 80%;
    margin: 0 0.5rem;
    opacity: 0.5;
  `,
};
