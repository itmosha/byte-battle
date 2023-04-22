import React from 'react'


interface Props {
    closeWindow: Function,
}

function RegisterWindow(props: Props): JSX.Element {
    return (
        <div
            className="w-screen h-screen absolute z-10 
                bg-black bg-opacity-20 animate-fadein"
        >
            <div className="w-80 px-8 py-6 absolute left-[calc(50%-160px)] top-[calc(50%-192px)] rounded-2xl
                text-gray-200/80 shadow-xl bg-gradient-to-br from-blue-950 to-gray-900
                "
            >
                <div className="text-center pb-4">
                    <h1 className="font-custom text-3xl">Register</h1>
                </div>
                <div className="items-center">
                    <input type="text" placeholder="Email" id="email" className="w-[256px] my-1 p-2 rounded-md bg-gray-900/50 border-2 border-gray-900/40 focus:border-blue-800 outline-none"></input>
                    <input type="text" placeholder="Username" id="username" className="w-[256px] my-1 p-2 rounded-md bg-gray-900/50 border-2 border-gray-900/40 focus:border-blue-800 outline-none"></input>
                    <input type="password" placeholder="Password" id="password" className="w-[256px] my-1 p-2 rounded-md bg-gray-900/50 border-2 border-gray-900/40 focus:border-blue-800 outline-none"></input>
                    <input type="password" placeholder="Confirm Password" id="confirmPassword" className="w-[256px] my-1 p-2 rounded-md bg-gray-900/50 border-2 border-gray-900/40 focus:border-blue-800 outline-none"></input>
                    <div className="flex px-12 mt-8">
                        <button 
                            type="submit" 
                            onClick={() => props.closeWindow(false)}
                            className="bg-blue-900 w-[160px] rounded-md py-2 hover:bg-blue-800 hover:shadow-blue-800/10 shadow-lg transition ease-in-out duration-200 font-custom"
                        >
                            Sign up
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default RegisterWindow;
