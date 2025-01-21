import { ReactNode, createContext, useEffect, useState } from 'react';
import { GetAppData, SaveAppData } from '../../../wailsjs/go/app/App';
import { persistence } from '../../../wailsjs/go/models';

export interface AppDataContextProps {
  appData: persistence.AppData;
  onSave: (value: persistence.AppData) => void;
  refetch: () => void;
}

export const appDataCtx = createContext<AppDataContextProps | undefined>(undefined);

interface Props {
  children?: ReactNode;
}
const AppDataProvider = ({ children }: Props) => {
  const [appData, setAppData] = useState<persistence.AppData>();

  useEffect(() => {
    GetAppData().then((data) => {
      setAppData(data);
    });
  }, []);

  const onSave = (value: persistence.AppData) => {
    setAppData(value);
    SaveAppData(value);
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
