export enum TileEnum {
  Chart = 'Chart',
  TimeAndSales = 'TimeAndSales',
  AccountOverview = 'AccountOverview',
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

export interface ChartConfig {
  ticker: string;
  timeframe: string;
  range: string;
}

export interface TimeAndSalesConfig {
  ticker: string;
}
