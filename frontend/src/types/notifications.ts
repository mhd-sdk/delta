export interface INotification {
  id?: number;
  title: string;
  subtitle?: string;
  type: 'error' | 'info' | 'info-square' | 'success' | 'warning' | 'warning-alt';
  caption?: string;
}
