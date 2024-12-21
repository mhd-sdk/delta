import { Menu, MenuItem, useContextMenu } from '@carbon/react';
import { css } from '@emotion/css';
import { useRef } from 'react';
import { TileInterface } from '../../types/tiles';
import { ChartTile } from './ChartTile';

interface Props {
  tile: TileInterface;
  onDelete: (id: string) => void;
}

export const Tile = ({ tile, onDelete }: Props) => {
  const el = useRef(null);
  const menuProps = useContextMenu(el);
  const renderTile = (tile: TileInterface) => {
    switch (tile.content.type) {
      case 'Chart':
        return <ChartTile />;
      default:
        return null;
    }
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
        <MenuItem
          label="Edit"
          onClick={(e) => {
            console.log('Edit clicked');
          }}
        />
        <MenuItem label="Delete" kind="danger" onClick={() => onDelete(tile.content.id)} />
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
