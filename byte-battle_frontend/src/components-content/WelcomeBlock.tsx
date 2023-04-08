import React from "react";

const WelcomeBlock = ():JSX.Element => {
  return (
    <div className="pt-20 flex flex-col items-center justify-center">
      <div className="mb-4 text-center">
        <p>Improve in programming<br/>through modern courses</p>
      </div>
      <div className="mb-4 text-center">
        <p>Develop your skills<br/>through problems solving</p>
      </div>
      <a href="#">Get Started</a>
    </div>
  )
}

export default WelcomeBlock