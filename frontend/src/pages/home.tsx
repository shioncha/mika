import { Link } from "react-router";

function HomePage() {
  return (
    <div>
      <h1>Mika</h1>
      <p>This is a lifelog app.</p>
      <Link to="/signin">Sign In</Link>
      <Link to="/signup">Sign Up</Link>
    </div>
  );
}

export default HomePage;
