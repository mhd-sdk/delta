import { useEffect } from 'react';
import { useAppData } from '../../hooks/useAppData';

interface Props {}

export const AccountOverview = ({}: Props): JSX.Element => {
  const { appData } = useAppData();

  useEffect(() => {}, [appData]);
  return <div>AccountOverview</div>;
};
