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
      <h1>Change Username</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">Username:</label>
          <input type="text" id="name" name="name" required />
          <button type="submit">Update Username</button>
        </div>
      </form>
    </div>
  );
}

export default UsernamePage;
