import { css, cx } from '@emotion/css';

interface Props {
  className?: string;
}

export const Separator = ({ className }: Props): JSX.Element => {
  return <div className={cx(styles.separator, className)}></div>;
};

const styles = {
  separator: css`
    border-left: 1px solid #d0d0d0;
    height: 80%;
    margin: 0 0.5rem;
    opacity: 0.5;
  `,
};
