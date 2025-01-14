import { Menu, MenuItem, useContextMenu } from '@carbon/react';
import { css } from '@emotion/css';
import { useRef } from 'react';
import { TileEnum, TileInterface } from '../../types/tiles';
import { AccountOverview } from './AccountOverview';
import { Chart } from './Chart';

interface Props {
  tile: TileInterface;
  onDelete: (id: string) => void;
}

export const Tile = ({ tile, onDelete }: Props) => {
  const el = useRef(null);
  const menuProps = useContextMenu(el);
  const renderTile = (tile: TileInterface) => {
    switch (tile.content.type) {
      case TileEnum.Chart:
        return <Chart config={tile.content.config} />;
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
              onClick={(e) => {
                console.log('Edit clicked');
              }}
            />
            <MenuItem
              label="Configure"
              onClick={(e) => {
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
          cursor: 'context-menu',
          width: '100%',
          height: '100%',
        }}
      >
        {renderTile(tile)}
      </div>
      <Menu
        label={'dfdf'}
        {...menuProps}
        mode="basic"
        onMouseDown={(e) => e.stopPropagation()} // Prevent drag activation
      >
        {renderMenu(tile.content.type)}
      </Menu>
    </>
  );
};

Tile.displayName = 'Tile';

const styles = {
  tile: css`
    background-color: #fff;
    border: 1px solid #ddd;
    padding: 0;
  `,
};
