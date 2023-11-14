import React, { createContext, useContext, useState } from "react";

const UserContext = createContext();

export function usePlayer() {
  const [player, setPlayer] = useContext(UserContext);
  return [player, setPlayer];
}

export function PlayerProvider({children}) {
  const [player, setPlayer] = useState({ username: "" });

  return (
    <UserContext.Provider value={[player, setPlayer]}>{children}</UserContext.Provider>
  );
}
