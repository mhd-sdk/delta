import { Button } from '@/components/ui/button';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';

import { cn } from '@/lib/utils';
import { useWebAuthnStore } from '@/stores/webAuthnStore';
import { zodResolver } from '@hookform/resolvers/zod';
import { HTMLAttributes, useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import { z } from 'zod';

type SignUpFormProps = HTMLAttributes<HTMLFormElement>;

const formSchema = z.object({
  username: z.string().min(1, { message: 'Please enter your username' }).min(3, { message: 'Username must be at least 3 characters' }),
});

export const SignUpForm = ({ className, ...props }: SignUpFormProps) => {
  const [isLoading, setIsLoading] = useState(false);
  // const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);
  const { register } = useWebAuthnStore();

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
      await register(data.username);
      toast.success('Registration successful', {
        description: "You've been successfully registered and logged in",
      });
      // navigate({ to: '/dashboard', replace: true });
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (err: any) {
      console.error(err);
      setError(err.response.data.error);
      toast.error('Registration failed', {
        description: err.response.data.error || 'An error occurred during registration. Please try again.',
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
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit" className="mt-2" disabled={isLoading}>
          {isLoading ? 'Processing...' : 'Register with WebAuthn'}
        </Button>

        <p className="text-xs text-muted-foreground mt-2">You will be prompted to use your device's biometric authentication</p>
      </form>
    </Form>
  );
};
