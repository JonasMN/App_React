import React, { useEffect, useState } from "react";
import GetGifs from "../services/GetGifs";
import Gif from "./gif";

export default function ListOfGifs({ params }) {
  const { keyword } = params;

  const [gifs, setGifs] = useState({loading: false,results: []});



  useEffect(function () {
      setGifs(
        actualGifs=>({loading:true,results:actualGifs.results})
        )

      GetGifs({ keyword })
      .then((gifs) => { 
        setGifs({loading:false,results:gifs})
      });
    },[keyword]);

    if(gifs.loading) return <i>Cargando 🚀</i>

    

  return (
    <div>
      {gifs.results.map(({ id, title, url }) => (
        <Gif id={id} key={id} title={title} url={url} />
      ))}
    </div>
  );
}
