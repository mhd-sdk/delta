import Alerts from '@/features/alerts';
import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/_authenticated/alerts/')({
  component: Alerts,
});
