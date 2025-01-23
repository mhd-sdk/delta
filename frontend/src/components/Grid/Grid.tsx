import { css } from '@emotion/css';
import { Responsive, WidthProvider } from 'react-grid-layout';
import { models } from '../../../wailsjs/go/models';
import { useAppData } from '../../hooks/useAppData';
import { Tile } from '../Tiles/Tile';

interface Props {
  layout: models.Tile[];
  onChange: (tiles: models.Tile[]) => void;
}
const ResponsiveReactGridLayout = WidthProvider(Responsive);

export const Grid = ({ layout, onChange }: Props): JSX.Element => {
  const layouts = {
    lg: layout.map(({ x, y, h, w, id }) => ({ i: id, x, y, h, w })),
  };

  const handleLayoutChange = (l: ReactGridLayout.Layout[]) => {
    const newTiles: models.Tile[] = l.map(({ i, x, y, h, w }) => ({
      x,
      y,
      h,
      w,
      id: i,
      data: {
        ...layout.find((tile) => tile.id === i)!.data,
      },
    })) as models.Tile[];
    onChange?.(newTiles);
  };

  const handleDelete = (id: string) => {
    const newTiles = layout.filter((tile) => tile.id !== id);
    onChange(newTiles);
  };

  const handleConfigChange = (tile: models.Tile) => {
    const updatedTiles = layout.map((t) => (t.id === tile.id ? tile : t));
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
      {layout.map((tile) => (
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
