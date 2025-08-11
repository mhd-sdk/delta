import { Button } from '@/components/ui/button';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { cn } from '@/lib/utils';
import { useWebAuthnStore } from '@/stores/webAuthnStore';
import { zodResolver } from '@hookform/resolvers/zod';
import { useNavigate } from '@tanstack/react-router';
import { HTMLAttributes, useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { z } from 'zod';

type UserAuthFormProps = HTMLAttributes<HTMLFormElement>;

const formSchema = z.object({
  username: z
    .string()
    .min(1, {
      message: 'Please enter your username',
    })
    .min(3, {
      message: 'Username must be at least 3 characters long',
    }),
});

export const SignInForm = ({ className, ...props }: UserAuthFormProps) => {
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);
  const { login } = useWebAuthnStore();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: '',
    },
  });

  const onSubmit = async (data: z.infer<typeof formSchema>) => {
    setIsLoading(true);
    setError(null);

    try {
      await login(data.username);
      toast.success('Login successful', {
        description: "You've been successfully logged in",
      });
      navigate({ to: '/dashboard', replace: true });
    } catch (err) {
      console.error(err);
      setError(err instanceof Error ? err.message : 'Login failed. Please try again.');
      toast.error('Login failed', {
        description: err instanceof Error ? err.message : 'Login failed. Please try again.',
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className={cn('grid gap-3', className)} {...props}>
        {error && <div className="bg-destructive/15 text-destructive text-sm p-3 rounded-md">{error}</div>}

        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem className="relative">
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input placeholder="Enter your username" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="mt-2" disabled={isLoading}>
          {isLoading ? 'Processing...' : 'Login with WebAuthn'}
        </Button>

        <p className="text-xs text-muted-foreground mt-2">You will be prompted to use your device's biometric authentication</p>
      </form>
    </Form>
  );
};
