import { createContext, useContext, useEffect, useState } from 'react';

type Theme = 'dark' | 'light';

type ThemeProviderProps = {
  children: React.ReactNode;
  defaultTheme?: Theme;
  storageKey?: string;
};

type ThemeProviderState = {
  theme: Theme;
  setTheme: (theme: Theme) => void;
};

const initialState: ThemeProviderState = {
  theme: 'dark',
  setTheme: () => null,
};

const ThemeProviderContext = createContext<ThemeProviderState>(initialState);

export function ThemeProvider({ children, defaultTheme = 'dark', storageKey = 'vite-ui-theme', ...props }: ThemeProviderProps) {
  const [theme, _setTheme] = useState<Theme>(() => (localStorage.getItem(storageKey) as Theme) || defaultTheme);

  useEffect(() => {
    const root = window.document.documentElement;
    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');

    const applyTheme = (theme: Theme) => {
      root.classList.remove('light', 'dark'); // Remove existing theme classes
      const systemTheme = mediaQuery.matches ? 'dark' : 'light';
      const effectiveTheme = theme === 'dark' ? systemTheme : theme;
      root.classList.add(effectiveTheme); // Add the new theme class
    };

    const handleChange = () => {
      if (theme === 'dark') {
        applyTheme('dark');
      }
    };

    applyTheme(theme);

    mediaQuery.addEventListener('change', handleChange);

    return () => mediaQuery.removeEventListener('change', handleChange);
  }, [theme]);

  const setTheme = (theme: Theme) => {
    localStorage.setItem(storageKey, theme);
    _setTheme(theme);
  };

  const value = {
    theme,
    setTheme,
  };

  return (
    <ThemeProviderContext.Provider {...props} value={value}>
      {children}
    </ThemeProviderContext.Provider>
  );
}

export const useTheme = () => {
  const context = useContext(ThemeProviderContext);

  if (context === undefined) throw new Error('useTheme must be used within a ThemeProvider');

  return context;
};
