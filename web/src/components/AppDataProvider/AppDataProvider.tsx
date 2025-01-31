import { ReactNode, createContext, useEffect, useState } from 'react';
import { GetAppData, SaveAppData } from '../../../wailsjs/go/app/App';
import { models } from '../../../wailsjs/go/models';

export interface AppDataContextProps {
  appData: models.AppData;
  onSave: (value: models.AppData) => void;
  refetch: () => void;
}

export const appDataCtx = createContext<AppDataContextProps | undefined>(undefined);

interface Props {
  children?: ReactNode;
}
const AppDataProvider = ({ children }: Props) => {
  const [appData, setAppData] = useState<models.AppData>();

  useEffect(() => {
    GetAppData().then((data) => {
      setAppData(data);
    });
  }, []);

  const onSave = async (value: models.AppData) => {
    await SaveAppData(value);
    refetch();
  };

  const refetch = async () => {
    const appData = await GetAppData();
    setAppData(appData);
  };

  if (!appData) {
    return <div>Could not load app data</div>;
  }

  return <appDataCtx.Provider value={{ appData, onSave, refetch }}>{children}</appDataCtx.Provider>;
};

export default AppDataProvider;
