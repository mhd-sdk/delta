import { useWebAuthnStore } from '@/stores/webAuthnStore';
import { FC } from 'react';

const DashboardPage: FC = () => {
  const { user, logout } = useWebAuthnStore();

  return (
    <div className="container mx-auto py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold">Dashboard</h1>
        <div className="flex items-center gap-4">
          <div className="text-sm">
            <span className="font-medium">Welcome, {user?.username}</span>
          </div>
          <button onClick={logout} className="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition-colors">
            Logout
          </button>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">{/* Dashboard content here */}</div>
    </div>
  );
};

export default DashboardPage;
