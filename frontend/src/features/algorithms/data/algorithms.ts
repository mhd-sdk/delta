export type Algorithm = {
  id: string;
  name: string;
  status: 'Running' | 'Stopped';
  balance: number; // Optional balance field
};
export const algorithms: Algorithm[] = [
  {
    id: '1',
    name: 'Algorithm A',
    status: 'Running',
    balance: 1000, // Example balance field
  },
  {
    id: '2',
    name: 'Algorithm B',
    status: 'Stopped',
    balance: 500, // Example balance field
  },
  {
    id: '3',
    name: 'Algorithm C',
    status: 'Running',
    balance: 750, // Example balance fieldw
  },
];
