import { css } from '@emotion/css';
import { Responsive, WidthProvider } from 'react-grid-layout';
import { TileInterface } from '../../types/tiles';
import { Tile } from '../Tiles/Tile';

interface Props {
  tiles: TileInterface[];
  isLocked: boolean;
  onChange?: (tiles: TileInterface[]) => void;
}
const ResponsiveReactGridLayout = WidthProvider(Responsive);

export const Grid = ({ tiles, isLocked, onChange }: Props): JSX.Element => {
  const layouts = {
    lg: tiles.map(({ content, x, y, h, w }) => ({ i: content.id, x, y, h, w })),
  };

  const handleLayoutChange = (l: ReactGridLayout.Layout[]) => {
    const newTiles: TileInterface[] = l.map(({ i, x, y, h, w }) => ({
      x,
      y,
      h,
      w,
      content: {
        ...tiles.find((tile) => tile.content.id === i)!.content,
      },
    }));
    onChange?.(newTiles);
  };

  const handleDelete = (id: string) => {
    const newTiles = tiles.filter((tile) => tile.content.id !== id);
    onChange?.(newTiles);
  };

  return (
    <ResponsiveReactGridLayout
      onLayoutChange={handleLayoutChange}
      className={styles.layout}
      useCSSTransforms={true}
      isDraggable={!isLocked}
      isResizable={!isLocked}
      compactType="vertical"
      breakpoints={{ lg: 0 }}
      layouts={layouts}
      cols={{ lg: 50 }}
      autoSize={true}
      rowHeight={10}
    >
      {tiles.map((tile) => (
        <div key={tile.content.id} className={styles.tile}>
          <Tile tile={tile} onDelete={handleDelete} />
        </div>
      ))}
    </ResponsiveReactGridLayout>
  );
};

const styles = {
  layout: css`
    height: calc(100vh - 3rem);
  `,
  tile: css`
    background-color: #fff;
    border: 1px solid #ddd;
    padding: 0;
  `,
};
