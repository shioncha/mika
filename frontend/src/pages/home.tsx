import { Link } from "react-router";

import style from "../styles/pages/home.module.css";

function HomePage() {
  return (
    <div className={style.container}>
      <h1 className={style.logo}>Mika</h1>
      <p>Mika is a open-source personal journal app.</p>
      <div className={style.links}>
        <Link to="/signup" className={style.link}>
          Try it on the demo site
        </Link>
        <Link to="https://github.com/shioncha/mika" className={style.link}>
          Host your own
        </Link>
      </div>
      <div className={style.features}>
        <h2>Features</h2>
        <div className={style.featureList}>
          <div className={style.feature}>
            <h3>Track your daily life</h3>
            <p>Write about your daily life, thoughts, and feelings.</p>
          </div>
          <div className={style.feature}>
            <h3>Manage with tags</h3>
            <p>You can add tags to your posts to categorize them.</p>
          </div>
          <div className={style.feature}>
            <h3>Set a reminder</h3>
            <p>Set reminders for your important things. (coming soon)</p>
          </div>
          <div className={style.feature}>
            <h3>Manage Tasks</h3>
            <p>
              Manage your tasks and to-do lists. You only need to add #task to
              your post. (coming soon)
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default HomePage;
