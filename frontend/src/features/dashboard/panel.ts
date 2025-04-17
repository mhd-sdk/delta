import { Range, Timeframe, Unit } from './timerange';

export const defaultTimeframes: Timeframe[] = [
  { n: 1, unit: Unit.min },
  { n: 5, unit: Unit.min },
  { n: 15, unit: Unit.min },
  { n: 30, unit: Unit.min },
  { n: 1, unit: Unit.hour },
  { n: 4, unit: Unit.hour },
  { n: 1, unit: Unit.day },
  { n: 1, unit: Unit.week },
  { n: 1, unit: Unit.month },
];

export interface Panel {
  height: number;
  width: number;
  x: number;
  y: number;
  id: string;
  data: PanelData;
}

export enum PanelType {
  Chart = 'Chart',
  TimeAndSales = 'TimeAndSales',
  AccountInfo = 'AccountInfo',
  TickerInfo = 'TickerInfo',
  Headlines = 'Headlines',
  Scanner = 'Scanner',
}

export type PanelData =
  | {
      type: PanelType.Chart;
      config: ChartConfig;
    }
  | {
      type: PanelType.TimeAndSales;
      config: TimeAndSalesConfig;
    }
  | {
      type: PanelType.AccountInfo;
    };

export type Configs = ChartConfig | TimeAndSalesConfig;

export interface ChartConfig {
  ticker: string;
  timeframe: Timeframe;
  range: Range;
}

export interface TimeAndSalesConfig {
  ticker: string;
}
