import SignIn from '@/features/sign-in';
import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/(auth)/sign-in')({
  component: SignIn,
});
