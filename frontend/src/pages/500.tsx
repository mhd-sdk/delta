import { useNavigate } from '@tanstack/react-router';
import { FC } from 'react';

const ServerError: FC = () => {
  const navigate = useNavigate();

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-background p-4">
      <h1 className="text-6xl font-bold text-red-500 mb-4">500</h1>
      <p className="text-xl mb-8">Server Error</p>
      <button onClick={() => navigate({ to: '/' })} className="px-4 py-2 bg-primary text-white rounded-md hover:bg-primary/90 transition-colors">
        Try Again
      </button>
    </div>
  );
};

export default ServerError;
