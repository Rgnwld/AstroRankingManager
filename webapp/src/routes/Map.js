import React, { useEffect, useState } from "react";
import "../assets/styles/basePage.css";
import "./styles/Home.css";
import Header from "../assets/components/Header";
import jscookie from "js-cookie";
import axios from "axios";
import { toast } from "react-toastify";
import { useParams } from "react-router-dom";
import ItemList from "../assets/components/ItemList";

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

function MapPage() {
  const {mapId} = useParams();

  const [mapInfo, setMapInfo] = useState([])

  async function GetMapInfo() {
    const res = await axios.get(
      "http://localhost:8080/v1/ranking/" +
        mapId +
        "?token=" +
        jscookie.get("access_token")
    );

    setMapInfo(res.data)
  }

  useEffect(() => {
    console.log("test")
    try {
      console.log(mapId)
      GetMapInfo();
    } catch (e) {
      console.error(e);
    }
  }, []);

  return (
    <div className="basePage homePage">
      <Header />
      {mapInfo.map(e => 
        <ItemList info={e} key={e.id}/>
      )}
      <div className="content"></div>
    </div>
  );
}

export default MapPage;
