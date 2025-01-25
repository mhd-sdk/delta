import { models } from '../../../wailsjs/go/models';
import { ChartConfig, TileType } from '../../types/tiles';
import { AccountInfo } from './AccountInfo/AccountInfo';
import { Chart } from './Chart/Chart';

interface Props {
  tile: models.Tile;
  onDelete: (id: string) => void;
  onConfigChange: (tile: models.Tile) => void;
}

export const Tile = ({ tile, onDelete, onConfigChange }: Props) => {
  const handleChartChange = (config: ChartConfig) => {
    if (tile.data.type === TileType.Chart) {
      onConfigChange({ ...tile, data: { type: TileType.Chart, config } } as models.Tile);
    }
  };

  const handleDelete = () => {
    onDelete(tile.id);
  };

  const renderTile = (tile: models.Tile) => {
    switch (tile.data.type) {
      case TileType.Chart:
        return <Chart onDelete={handleDelete} config={tile.data.config} onConfigChange={handleChartChange} />;
      case TileType.AccountInfo:
        return <AccountInfo />;
      default:
        return null;
    }
  };

  return renderTile(tile);
};

Tile.displayName = 'Tile';
