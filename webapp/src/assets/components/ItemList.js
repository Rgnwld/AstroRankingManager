import React from 'react'

export default function ItemList({ info }) {
    const {id,username, timeInSeconds} = info

  return (
    <div>{username} - {timeInSeconds}</div>
  )
}
