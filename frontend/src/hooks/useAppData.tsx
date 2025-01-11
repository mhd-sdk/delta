import { useContext } from 'react';
import { AppDataContextProps, appDataCtx } from '../components/AppDataProvider/AppDataProvider';

// Custom hook to access the game context
export const useAppData = (): AppDataContextProps => {
  const context = useContext(appDataCtx);

  if (!context) {
    throw new Error('usePreferences must be used within a GameWrapper');
  }

  return context;
};
