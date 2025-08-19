import UnauthorisedError from '@/features/errors/unauthorized-error';
import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/(errors)/401')({
  component: UnauthorisedError,
});
