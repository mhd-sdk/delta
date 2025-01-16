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
  symbol: string;
  timeframe: string;
  startDate: string;
  endDate: string;
}

export interface TimeAndSalesConfig {
  symbol: string;
}
