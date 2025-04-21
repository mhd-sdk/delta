import { createFileRoute, redirect } from '@tanstack/react-router';

// Define the index route with a redirect to authenticated route
export const Route = createFileRoute('/')({
  beforeLoad: () => {
    // Check if user is authenticated and redirect accordingly
    const authData = localStorage.getItem('auth');
    if (authData) {
      const parsedAuth = JSON.parse(authData);
      if (parsedAuth?.isAuthenticated) {
        return redirect({ to: '/dashboard' });
      }
    }
    return redirect({ to: '/sign-in' });
  },
  component: () => null,
});
