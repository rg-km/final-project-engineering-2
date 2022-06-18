import { createContext, useState, useMemo, useEffect } from "react";
import { getSession } from "../components/api/auth";

export const SessionContext = createContext({
  // TODO: answer here
});

export const SessionProvider = ({ children }) => {
  // TODO: answer here
  const [isLoggedIn, setIsLoggedIn] = useState(null);
  const [user, setUser] = useState({});

  useEffect(() => {
    const loadSession = async () => {
      const session = await getSession();
      setUser(session?.data);
      if (session?.status === 200) {
        setIsLoggedIn(true);
      } else {
        setIsLoggedIn(false);
      }
    };
    loadSession();
  }, []);

  const value = useMemo(
    () => ({
      user,
      setUser,
      isLoggedIn,
      setIsLoggedIn,
    }),
    [isLoggedIn, user]
  );

  return (
    <SessionContext.Provider value={value}>{children}</SessionContext.Provider>
  );
};
