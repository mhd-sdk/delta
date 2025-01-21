import { css } from '@emotion/css';
import { Responsive, WidthProvider } from 'react-grid-layout';
import { useAppData } from '../../hooks/useAppData';
import { TileInterface } from '../../types/tiles';
import { Tile } from '../Tiles/Tile';

interface Props {
  tiles: TileInterface[];
  onChange: (tiles: TileInterface[]) => void;
}
const ResponsiveReactGridLayout = WidthProvider(Responsive);

export const Grid = ({ tiles, onChange }: Props): JSX.Element => {
  const layouts = {
    lg: tiles.map(({ x, y, h, w, id }) => ({ i: id, x, y, h, w })),
  };

  const handleLayoutChange = (l: ReactGridLayout.Layout[]) => {
    const newTiles: TileInterface[] = l.map(({ i, x, y, h, w }) => ({
      x,
      y,
      h,
      w,
      id: i,
      content: {
        ...tiles.find((tile) => tile.id === i)!.content,
      },
    }));
    onChange?.(newTiles);
  };

  const handleDelete = (id: string) => {
    const newTiles = tiles.filter((tile) => tile.id !== id);
    onChange(newTiles);
  };

  const handleConfigChange = (tile: TileInterface) => {
    const updatedTiles = tiles.map((t) => (t.id === tile.id ? tile : t));
    onChange(updatedTiles);
  };

  const { appData } = useAppData();

  return (
    <ResponsiveReactGridLayout
      onLayoutChange={handleLayoutChange}
      className={styles.layout}
      useCSSTransforms={true}
      compactType="vertical"
      breakpoints={{ lg: 0 }}
      layouts={layouts}
      cols={{ lg: 50 }}
      autoSize={true}
      rowHeight={10}
      draggableHandle=".drag-handle"
      draggableCancel=".drag-cancel"
    >
      {tiles.map((tile) => (
        <div key={tile.id} className={styles.tile(appData.preferences.generalPreferences.theme)}>
          <Tile tile={tile} onDelete={handleDelete} onConfigChange={handleConfigChange} />
        </div>
      ))}
    </ResponsiveReactGridLayout>
  );
};

const styles = {
  layout: css`
    height: calc(100vh - 3rem) !important;
  `,
  tile: (theme: string) => css`
    background-color: ${theme === 'light' ? '#fff' : '#393939'};
    border: 1px solid ${theme === 'light' ? '#ddd' : '#505050'};
    padding: 0;
  `,
};
