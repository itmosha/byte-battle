import React from "react"
import Header from "./Header"
import WelcomeBlock from "../components-content/WelcomeBlock"

const MainPageContent = ():JSX.Element => {
  return (
    <div className="main-page">
      <WelcomeBlock />
    </div>
  )
}

export default MainPageContent