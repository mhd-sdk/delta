import { css } from '@emotion/css';

interface Props {}
1;

export const Account = (): JSX.Element => {
  return <div>Account</div>;
};

const styles = {
  mb: (value: number) => css`
    margin-bottom: ${value}rem !important;
  `,
};
