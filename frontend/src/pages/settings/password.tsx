import apiClient from "../../libs/api";

function PasswordPage() {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    apiClient
      .patch("/account", {
        password: event.currentTarget.password.value,
      })
      .then(() => {
        alert("Password updated successfully!");
      })
      .catch((error) => {
        console.error("Error updating password:", error);
        alert("Failed to update password. Please try again.");
      });
  };

  return (
    <div>
      <h1>Change Password</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="password">Password:</label>
          <input type="password" id="password" name="password" required />
          <button type="submit">Update Password</button>
        </div>
      </form>
    </div>
  );
}

export default PasswordPage;
