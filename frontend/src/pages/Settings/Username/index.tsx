import Button from "../../../components/ui/Button";
import InputText from "../../../components/ui/InputText";
import apiClient from "../../../lib/api";

function UsernamePage() {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    apiClient
      .patch("/account", {
        name: event.currentTarget.username.value,
      })
      .then(() => {
        alert("Username updated successfully!");
      })
      .catch(() => {
        console.error("Error updating username:");
        alert("Failed to update username. Please try again.");
      });
  };

  return (
    <div>
      <h1>Change Name</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="username">New Name</label>
          <InputText type="text" id="username" name="username" required />
          <Button type="submit">Update Name</Button>
        </div>
      </form>
    </div>
  );
}

export default UsernamePage;
