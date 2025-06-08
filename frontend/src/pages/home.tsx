import { Link } from "react-router";
import style from "../styles/pages/home.module.css";

function HomePage() {
  return (
    <div className={style.container}>
      <h1>Mika</h1>
      <p>This is a lifelog app.</p>
      <Link to="/signin" className={style.link}>Sign In</Link>
      <Link to="/signup" className={style.link}>Sign Up</Link>
    </div>
  );
}

export default HomePage;
