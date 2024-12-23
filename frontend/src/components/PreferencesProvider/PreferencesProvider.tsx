import { ReactNode, createContext, useState } from 'react';
import { Preferences } from '../../types/preferences';

export interface PreferencesContextProps {
  preferences: Preferences;
  onPreferencesChange: (preferences: Preferences) => void;
  onSave: () => void;
  onCancel: () => void;
}

export const preferencesCtx = createContext<PreferencesContextProps | undefined>(undefined);

interface Props {
  children?: ReactNode;
}
const PreferencesProvider = ({ children }: Props) => {
  const [isDirty, setIsDirty] = useState(false);
  const [preferences, setPreferences] = useState<Preferences>({} as Preferences);
  const onPreferencesChange = (preferences: Preferences) => {
    setPreferences(preferences);
    setIsDirty(true);
  };
  const onCancel = () => {
    setIsDirty(false);
  };

  const onSave = () => {
    setIsDirty(false);
  };
  return <preferencesCtx.Provider value={{ preferences, onPreferencesChange, onSave, onCancel }}>{children}</preferencesCtx.Provider>;
};

export default PreferencesProvider;
