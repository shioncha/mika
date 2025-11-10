import { useEffect, useState } from "react";

import Button from "../../../components/ui/Button";
import List from "../../../components/ui/List";
import ListElementSession from "../../../features/settings/components/ListElementSession";
import { useAuth } from "../../../hooks/useAuth";
import { getAllSessions, revokeAllSessions } from "../../../lib/AuthService";
import type { Session } from "../../../types/session";

function SessionPage() {
  const [sessions, setSessions] = useState<Session[] | null>(null);
  const { signOut } = useAuth();

  useEffect(() => {
    getAllSessions()
      .then((data) => {
        setSessions(data);
      })
      .catch(() => {
        setSessions([]);
      });
  }, []);

  const handleRevokeAllSessions = async () => {
    if (!sessions) return;

    await revokeAllSessions();
    signOut();
  };

  if (sessions === null) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <h1>Sessions</h1>
      <p>List of active sessions for your account.</p>
      <Button onClick={handleRevokeAllSessions}>
        Sign out of all sessions
      </Button>
      <List>
        {sessions.map((session) => (
          <ListElementSession
            key={session.ID}
            userAgent={session.DeviceInfo}
            ipAddress={session.IPAddress}
          />
        ))}
      </List>
    </>
  );
}

export default SessionPage;
