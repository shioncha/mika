import { useContext, useState } from "react";
import { Link, useNavigate } from "react-router";

import Button from "../components/elements/Button.tsx";
import { AuthContext } from "../hooks/auth_context.tsx";
import style from "../styles/pages/signin.module.css";

type SignInResponse = {
  token: string;
  refresh_token: string;
};

function SignInPage() {
  const { login } = useContext(AuthContext);
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
      const response = await fetch("/api/sign-in", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });
      const data: SignInResponse = await response.json();

      if (response.status == 200) {
        login(data.token, data.refresh_token);
        window.localStorage.setItem("token", data.token);
        window.localStorage.setItem("refresh_token", data.refresh_token);
        navigate("/");
      } else if (response.status == 400 || response.status == 401) {
        setError("Invalid email or password. Please try again.");
      } else {
        setError("Internal server error. Please try again later.");
      }
    } catch {
      setError("Network error. Please check your connection.");
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
          <input
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
          <input
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
