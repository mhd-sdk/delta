import { useWebAuthnStore } from '@/stores/webAuthnStore';
import { handleServerError } from '@/utils/handle-server-error';
import { QueryCache, QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RouterProvider, createRouter } from '@tanstack/react-router';
import { AxiosError } from 'axios';
import { StrictMode, useEffect } from 'react';
import ReactDOM from 'react-dom/client';
import { toast } from 'sonner';
import { FontProvider } from './context/font-context';
import { ThemeProvider } from './context/theme-context';
import './index.css';
import '/node_modules/react-grid-layout/css/styles.css';
import '/node_modules/react-resizable/css/styles.css';

// Import generated routes
import { routeTree } from './routeTree.gen';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: (failureCount, error) => {
        // eslint-disable-next-line no-console
        if (import.meta.env.DEV) console.log({ failureCount, error });

        if (failureCount >= 0 && import.meta.env.DEV) return false;
        if (failureCount > 3 && import.meta.env.PROD) return false;

        return !(error instanceof AxiosError && [401, 403].includes(error.response?.status ?? 0));
      },
      refetchOnWindowFocus: import.meta.env.PROD,
      staleTime: 10 * 1000, // 10s
    },
    mutations: {
      onError: (error) => {
        handleServerError(error);

        if (error instanceof AxiosError) {
          if (error.response?.status === 304) {
            toast.error('Content not modified!');
          }
        }
      },
    },
  },
  queryCache: new QueryCache({
    onError: (error) => {
      if (error instanceof AxiosError) {
        if (error.response?.status === 401) {
          toast.error('Session expired!');
          useWebAuthnStore.getState().logout();
          router.navigate({ to: '/login' });
        }
        if (error.response?.status === 500) {
          toast.error('Internal Server Error!');
          router.navigate({ to: '/500' });
        }
        if (error.response?.status === 403) {
          // router.navigate("/forbidden", { replace: true });
        }
      }
    },
  }),
});

// Create a new router instance
const router = createRouter({
  routeTree,
  context: { queryClient },
  defaultPreload: 'intent',
  defaultPreloadStaleTime: 0,
});

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

// Auth token checker component
const AuthChecker = () => {
  const { checkAuth } = useWebAuthnStore();

  useEffect(() => {
    // Initial auth check
    const initialCheck = async () => {
      await checkAuth();
    };
    initialCheck();

    // Check auth every 2 hours
    const interval = setInterval(
      async () => {
        const isStillValid = await checkAuth();
        if (!isStillValid) {
          const pathname = window.location.pathname;
          if (pathname !== '/login' && pathname !== '/register') {
            router.navigate({ to: '/login' });
            toast.error('Session expired. Please login again.');
          }
        }
      },
      2 * 60 * 60 * 1000
    ); // 2 hours

    return () => clearInterval(interval);
  }, [checkAuth]);

  return null;
};

// Render the app
const rootElement = document.getElementById('root')!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
          <FontProvider>
            <AuthChecker />
            <RouterProvider router={router} />
          </FontProvider>
        </ThemeProvider>
      </QueryClientProvider>
    </StrictMode>
  );
}
