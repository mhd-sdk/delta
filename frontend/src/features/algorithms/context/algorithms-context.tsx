import useDialogState from '@/hooks/use-dialog-state';
import React, { useState } from 'react';
import { Algorithm } from '../data/algorithms';

type UsersDialogType = 'edit';

interface UsersContextType {
  open: UsersDialogType | null;
  setOpen: (str: UsersDialogType | null) => void;
  currentRow: Algorithm | null;
  setCurrentRow: React.Dispatch<React.SetStateAction<Algorithm | null>>;
}

const AlgorithmsContext = React.createContext<UsersContextType | null>(null);

interface Props {
  children: React.ReactNode;
}

export default function UsersProvider({ children }: Props) {
  const [open, setOpen] = useDialogState<UsersDialogType>(null);
  const [currentRow, setCurrentRow] = useState<Algorithm | null>(null);

  return <AlgorithmsContext value={{ open, setOpen, currentRow, setCurrentRow }}>{children}</AlgorithmsContext>;
}

export const useAlgorithms = () => {
  const algorithmsContext = React.useContext(AlgorithmsContext);

  if (!algorithmsContext) {
    throw new Error('useAlgorithms has to be used within <UsersContext>');
  }

  return algorithmsContext;
};
