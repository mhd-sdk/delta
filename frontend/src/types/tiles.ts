export enum TileEnum {
  Chart = 'Chart',
  Positions = 'Positions',
  DepthOfMarket = 'DepthOfMarket',
}

export interface TileInterface {
  x: number;
  y: number;
  w: number;
  h: number;
  content: {
    id: string;
    type: TileEnum;
  };
}
