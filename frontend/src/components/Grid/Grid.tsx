import { css } from '@emotion/css';
import { Responsive, WidthProvider } from 'react-grid-layout';
import { TileInterface } from '../../types/tiles';
import { ChartTile } from '../Tiles/ChartTile';
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

  const renderTile = (tile: TileInterface) => {
    switch (tile.content.type) {
      case 'Chart':
        return <ChartTile />;
      default:
        return null;
    }
  };

  return (
    <ResponsiveReactGridLayout
      useCSSTransforms={true}
      isDraggable={!isLocked}
      isResizable={!isLocked}
      compactType="vertical"
      className={styles.layout}
      layouts={layouts}
      breakpoints={{ lg: 0 }}
      cols={{ lg: 50 }}
      autoSize={true}
      rowHeight={10}
      onLayoutChange={handleLayoutChange}
    >
      {tiles.map((tile) => (
        <div key={tile.content.id} className={styles.tile}>
          <Tile tile={tile} />
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
