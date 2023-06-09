import React, { ChangeEvent, useState } from 'react'


interface Props {
    closeWindow: Function,
}

function LoginWindow(props: Props): JSX.Element {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");


    const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { id, value } = e.target;

        switch (id) {
            case "username" : {
                setUsername(value);
                break;
            }
            case "password" : {
                setPassword(value);
                break;
            }
        }
    }

    const handleSubmit = async () => {
        try {
            const url = `${import.meta.env.PUBLIC_PROTOCOL}://${import.meta.env.PUBLIC_HOSTNAME}:8080/api/login/`;
            const response = await fetch(url, {
                method: 'POST',
                body: JSON.stringify({
                    "Username": username,
                    "Password": password,
                })
            });

            if (response.status === 200) {
                const responseJSON = await response.json();

                localStorage.setItem("token", responseJSON.token);
                props.closeWindow(false);
            } else {
                alert('Failed!');
            }
        } catch (err: any) {
            console.log(err);
        }
    }

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
                    <h1 className="font-custom text-3xl">Login</h1>
                </div>
                <div className="items-center">
                    <input 
                        type="text" 
                        placeholder="Username" 
                        id="username" 
                        value={username} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className="w-[256px] my-1 p-2 rounded-md bg-gray-900/50 border-2 border-gray-900/40 focus:border-blue-800 outline-none"
                    />
                    <input 
                        type="password" 
                        placeholder="Password" 
                        id="password" 
                        value={password} 
                        onChange={(e: ChangeEvent<HTMLInputElement>) => handleInputChange(e)}
                        className="w-[256px] my-1 p-2 rounded-md bg-gray-900/50 border-2 border-gray-900/40 focus:border-blue-800 outline-none"
                    />
                    <div className="flex px-12 mt-8">
                        <button 
                            type="submit" 
                            onClick={() => { handleSubmit(); } }
                            className="bg-blue-900 w-[160px] rounded-md py-2 hover:bg-blue-800 hover:shadow-blue-800/10 shadow-lg transition ease-in-out duration-200 font-custom"
                        >
                            Sign in
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default LoginWindow;
