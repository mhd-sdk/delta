import { Dropdown, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { ChartConfig } from '../../../types/tiles';
import { TickerSelect } from '../../TickerSelect/TickerSelect';

interface Props {
  isLocked: boolean;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

export const Chart = ({ config, isLocked, onConfigChange }: Props): JSX.Element => {
  const timeframes = ['1m', '5m', '15m', '30m', '1h', '4h', '1d'];
  const ranges = ['1 day', '3 days', '1 week', '1 month', '3 months', '6 months', '1 year', '5 year'];

  const handleTickerChange = (ticker: string) => {
    onConfigChange({ ...config, ticker });
  };

  const handleTimeframeChange = (timeframe: string) => {
    onConfigChange({ ...config, timeframe });
  };

  const handleRangeChange = (range: string) => {
    onConfigChange({ ...config, range });
  };

  return (
    <div className={styles.height100}>
      <div className={styles.header}>
        <TickerSelect disabled={!isLocked} value={config.ticker} onChange={handleTickerChange} />
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
        <Dropdown
          disabled={!isLocked}
          className={styles.timeframeContainer}
          size="sm"
          id="inline"
          initialSelectedItem={config.range}
          label="range"
          // type="inline"
          items={ranges}
          hideLabel
          titleText={undefined}
          onChange={({ selectedItem }) => selectedItem && handleRangeChange(selectedItem)}
        />
        <div className={styles.overflowMenu(isLocked)}>
          <OverflowMenu disabled={!isLocked} iconDescription="toto" direction="bottom" size="sm" flipped align="bottom">
            <OverflowMenuItem itemText="Ticker info" />
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
  tickerContainer: css`
    width: 120px;
  `,
  timeframeContainer: css`
    width: 150px;
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

    .cds--list-box {
      border: 0px !important;
      :hover {
      }
    }
  `,
};
