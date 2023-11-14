import React, { useEffect } from "react";
import "../assets/styles/basePage.css";
import "./styles/Home.css";
import Header from "../assets/components/Header";
import { toast } from "react-toastify";
import { useNavigate } from "react-router-dom";

export const mapList = [
  "1-1",
  "1-2",
  "1-3",
  "1-4",
  "1-5",
  "1-6",
  "1-7",
  "2-1",
  "3-1",
  "4-1",
];

function HomePage() {
  return (
    <div className="basePage homePage">
      <Header />

      <div className="content">
        {mapList.map((el) => (
          <MapItem mapId={el} key={el}/>
        ))}
      </div>
    </div>
  );
}

function MapItem({ mapId }) {
  const nav = useNavigate()
  
  function Redirect(){
   nav("/map/"+mapId)
  }

  return <div onClick={Redirect} className="mapItem">{mapId}</div>;
}

export default HomePage;
