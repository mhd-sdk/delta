import { ChartConfig, Panel, PanelType } from '../panel';
import { AccountInfo } from './AccountInfo';
import { Chart } from './chart';

interface Props {
  panel: Panel;
  onDelete: (id: string) => void;
  onConfigChange: (panel: Panel) => void;
}

export const DraggablePanel = ({ panel, onDelete, onConfigChange }: Props) => {
  const handleChartChange = (config: ChartConfig) => {
    if (panel.data.type === PanelType.Chart) {
      onConfigChange({ ...panel, data: { type: PanelType.Chart, config } } as Panel);
    }
  };

  const handleDelete = () => {
    onDelete(panel.id);
  };

  const renderPanel = (panel: Panel) => {
    switch (panel.data.type) {
      case PanelType.Chart:
        return <Chart onDelete={handleDelete} config={panel.data.config} onConfigChange={handleChartChange} />;
      case PanelType.AccountInfo:
        return <AccountInfo />;
      default:
        return null;
    }
  };

  return renderPanel(panel);
};
