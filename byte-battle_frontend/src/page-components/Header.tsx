import React from "react";
import "../../src/index.css"

const Header = () => {
  return (
    <div className=" bg-gradient-to-b from-header to-header fixed top-0 w-full z-10 h-16 p-2 flex flex-row justify-between items-center">
      <div className="w-1/4 px-4">
        <a className="inline-block" href="#">LOGO</a> 
      </div>
      <div className="w-3/4 px-4 flex flex-row justify-between">
        <a className="inline-block" href="#">Курсы</a>      
        <a className="inline-block" href="#">Войти</a>  
      </div>     
    </div>
  )
}

export default Header