import Button from "../../components/elements/Button";
import InputText from "../../components/elements/InputText";
import apiClient from "../../libs/api";

function UsernamePage() {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    apiClient
      .patch("/account", {
        name: event.currentTarget.name.value,
      })
      .then(() => {
        alert("Username updated successfully!");
      })
      .catch((error) => {
        console.error("Error updating username:", error);
        alert("Failed to update username. Please try again.");
      });
  };

  return (
    <div>
      <h1>Change Name</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">New Name</label>
          <InputText type="text" id="name" name="name" required />
          <Button type="submit">Update Name</Button>
        </div>
      </form>
    </div>
  );
}

export default UsernamePage;
