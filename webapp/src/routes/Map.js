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

function MapPage() {
  return (
    <div className="basePage homePage">
      <Header />

      <div className="content">
      
      </div>
    </div>
  );
}

export default MapPage;
