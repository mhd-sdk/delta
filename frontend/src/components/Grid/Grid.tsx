import React from 'react';
import { css } from '@emotion/css';
import { Responsive, WidthProvider } from 'react-grid-layout';

interface Props {
  tiles: { id: string; content: JSX.Element; x: number; y: number; w: number; h: number }[];
}
const ResponsiveReactGridLayout = WidthProvider(Responsive);

export const Grid = ({ tiles }: Props): JSX.Element => {
  const layouts = {
    lg: tiles.map(({ id, x, y, h, w }) => ({ i: id, x, y, h, w })),
  };
  return (
    <ResponsiveReactGridLayout
      compactType={null}
      className={styles.layout}
      layouts={layouts}
      breakpoints={{ lg: 0 }}
      cols={{ lg: 80 }}
      autoSize={true}
      rowHeight={10}
    >
      {tiles.map(({ id, content }) => (
        <div key={id} className="bg-gray-300">
          {content}
        </div>
      ))}
    </ResponsiveReactGridLayout>
  );
};

const styles = {
  layout: css`
    height: calc(100vh - 3rem);
  `,
};
