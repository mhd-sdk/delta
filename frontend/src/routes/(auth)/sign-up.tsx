import SignUp from '@/features/sign-up';
import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/(auth)/sign-up')({
  component: SignUp,
});
