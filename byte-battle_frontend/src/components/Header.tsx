import React, { useState, useEffect } from 'react';
import RegisterWindow from './RegisterWindow';
import LoginWindow from './LoginWindow';


function Header(): JSX.Element {

  const [showRegisterWindow, setShowRegisterWindow] = useState(false);
  const [showLoginWindow, setShowLoginWindow] = useState(false);

  useEffect(() => {
    const logoButton = document.getElementById("logo-button");
    logoButton.addEventListener('click', () => window.location.replace(`${import.meta.env.PUBLIC_PROTOCOL}://${import.meta.env.PUBLIC_HOSTNAME}:3000/`));
  }, []);

  return (
    <div>
      {showRegisterWindow ? (
        <RegisterWindow closeWindow={setShowRegisterWindow} />
      ) : null}
      {
        showLoginWindow ? (
          <LoginWindow closeWindow={setShowLoginWindow} />
        ) : null}
      <div className="flex text-lg font-sans px-6 py-3 justify-between relative">
        <button
          id="logo-button"
        >
          <h1 className="text-gray-300/90 font-bold text-2xl">
            BYTE BATTLE
          </h1>
        </button>
        <div className="text-gray-200/80 font-medium font-custom">
          <button
            onClick={() => setShowRegisterWindow(!showRegisterWindow)}
            className="mr-4 px-4 py-1 rounded-full bg-gray-600/20
                            hover:bg-blue-600/20 hover:shadow-blue-600/10
                            transition ease-in-out duration-200 shadow-lg"
          >
            <h1>
              Register
            </h1>
          </button>
          <button
            onClick={() => setShowLoginWindow(!showLoginWindow)}
            className="mr-2 px-4 py-1 rounded-full bg-gray-600/20
                            hover:bg-blue-600/30 hover:shadow-blue-600/10
                            transition ease-in-out duration-200 shadow-lg"
          >
            <h1>
              Login
            </h1>
          </button>
        </div>
      </div>
    </div>
  );
};

export default Header;