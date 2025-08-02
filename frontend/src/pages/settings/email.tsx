import Button from "../../components/elements/Button";
import InputText from "../../components/elements/InputText";
import apiClient from "../../libs/api";

function EmailPage() {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    apiClient
      .patch("/account", {
        email: event.currentTarget.email.value,
      })
      .then(() => {
        alert("Email updated successfully!");
      })
      .catch((error) => {
        console.error("Error updating email:", error);
        alert("Failed to update email. Please try again.");
      });
  };

  return (
    <div>
      <h1>Change Email</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">New Email</label>
          <InputText type="email" id="email" name="email" required />
          <Button type="submit">Update Email Address</Button>
        </div>
      </form>
    </div>
  );
}

export default EmailPage;
