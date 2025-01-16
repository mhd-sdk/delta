import { Dropdown, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { useState } from 'react';
import { ChartConfig } from '../../../types/tiles';
import { SymbolSelect } from '../../SymbolSelect/SymbolSelect';

interface Props {
  isLocked: boolean;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

export const Chart = ({ config, isLocked, onConfigChange }: Props): JSX.Element => {
  const timeframes = ['1m', '5m', '15m', '30m', '1h', '4h', '1d'];
  const [isSymbolSearchOpen, setIsSymbolSearchOpen] = useState(false);
  const handleSymbolChange = (symbol: string) => {
    onConfigChange({ ...config, symbol });
  };

  const handleTimeframeChange = (timeframe: string) => {
    onConfigChange({ ...config, timeframe });
  };

  return (
    <div className={styles.height100}>
      <div className={styles.header}>
        <SymbolSelect disabled={!isLocked} value={config.symbol} onChange={handleSymbolChange} />
        <Dropdown
          disabled={!isLocked}
          className={css`
            width: 100px;
          `}
          size="sm"
          id="inline"
          initialSelectedItem={config.timeframe}
          label="Timeframe"
          // type="inline"
          items={timeframes}
          titleText={undefined}
          onChange={({ selectedItem }) => selectedItem && handleTimeframeChange(selectedItem)}
        />
        <div className={styles.overflowMenu(isLocked)}>
          <OverflowMenu disabled={!isLocked} iconDescription="toto" direction="bottom" size="sm" flipped align="bottom">
            <OverflowMenuItem itemText="Symbol info" />
            <OverflowMenuItem itemText="Link" />
            <OverflowMenuItem hasDivider isDelete itemText="Delete" />
          </OverflowMenu>
        </div>
      </div>
    </div>
  );
};

const styles = {
  height100: css`
    height: 100%;
  `,
  symbolWrapper: css`
    width: 120px;
  `,
  overflowMenu: (isLocked: boolean) => css`
    margin-left: auto;
    /* if locked, opacity to 50% */
    opacity: ${isLocked ? 1 : 0.3};
  `,
  header: css`
    display: flex;
    width: 100%;
    #symbol-combobox + .cds--list-box__selection {
      display: none !important;
    }
    #symbol-combobox {
      padding-inline-end: 0rem !important;
      border: 0px !important;
    }
    .cds--list-box {
      border: 0px !important;
      :hover {
      }
    }
  `,
};
