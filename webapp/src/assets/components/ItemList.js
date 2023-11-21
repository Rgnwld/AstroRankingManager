import React from "react";

export default function ItemList({ info }) {
  const { id, username, timeInMiliSeconds } = info;

  return (
    <ol className="itemObject">
      <span>{username}:</span> <span>{timeInMiliSeconds}</span>
    </ol>
  );
}
