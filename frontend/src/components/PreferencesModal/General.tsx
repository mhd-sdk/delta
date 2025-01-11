import { RadioTile, TileGroup } from '@carbon/react';
import { css } from '@emotion/css';
import { persistence } from '../../../wailsjs/go/models';

interface Props {
  onChange: (value: persistence.GeneralPreferences) => void;
  value: persistence.GeneralPreferences;
}
1;

export const General = ({ value, onChange }: Props): JSX.Element => {
  const handleThemeChange = (theme: 'light' | 'dark') => {
    onChange({
      ...value,
      theme,
    });
  };

  return (
    <div>
      <TileGroup defaultSelected={value.theme} onChange={handleThemeChange} valueSelected={value.theme} legend="Theme" name="Theme">
        <RadioTile id="light" value="light" className={styles.mb(0.5)}>
          Light
        </RadioTile>
        <RadioTile id="dark" value="dark">
          Dark
        </RadioTile>
      </TileGroup>
    </div>
  );
};

const styles = {
  mb: (value: number) => css`
    margin-bottom: ${value}rem !important;
  `,
};
