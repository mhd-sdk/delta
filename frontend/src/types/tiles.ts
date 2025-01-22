import { RangeEnum } from './range';
import { Timeframe } from './timeframe';

export enum TileType {
  Chart = 'Chart',
  TimeAndSales = 'TimeAndSales',
  AccountInfo = 'AccountInfo',
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
  data: TileData;
}

export type TileData =
  | {
      type: TileType.Chart;
      config: ChartConfig;
    }
  | {
      type: TileType.TimeAndSales;
      config: TimeAndSalesConfig;
    }
  | {
      type: TileType.AccountInfo;
    };

export type Configs = ChartConfig | TimeAndSalesConfig;

export interface ChartConfig {
  ticker: string;
  timeframe: Timeframe;
  range: RangeEnum;
}

export interface TimeAndSalesConfig {
  ticker: string;
}
