import { useNavigate } from '@tanstack/react-router';
import { FC } from 'react';

const NotFound: FC = () => {
  const navigate = useNavigate();

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-background p-4">
      <h1 className="text-6xl font-bold text-primary mb-4">404</h1>
      <p className="text-xl mb-8">Page not found</p>
      <button onClick={() => navigate({ to: '/' })} className="px-4 py-2 bg-primary text-white rounded-md hover:bg-primary/90 transition-colors">
        Return to Dashboard
      </button>
    </div>
  );
};

export default NotFound;
