export const ChartToolbar = ({ config, onConfigChange, onDelete }: Props): JSX.Element => {
  const handleTickerChange = (ticker: string) => {
    onConfigChange({ ...config, ticker });
  };
  const handleTimeframeChange = (timeframe: Timeframe) => {
    onConfigChange({ ...config, timeframe, range: calcOptimizedRange(timeframe) });
  };

  const { appData } = useAppData();
  const isDarkMode = appData.preferences.generalPreferences.theme === 'dark';

  return (
    <div className={cx('drag-handle', styles.header(isDarkMode))}>
      <TickerSelect value={config.ticker} onChange={handleTickerChange} />
      <Dropdown
        className={cx('drag-cancel', styles.timeframeContainer)}
        size="sm"
        id="inline"
        initialSelectedItem={config.timeframe}
        label="Timeframe"
        items={defaultTimeframes}
        titleText=""
        itemToString={(item: Timeframe) => `${item.n} ${item.unit}`}
        onChange={({ selectedItem }) => selectedItem && handleTimeframeChange(selectedItem)}
      />

      <div className={cx('drag-cancel', styles.overflowMenu)}>
        <OverflowMenu iconDescription="toto" direction="bottom" size="sm" flipped align="bottom">
          <OverflowMenuItem itemText="Ticker info" />
          <OverflowMenuItem itemText="Link" />
          <OverflowMenuItem hasDivider isDelete itemText="Delete" onClick={onDelete} />
        </OverflowMenu>
      </div>
    </div>
  );
};
