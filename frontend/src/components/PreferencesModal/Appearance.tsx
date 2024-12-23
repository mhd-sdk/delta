import { RadioTile, TileGroup } from '@carbon/react';
import { css } from '@emotion/css';
import { AppearanceSettings } from '../../types/preferences';

interface Props {
  onChange: (value: AppearanceSettings) => void;
  value: AppearanceSettings;
}
1;

export const Appearance = ({ value, onChange }: Props): JSX.Element => {
  return (
    <div>
      <TileGroup valueSelected={value.theme} legend="Radio Tile Group" name="radio tile group">
        <RadioTile id="Light" value="Light" className={styles.mb(0.5)}>
          Light
        </RadioTile>
        <RadioTile id="Dark" value="Dark">
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
