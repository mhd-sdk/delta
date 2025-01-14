import { Search } from '@carbon/icons-react';
import { ComboBox, Dropdown, IconButton, OverflowMenu, OverflowMenuItem } from '@carbon/react';
import { css } from '@emotion/css';
import { ChartConfig } from '../../types/tiles';

interface Props {
  config: ChartConfig;
}

export const Chart = ({ config }: Props): JSX.Element => {
  const timeframes = [
    { id: '1', text: '1m' },
    { id: '5', text: '5m' },
    { id: '15', text: '15m' },
    { id: '30', text: '30m' },
    { id: '60', text: '1H' },
    { id: '240', text: '4H' },
    { id: '1440', text: '1D' },
  ];
  const symbols = ['AAPL', 'GOOGL', 'MSFT', 'AMZN', 'TSLA', 'NFLX', 'FB', 'NVDA', 'INTC', 'AMD', 'CSCO', 'QCOM'];
  return (
    <div>
      <div className={styles.header}>
        <IconButton align="bottom-left" label="Advanced search" kind="ghost" size="sm">
          <Search />
        </IconButton>
        <div className={styles.symbolWrapper}>
          <ComboBox
            initialSelectedItem={config.symbol}
            typeahead
            selectedItem={config.symbol}
            inlist={true}
            size="sm"
            onChange={() => {}}
            id="symbol-combobox"
            items={symbols}
          />
        </div>
        <Dropdown
          size="sm"
          id="inline"
          initialSelectedItem={timeframes[1]}
          label="Option 1"
          type="inline"
          items={timeframes}
          // remove parenthesis
          itemToString={(item) => (item ? item.text : '')}
          titleText={undefined}
        />
        <div className={styles.overflowMenu}>
          <OverflowMenu iconDescription="toto" direction="bottom" size="sm" flipped align="bottom">
            <OverflowMenuItem itemText="Stop app" />
            <OverflowMenuItem itemText="Restart app" />
            <OverflowMenuItem itemText="Rename app" />
            <OverflowMenuItem itemText="Clone and move app" disabled requireTitle />
            <OverflowMenuItem itemText="Edit routes and access" requireTitle />
            <OverflowMenuItem hasDivider isDelete itemText="Delete app" />
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
  overflowMenu: css`
    margin-left: auto;
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
