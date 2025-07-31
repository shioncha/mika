import { useState } from "react";
import { Link, useNavigate } from "react-router";

import Button from "../components/elements/Button.tsx";
import InputText from "../components/elements/InputText.tsx";
import { useAuth } from "../hooks/auth_context.tsx";
import type { AuthResponse, SignInCredentials } from "../libs/AuthService";
import { authService } from "../libs/AuthService";
import style from "../styles/pages/signin.module.css";

function SignInPage() {
  const { signIn } = useAuth();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  async function handleSignIn(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!email || !password) {
      setError("Email and password are required.");
      return;
    }

    setLoading(true);
    try {
      const credentials: SignInCredentials = { email, password };
      const response: AuthResponse = await authService.signIn(credentials);

      if (response && response.token !== "") {
        signIn(response.token);
        navigate("/");
      } else {
        setError("Invalid email or password. Please try again.");
      }
    } catch {
      console.error("Sign in failed:");
      setError("An error occurred while signing in. Please try again later.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className={style.container}>
      <h1>Sign In</h1>
      <p>Sign in to your account</p>
      {error && <p className={style.error}>{error}</p>}
      <form autoComplete="on" onSubmit={handleSignIn} className={style.form}>
        <div className={style.formGroup}>
          <label htmlFor="email" className={style.label}>
            Email
          </label>
          <InputText
            id="email"
            name="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            aria-required="true"
            aria-invalid={error ? "true" : "false"}
          />
        </div>
        <div className={style.formGroup}>
          <label htmlFor="password" className={style.label}>
            Password
          </label>
          <InputText
            hasPasswordMask
            id="password"
            name="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            minLength={8}
            required
            aria-required="true"
            aria-invalid={error ? "true" : "false"}
          />
        </div>
        <Button type="submit" disabled={loading}>
          {loading ? "Loading..." : "Sign In"}
        </Button>
      </form>
      <p>
        Don't have an account? <Link to="/signup">Sign Up</Link>
      </p>
    </div>
  );
}

export default SignInPage;
