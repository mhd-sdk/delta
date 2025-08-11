import { JSX } from 'react';
import RGL, { WidthProvider } from 'react-grid-layout';
import { Panel } from '../panel';
import { DraggablePanel } from './draggablePanel';

interface Props {
  panels: Panel[];
  onChange: (panels: Panel[]) => void;
}
const ReactGridLayout = WidthProvider(RGL);

export const Grid = ({ panels, onChange }: Props): JSX.Element => {
  const handleLayoutChange = (layout: ReactGridLayout.Layout[]) => {
    const newPanels: Panel[] = [];
    layout.forEach((item) => {
      const panel = panels.find((p) => p.id === item.i);
      if (panel) {
        newPanels.push({
          ...panel,
          x: item.x,
          y: item.y,
          width: item.w,
          height: item.h,
        });
      }
      return null;
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

  const layout = panels.map((panel) => ({ i: panel.id, x: panel.x, y: panel.y, w: panel.width, h: panel.height, minW: 5, minH: 3 }));

  return (
    <ReactGridLayout
      layout={layout}
      onLayoutChange={handleLayoutChange}
      useCSSTransforms={true}
      compactType="vertical"
      cols={20}
      autoSize={true}
      rowHeight={50}
      draggableHandle=".drag-handle"
      draggableCancel=".drag-cancel"
      className="w-full"
    >
      {panels.map((panel) => (
        // <Chart onDelete={() => handleDelete(panel.id)} config={panel.data.config as ChartConfig} onConfigChange={handleConfigChange} />
        <div key={panel.id} data-grid={{ i: panel.id, x: panel.x, y: panel.y, w: panel.width, h: panel.height }}>
          <DraggablePanel onDelete={() => handleDelete(panel.id)} onConfigChange={handleConfigChange} panel={panel} />
        </div>
      ))}
    </ReactGridLayout>
  );
};
