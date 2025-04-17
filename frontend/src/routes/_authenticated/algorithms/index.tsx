import Algorithms from '@/features/algorithms';
import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/_authenticated/algorithms/')({
  component: Algorithms,
});
