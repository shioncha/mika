import Button from "../../../components/ui/Button";
import InputText from "../../../components/ui/InputText";
import apiClient from "../../../lib/api";

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
          <label htmlFor="password">New Password</label>
          <InputText
            hasPasswordMask={true}
            type="password"
            id="password"
            name="password"
            required
          />
          <Button type="submit">Update Password</Button>
        </div>
      </form>
    </div>
  );
}

export default PasswordPage;
