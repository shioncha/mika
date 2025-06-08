import { useContext, useState } from "react";
import { Link, useNavigate } from "react-router";

import { AuthContext } from "../hooks/auth_context.tsx";
import style from "../styles/pages/signup.module.css";

type SignUpResponse = {
  token: string;
  refresh_token: string;
};

function SignUpPage() {
  const { login } = useContext(AuthContext);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordConfirm, setPasswordConfirm] = useState("");

  const navigate = useNavigate();

  async function handleSignUp(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!name || !email || !password || !passwordConfirm) {
      setError("All fields are required.");
      return;
    }
    if (password !== passwordConfirm) {
      setError("Passwords do not match.");
      return;
    }

    setLoading(true);
    try {
      const response = await fetch("/api/sign-up", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name, email, password, password_confirm: passwordConfirm }),
      });
      const data: SignUpResponse = await response.json();

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
      <h1>Sign Up</h1>
      <p>Create your account</p>
      {error && <p className={style.error}>{error}</p>}
      <form autoComplete="on" onSubmit={handleSignUp} className={style.form}>
        <div className={style.formGroup}>
          <label htmlFor="name" className={style.label}>Name</label>
          <input
            id="name"
            name="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            aria-required="true"
            aria-invalid={error ? "true" : "false"}
            className={style.input}
          />
        </div>
        <div className={style.formGroup}>
          <label htmlFor="email" className={style.label}>Email</label>
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
          <label htmlFor="password" className={style.label}>Password</label>
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
        <div className={style.formGroup}>
          <label htmlFor="passwordConfirm" className={style.label}>Confirm Password</label>
          <input
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
        <button type="submit" disabled={loading} className={style.button}>
          {loading ? "Loading..." : "Sign Up"}
        </button>
      </form>
      <p>
        Already have an account? <Link to="/signin">Sign In</Link>
      </p>
    </div>
  );
}

export default SignUpPage;
