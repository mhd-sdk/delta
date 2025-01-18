import { Dropdown, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { ChartConfig, Range, RangeOptions, Timeframe, TimeframeOptions } from '../../../types/tiles';
import { TickerSelect } from '../../TickerSelect/TickerSelect';

interface Props {
  isLocked: boolean;
  onDelete: () => void;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

export const Toolbar = ({ config, isLocked, onConfigChange, onDelete }: Props): JSX.Element => {
  const handleTickerChange = (ticker: string) => {
    onConfigChange({ ...config, ticker });
  };
  const handleTimeframeChange = (timeframe: Timeframe) => {
    onConfigChange({ ...config, timeframe });
  };
  const handleRangeChange = (range: Range) => {
    onConfigChange({ ...config, range });
  };
  return (
    <div className={styles.header}>
      <TickerSelect disabled={!isLocked} value={config.ticker} onChange={handleTickerChange} />
      <Dropdown
        disabled={!isLocked}
        className={styles.timeframeContainer}
        size="sm"
        id="inline"
        initialSelectedItem={config.timeframe}
        label="Timeframe"
        items={TimeframeOptions}
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
        items={RangeOptions}
        hideLabel
        titleText={undefined}
        onChange={({ selectedItem }) => selectedItem && handleRangeChange(selectedItem)}
        defaultValue={config.range}
      />
      <div className={styles.overflowMenu(isLocked)}>
        <OverflowMenu disabled={!isLocked} iconDescription="toto" direction="bottom" size="sm" flipped align="bottom">
          <OverflowMenuItem itemText="Ticker info" />
          <OverflowMenuItem itemText="Link" />
          <OverflowMenuItem hasDivider isDelete itemText="Delete" onClick={onDelete} />
        </OverflowMenu>
      </div>
    </div>
  );
};

const styles = {
  height100: css`
    height: 100%;
  `,
  tickerContainer: css`
    width: 150px;
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
