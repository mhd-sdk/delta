export enum TileEnum {
  Chart = 'Chart',
  TimeAndSales = 'TimeAndSales',
  AccountOverview = 'AccountOverview',
}

export interface TileInterface {
  x: number;
  y: number;
  w: number;
  h: number;
  content: TileContent;
}

export type TileContent =
  | {
      id: string;
      type: TileEnum.Chart;
      config: ChartConfig;
    }
  | {
      id: string;
      type: TileEnum.TimeAndSales;
      settings: TimeAndSalesConfig;
    }
  | {
      id: string;
      type: TileEnum.AccountOverview;
    };

export interface ChartConfig {
  symbol: string;
  timeframe: string;
  startDate: string;
  endDate: string;
}

export interface TimeAndSalesConfig {
  symbol: string;
}
