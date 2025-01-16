export const getThemeCode = (code: string) => {
  const themeCode = code === 'light' ? 'g10' : 'g90';
  return themeCode;
};
