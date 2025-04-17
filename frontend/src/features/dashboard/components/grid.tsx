import { JSX } from 'react';
import ReactGridLayout from 'react-grid-layout';
import { Panel } from '../panel';
import { DraggablePanel } from './draggablePanel';

interface Props {
  panels: Panel[];
  onChange: (panels: Panel[]) => void;
}

export const Grid = ({ panels, onChange }: Props): JSX.Element => {
  const handleLayoutChange = (layout: ReactGridLayout.Layout[]) => {
    const newPanels = panels.map((panel) => {
      const newLayout = layout.find((l: ReactGridLayout.Layout) => l.i === panel.id);
      if (newLayout) {
        return { ...panel, x: newLayout.x, y: newLayout.y, w: newLayout.w, h: newLayout.h };
      }
      return panel;
    });
    onChange(newPanels);
  };

  const handleDelete = (id: string) => {
    const newPanels = panels.filter((panel) => panel.id !== id);
    onChange(newPanels);
  };

  const handleConfigChange = (panel: Panel) => {
    const updatedPanels = panels.map((t) => (t.id === panel.id ? panel : t));
    onChange(updatedPanels);
  };

  const layout = panels.map((panel) => ({ i: panel.id, x: panel.x, y: panel.y, w: panel.width, h: panel.height }));

  return (
    <ReactGridLayout
      layout={layout}
      onLayoutChange={handleLayoutChange}
      useCSSTransforms={true}
      compactType="vertical"
      cols={50}
      autoSize={true}
      rowHeight={10}
      // draggableHandle=".drag-handle"
      draggableCancel=".drag-cancel"
    >
      {panels.map((panel) => (
        <div key={panel.id}>
          <DraggablePanel panel={panel} onDelete={handleDelete} onConfigChange={handleConfigChange} />
        </div>
      ))}
    </ReactGridLayout>
  );
};
