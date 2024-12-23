export interface GeneralSettings {
  language: 'French' | 'English';
}

export interface AppearanceSettings {
  theme: 'Light' | 'Dark';
}

export interface Preferences {
  general: GeneralSettings;
  appearance: AppearanceSettings;
}
