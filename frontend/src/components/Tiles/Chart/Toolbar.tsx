import { Dropdown, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { ChartConfig, Range, Timeframe, TimeframeOptions } from '../../../types/tiles';
import { TickerSelect } from '../../TickerSelect/TickerSelect';

interface Props {
  isLocked: boolean;
  onDelete: () => void;
  config: ChartConfig;
  onConfigChange: (config: ChartConfig) => void;
}

const getOptimalRange = (timeframe: Timeframe): Range => {
  switch (timeframe) {
    case Timeframe.oneMin:
      return Range.oneWeek;
    case Timeframe.fiveMin:
      return Range.oneWeek;
    case Timeframe.fifteenMin:
      return Range.oneMonth;
    case Timeframe.thirtyMin:
      return Range.sixMonths;
    case Timeframe.oneHour:
      return Range.sixMonths;
    case Timeframe.fourHour:
      return Range.oneYear;
    case Timeframe.oneDay:
      return Range.fiveYear;
    case Timeframe.oneWeek:
      return Range.fiveYear;
    case Timeframe.oneMonth:
      return Range.fiveYear;
    default:
      return Range.oneDay;
  }
};

export const Toolbar = ({ config, isLocked, onConfigChange, onDelete }: Props): JSX.Element => {
  const handleTickerChange = (ticker: string) => {
    onConfigChange({ ...config, ticker });
  };
  const handleTimeframeChange = (timeframe: Timeframe) => {
    onConfigChange({ ...config, timeframe, range: getOptimalRange(timeframe) });
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
