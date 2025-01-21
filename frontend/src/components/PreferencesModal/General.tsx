import { Button, RadioTile, TileGroup } from '@carbon/react';
import { css } from '@emotion/css';
import { ResetPreferences } from '../../../wailsjs/go/app/App';
import { persistence } from '../../../wailsjs/go/models';
import { useAppData } from '../../hooks/useAppData';

interface Props {
  onChange: (value: persistence.GeneralPreferences) => void;
  value: persistence.GeneralPreferences;
}

export const General = ({ value, onChange }: Props): JSX.Element => {
  const handleThemeChange = (theme: 'light' | 'dark') => {
    onChange({
      ...value,
      theme,
    });
  };

  const { refetch } = useAppData();
  console.log({ value });

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
      <Button
        onClick={async () => {
          await ResetPreferences();
          await refetch();
        }}
      >
        Reset preferences
      </Button>
    </div>
  );
};

const styles = {
  mb: (value: number) => css`
    margin-bottom: ${value}rem !important;
  `,
};
