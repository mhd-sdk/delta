import WebAuthnForm from '../components/WebAuthnForm';

const RegisterPage = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-background p-4">
      <WebAuthnForm mode="register" />
    </div>
  );
};

export default RegisterPage;
