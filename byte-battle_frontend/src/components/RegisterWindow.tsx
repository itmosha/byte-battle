import React from 'react'


interface Props {
    closeWindow: Function,
}

function RegisterWindow(props: Props): JSX.Element {
    return (
        <div
            className="w-screen h-screen absolute bg-black bg-opacity-30 z-10"
        >
            <div className="h-2/5 w-2/5 absolute left-1/3 text-yellow-200">
                <h1 className="font-custom">Register</h1>
                <button onClick={() => props.closeWindow(false)}>
                    Toggle
                </button>
            </div>
        </div>
    );
};

export default RegisterWindow;
