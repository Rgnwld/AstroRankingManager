import React, { useEffect, useState } from "react";
import secondsToTime from "../aux/time";

export default function ItemList({ info }) {
  const { id, username, timeInMiliSeconds } = info;
  const [gameTime,setGameTime]=useState("");

  useEffect(() => {
    setGameTime(secondsToTime(timeInMiliSeconds))
  }, [info])


  return (
    <ol className="itemObject">
       <span>{gameTime}</span> - <span>{username}</span>
    </ol>
  );
}
