import { useContext, useState } from "react";
import { useNavigate } from "react-router";

import { AuthContext } from "../hooks/auth_context.tsx";

type SignInResponse = {
  token: string;
  refresh_token: string;
};

function SignInPage() {
  const { login } = useContext(AuthContext);
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  async function handleSignIn(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const email = formData.get("email") as string;
    const password = formData.get("password") as string;
    setLoading(true);
    const response = await fetch("/api/sign-in", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });
    setLoading(false);
    const data: SignInResponse = await response.json();
    if (response.status == 200) {
      login(data.token, data.refresh_token);
      window.localStorage.setItem("token", data.token);
      window.localStorage.setItem("refresh_token", data.refresh_token);
      navigate("/");
    }
    if (response.status == 401) {
      alert("Invalid email or password");
    }
    if (response.status == 500) {
      alert("Server error");
    }
  }

  return (
    <div>
      <h1>Sign In</h1>
      <p>Sign in to your account</p>
      <form onSubmit={handleSignIn}>
        <label>
          Email:
          <input type="email" name="email" required />
        </label>
        <br />
        <label>
          Password:
          <input type="password" name="password" required />
        </label>
        <button type="submit" disabled={loading}>
          {loading ? "Loading..." : "Sign In"}
        </button>
      </form>
    </div>
  );
}

export default SignInPage;
