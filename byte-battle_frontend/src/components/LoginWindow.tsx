
import React from 'react'


interface Props {
    closeWindow: Function,
}

function LoginWindow(props: Props): JSX.Element {
    return (
        <div
            className="w-screen h-screen absolute bg-black bg-opacity-50 z-10"
        >
            <div className="h-2/5 w-2/5">
                <h1>Login</h1>
                <button onClick={() => props.closeWindow(false)}>
                    Toggle
                </button>
            </div>
        </div>
    );
};

export default LoginWindow;
