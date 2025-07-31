import { useState } from "react";
import { Link, useNavigate } from "react-router";

import Button from "../components/elements/Button.tsx";
import InputText from "../components/elements/InputText.tsx";
import { useAuth } from "../hooks/auth_context.tsx";
import type { AuthResponse, SignUpCredentials } from "../libs/AuthService";
import { authService } from "../libs/AuthService";
import style from "../styles/pages/signup.module.css";

function SignUpPage() {
  const { signIn } = useAuth();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordConfirm, setPasswordConfirm] = useState("");

  const navigate = useNavigate();

  async function handleSignUp(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!email || !password) {
      setError("Email and password are required.");
      return;
    }

    setLoading(true);
    try {
      const credentials: SignUpCredentials = {
        name,
        email,
        password,
        password_confirm: passwordConfirm,
      };
      const response: AuthResponse = await authService.signUp(credentials);

      if (response.token) {
        signIn(response.token);
        navigate("/");
      } else {
        setError("Invalid email or password. Please try again.");
      }
    } catch {
      console.error("Sign up failed:");
      setError("An error occurred while signing up. Please try again later.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className={style.container}>
      <h1>Sign Up</h1>
      <p>Create your account</p>
      {error && <p className={style.error}>{error}</p>}
      <form autoComplete="on" onSubmit={handleSignUp} className={style.form}>
        <div className={style.formGroup}>
          <label htmlFor="name" className={style.label}>
            Name
          </label>
          <InputText
            id="name"
            name="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            aria-required="true"
            aria-invalid={error ? "true" : "false"}
          />
        </div>
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
            className={style.input}
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
            className={style.input}
          />
        </div>
        <div className={style.formGroup}>
          <label htmlFor="passwordConfirm" className={style.label}>
            Confirm Password
          </label>
          <InputText
            hasPasswordMask
            id="passwordConfirm"
            name="passwordConfirm"
            type="password"
            value={passwordConfirm}
            onChange={(e) => setPasswordConfirm(e.target.value)}
            minLength={8}
            required
            aria-required="true"
            aria-invalid={error ? "true" : "false"}
            className={style.input}
          />
        </div>
        <Button type="submit" disabled={loading}>
          {loading ? "Loading..." : "Sign Up"}
        </Button>
      </form>
      <p>
        Already have an account? <Link to="/signin">Sign In</Link>
      </p>
    </div>
  );
}

export default SignUpPage;
