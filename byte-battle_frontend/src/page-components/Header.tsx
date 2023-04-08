import React from "react";
import "../../src/index.css"

const Header = ():JSX.Element => {
  return (
    <div className="header">
      <div className="w-1/4 px-4">
        <a className="inline-block" href="#">LOGO</a> 
      </div>
      <div className="w-3/4 px-4 flex flex-row justify-between">
        <a className="inline-block" href="#">Курсы</a>      
        <a className="inline-block" href="#">Войти</a>  
        <a className="inline-block" href="#">Регистрация</a>  
      </div>     
    </div>
  )
}

export default Header