import { Menu, MenuItem, useContextMenu } from '@carbon/react';
import { useRef } from 'react';
import { TileEnum, TileInterface } from '../../types/tiles';
import { AccountOverview } from './AccountOverview';
import { Chart } from './Chart/Chart';

interface Props {
  tile: TileInterface;
  onDelete: (id: string) => void;
  isLocked: boolean;
}

export const Tile = ({ tile, isLocked, onDelete }: Props) => {
  const el = useRef<HTMLDivElement | null>(null);
  const menuProps = useContextMenu(el);
  const renderTile = (tile: TileInterface) => {
    switch (tile.content.type) {
      case TileEnum.Chart:
        return <Chart isLocked={isLocked} config={tile.content.config} />;
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
              label="Advanced search"
              onClick={() => {
                console.log('Edit clicked');
              }}
            />
            <MenuItem
              label="Symbol Info"
              onClick={() => {
                console.log('Edit clicked');
              }}
            />
            <MenuItem
              label="Link"
              onClick={() => {
                console.log('Edit clicked');
              }}
            />
          </>
        )}
        <MenuItem label="Delete" kind="danger" onClick={() => onDelete(tile.content.id)} />
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
