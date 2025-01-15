import { Dropdown, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { ChartConfig } from '../../../types/tiles';
import { SymbolSelect } from '../../SymbolSelect/SymbolSelect';

interface Props {
  isLocked: boolean;
  config: ChartConfig;
}

export const Chart = ({ config, isLocked }: Props): JSX.Element => {
  const timeframes = ['1m', '5m', '15m', '30m', '1h', '4h', '1d'];
  const symbols = ['AAPL', 'GOOGL', 'MSFT', 'AMZN', 'TSLA', 'NFLX', 'FB', 'NVDA', 'INTC', 'AMD', 'CSCO', 'QCOM'];
  return (
    <div>
      <div className={styles.header}>
        {/* <IconButton align="bottom-left" label="Advanced search" kind="ghost" size="sm">
          <Search />
        </IconButton> */}
        {/* <div className={styles.symbolWrapper}> */}
        <SymbolSelect value={config.symbol} onChange={() => {}} />
        {/* </div> */}
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
        />
        <div className={styles.overflowMenu(isLocked)}>
          <OverflowMenu disabled={!isLocked} iconDescription="toto" direction="bottom" size="sm" flipped align="bottom">
            <OverflowMenuItem itemText="Advanced search" />
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
