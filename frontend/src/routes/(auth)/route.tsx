import { createFileRoute, Outlet } from '@tanstack/react-router';

// Auth layout route
export const Route = createFileRoute('/(auth)')({
  component: AuthLayout,
});

function AuthLayout() {
  return (
    <div className="auth-layout">
      <Outlet />
    </div>
  );
}
