import { css } from '@emotion/css';
import { GeneralSettings } from '../../types/preferences';

interface Props {
  onChange: (value: GeneralSettings) => void;
  value: GeneralSettings;
}
1;

export const General = ({ value, onChange }: Props): JSX.Element => {
  return <div></div>;
};

const styles = {
  mb: (value: number) => css`
    margin-bottom: ${value}rem !important;
  `,
};
