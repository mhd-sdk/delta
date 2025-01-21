export enum TileEnum {
  Chart = 'Chart',
  TimeAndSales = 'TimeAndSales',
  AccountOverview = 'AccountOverview',
  TickerInfo = 'TickerInfo',
  Headlines = 'Headlines',
  Scanner = 'Scanner',
}

export interface TileInterface {
  id: string;
  x: number;
  y: number;
  w: number;
  h: number;
  content: TileContent;
}

export type TileContent =
  | {
      type: TileEnum.Chart;
      config: ChartConfig;
    }
  | {
      type: TileEnum.TimeAndSales;
      config: TimeAndSalesConfig;
    }
  | {
      type: TileEnum.AccountOverview;
    };

export type Configs = ChartConfig | TimeAndSalesConfig;

export enum Range {
  oneDay = '1day',
  threeDays = '3days',
  oneWeek = '1week',
  oneMonth = '1month',
  threeMonths = '3months',
  sixMonths = '6months',
  oneYear = '1year',
  fiveYear = '5year',
}
export const RangeOptions: Range[] = [
  Range.oneDay,
  Range.threeDays,
  Range.oneWeek,
  Range.oneMonth,
  Range.threeMonths,
  Range.sixMonths,
  Range.oneYear,
  Range.fiveYear,
];

export enum Timeframe {
  oneMin = '1min',
  fiveMin = '5min',
  fifteenMin = '15min',
  thirtyMin = '30min',
  oneHour = '1hour',
  fourHour = '4hour',
  oneDay = '1day',
  oneWeek = '1week',
  oneMonth = '1month',
}

export const TimeframeOptions: Timeframe[] = [
  Timeframe.oneMin,
  Timeframe.fiveMin,
  Timeframe.fifteenMin,
  Timeframe.thirtyMin,
  Timeframe.oneHour,
  Timeframe.fourHour,
  Timeframe.oneDay,
  Timeframe.oneWeek,
  Timeframe.oneMonth,
];

export interface ChartConfig {
  ticker: string;
  timeframe: Timeframe;
  range: Range;
}

export interface TimeAndSalesConfig {
  ticker: string;
}
