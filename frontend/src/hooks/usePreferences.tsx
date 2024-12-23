import { useContext } from 'react';
import { PreferencesContextProps, preferencesCtx } from '../components/PreferencesProvider/PreferencesProvider';

// Custom hook to access the game context
export const usePreferences = (): PreferencesContextProps => {
  const context = useContext(preferencesCtx);

  if (!context) {
    throw new Error('usePreferences must be used within a GameWrapper');
  }

  return context;
};
