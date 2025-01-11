import { ReactNode, createContext, useEffect, useState } from 'react';
import { GetAppData, SaveAppData } from '../../../wailsjs/go/main/App';
import { persistence } from '../../../wailsjs/go/models';

export interface AppDataContextProps {
  appData: persistence.AppData;
  onSave: (value: persistence.AppData) => void;
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

  if (!appData) {
    return <div>Could not load app data</div>;
  }

  return <appDataCtx.Provider value={{ appData, onSave }}>{children}</appDataCtx.Provider>;
};

export default AppDataProvider;
