import React, { useState, useEffect, useRef } from 'react';
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
            { showRegisterWindow ? (
                <RegisterWindow closeWindow={setShowRegisterWindow} />
            ) : null }
            {
                showLoginWindow ? (
                    <LoginWindow closeWindow={setShowLoginWindow} />
            ) : null }
            <div className="flex text-lg font-sans px-6 py-3 justify-between relative">
                <button
                    id="logo-button"
                >
                    <h1 className="text-gray-300/90 font-semibold text-2xl">
                        BYTE BATTLE
                    </h1>
                </button>
                <div className="text-gray-300/70 font-medium">
                    <button 
                        onClick={() => setShowRegisterWindow(!showRegisterWindow)}
                        className="mr-2 px-4 py-1 rounded-full uppercase
                            hover:bg-gray-500 hover:bg-opacity-20 hover:shadow-gray-500/20
                            transition ease-in-out duration-200 shadow-lg"
                    >
                        <h1>
                            Register
                        </h1>
                    </button>
                    <button 
                        onClick={() => setShowLoginWindow(!showLoginWindow)}
                        className="mr-2 px-4 py-1 rounded-full uppercase
                            hover:bg-gray-500 hover:bg-opacity-20 hover:shadow-gray-500/20
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