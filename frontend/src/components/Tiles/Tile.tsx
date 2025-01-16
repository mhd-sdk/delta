import { Menu, MenuItem, useContextMenu } from '@carbon/react';
import { useRef } from 'react';
import { ChartConfig, TileEnum, TileInterface } from '../../types/tiles';
import { AccountOverview } from './AccountOverview';
import { Chart } from './Chart/Chart';

interface Props {
  tile: TileInterface;
  onDelete: (id: string) => void;
  isLocked: boolean;
  onConfigChange?: (tile: TileInterface) => void;
}

export const Tile = ({ tile, isLocked, onDelete, onConfigChange }: Props) => {
  const el = useRef<HTMLDivElement | null>(null);
  const menuProps = useContextMenu(el);

  const handleChartChange = (config: ChartConfig) => {
    if (tile.content.type === TileEnum.Chart) {
      onConfigChange?.({ ...tile, content: { ...tile.content, config } });
    }
  };

  const renderTile = (tile: TileInterface) => {
    switch (tile.content.type) {
      case TileEnum.Chart:
        return <Chart isLocked={isLocked} config={tile.content.config} onConfigChange={handleChartChange} />;
      case TileEnum.AccountOverview:
        return <AccountOverview />;
      default:
        return null;
    }
  };
  const renderMenu = (type: TileEnum) => {
    return (
      <>
        {type === TileEnum.Chart && (
          <>
            <MenuItem
              label="Symbol Info"
              onClick={() => {
                console.log('Edit clicked');
              }}
            />
            <MenuItem
              label="Link"
              onClick={() => {
                console.log('Link clicked');
              }}
            />
          </>
        )}
        <MenuItem label="Delete" kind="danger" onClick={() => onDelete(tile.id)} />
      </>
    );
  };

  return (
    <>
      <div
        ref={el}
        style={{
          width: '100%',
          height: '100%',
        }}
      >
        {renderTile(tile)}
      </div>
      <Menu
        label="Tile menu"
        {...menuProps}
        mode="full"
        onMouseDown={(e) => e.stopPropagation()} // Prevent drag activation
      >
        {renderMenu(tile.content.type)}
      </Menu>
    </>
  );
};

Tile.displayName = 'Tile';
